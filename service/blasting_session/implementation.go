package blasting_session

import (
	"fmt"
	"lalokal/domain/blasting_log"
	"lalokal/domain/blasting_session"
	"lalokal/domain/http_response"
	"lalokal/domain/keyword"
	"lalokal/domain/selected_tweet"
	"lalokal/domain/twitter_api_token"
	twitter_http_request "lalokal/infrastructure/lib/twitter_api"
)

type blastingSessionService struct {
	blastingLogRepository     blasting_log.Repository
	blastingSessionRepository blasting_session.Repository
	twitterAPITokenRepository twitter_api_token.Repository
	keywordRepository         keyword.Repository
	twitter                   twitter_http_request.Contract
}

func BlastingSessionService(blr *blasting_log.Repository, bsr *blasting_session.Repository, tat *twitter_api_token.Repository,
	kr *keyword.Repository, tw *twitter_http_request.Contract) blasting_session.Service {
	return &blastingSessionService{
		blastingLogRepository:     *blr,
		blastingSessionRepository: *bsr,
		twitterAPITokenRepository: *tat,
		keywordRepository:         *kr,
		twitter:                   *tw,
	}
}

func (s *blastingSessionService) Store(input *blasting_session.BlastingSession) (response *http_response.Response) {
	if msg, isfail := storeValidation(input); isfail {
		return &http_response.Response{
			Message: msg,
			Status:  400,
		}
	}

	input.Status = "waiting"

	// store blasting session
	inserted_id, err := s.blastingSessionRepository.Insert(input)
	if err != nil {
		return &http_response.Response{
			Message: "kesalahan ketika menyimpan sesi blasting",
			Status:  500,
		}
	}

	input.Id = inserted_id

	return &http_response.Response{
		Message: "sesi blasting tersimpan",
		Success: true,
		Status:  200,
		Data:    input,
	}
}

func (s *blastingSessionService) Update(input *blasting_session.BlastingSession) (response *http_response.Response) {
	if msg, isfail := updateValidation(input); isfail {
		return &http_response.Response{
			Message: msg,
			Status:  400,
		}
	}

	if err := s.blastingSessionRepository.Update(input); err != nil {
		return &http_response.Response{
			Message: "kesalahan ketika update sesi blasting",
			Status:  500,
		}
	}

	return &http_response.Response{
		Message: "sesi blasting terupdate",
		Success: true,
		Status:  200,
		Data:    input,
	}
}

func (s *blastingSessionService) ReadAll(topic_id string) (response *http_response.Response) {
	if topic_id == "" {
		return &http_response.Response{
			Message: "id topik tidak boleh kosong",
			Status:  400,
		}
	}

	result := s.blastingSessionRepository.FindByTopicId(topic_id)
	for i := 0; i < len(result); i++ {
		total_count, success_count, failed_count, success_percentage, fail_percentage := s.blastingLogRepository.LogPercentage(result[i].Id)
		result[i].FailedPercentage = fail_percentage
		result[i].SuccessPercentage = success_percentage
		result[i].FailedCount = failed_count
		result[i].SuccessCount = success_count
		result[i].TotalCount = total_count
	}

	return &http_response.Response{
		Message: "sesi blasting berhasil diambil",
		Success: true,
		Status:  200,
		Data:    result,
	}
}

func (s *blastingSessionService) Detail(blasting_session_id string) (response *http_response.Response) {
	if blasting_session_id == "" {
		return &http_response.Response{
			Message: "id tidak boleh kosong",
			Status:  400,
		}
	}

	result := s.blastingSessionRepository.FindById(blasting_session_id)

	return &http_response.Response{
		Message: "sesi blasting berhasil diambil",
		Success: true,
		Status:  200,
		Data:    result,
	}
}

func (s *blastingSessionService) Scrape(blasting_session_id string) (response *http_response.Response) {
	tweets := []twitter_http_request.ScrapedTweet{}

	if blasting_session_id == "" {
		return &http_response.Response{
			Message: "id tidak boleh kosong",
			Status:  400,
		}
	}

	// retrieve blasting session
	blasting_session := s.blastingSessionRepository.FindById(blasting_session_id)

	// if blasting session not exists
	if blasting_session.Id == "" {
		return &http_response.Response{
			Message: "sesi blasting tidak ditemukan",
			Status:  404,
		}
	}

	// retireve twitter token
	twitter_token := s.twitterAPITokenRepository.FindOneByTopicId(blasting_session.TopicId)
	if twitter_token.APIToken == "" {
		return &http_response.Response{
			Message: "api token twitter untuk topik tidak ada",
			Status:  404,
		}
	}

	// retrieve keyword
	keywords := s.keywordRepository.FindByTopicId(blasting_session.TopicId)
	if len(keywords) == 0 {
		return &http_response.Response{
			Message: "kata kunci kosong",
			Status:  404,
		}
	}

	for _, k := range keywords {
		scraped_tweets, err := s.twitter.Search(k.Keyword, twitter_token.APIToken)
		if err != nil {
			fmt.Println(err)
			return &http_response.Response{
				Message: "kesalahan ketika mengambil tuitan",
				Status:  500,
			}
		}

		tweets = append(tweets, scraped_tweets.Data...)
	}

	return &http_response.Response{
		Message: "tuitan berhasil diambil",
		Success: true,
		Status:  200,
		Data:    tweets,
	}
}

func (s *blastingSessionService) Blast(blasting_session_id string, tweets []selected_tweet.SelectedTweet) (response *http_response.Response) {
	// check blasting session id
	if blasting_session_id == "" {
		return &http_response.Response{
			Message: "id tidak boleh kosong",
			Status:  400,
		}
	}

	// check selected tweet, if no selected tweet then retrun error response 400
	if len(tweets) == 0 {
		return &http_response.Response{
			Message: "tidak ada tuitan yang dipilih",
			Status:  400,
		}
	}

	// retrieve blasting session
	blasting_session := s.blastingSessionRepository.FindById(blasting_session_id)

	// if blasting session not exists
	if blasting_session.Id == "" {
		return &http_response.Response{
			Message: "sesi blasting tidak ditemukan",
			Status:  404,
		}
	}

	// if blasting session status is finish
	if blasting_session.Status == "revoked" {
		return &http_response.Response{
			Message: "sesi blasting sudah selesai",
			Status:  403,
		}
	}

	// retireve twitter token
	twitter_token := s.twitterAPITokenRepository.FindOneByTopicId(blasting_session.TopicId)
	if twitter_token.APIToken == "" || twitter_token.AccessSecret == "" || twitter_token.AccessToken == "" || twitter_token.ConsumerKey == "" || twitter_token.ConsumerSecret == "" {
		return &http_response.Response{
			Message: "token twitter tidak boleh kosong",
			Status:  404,
		}
	}

	// start BLASTING!!
	for _, tweet := range tweets {
		EO := twitter_http_request.EOMap{
			"event": twitter_http_request.EOMap{
				"type": "message_create",
				"message_create": twitter_http_request.EOMap{
					"target":       twitter_http_request.EOMap{"recipient_id": tweet.AuthorId},
					"message_data": twitter_http_request.EOMap{"text": blasting_session.Message},
				},
			},
		}

		_, DER := s.twitter.DirectMessage(*twitter_token, EO)
		fmt.Println(DER)
		if DER != nil {
			err := s.blastingLogRepository.Insert(&blasting_log.BlastingLogDomain{
				Status:            "not sent",
				BlastingSessionId: blasting_session_id,
				TopicId:           blasting_session.TopicId,
			})

			if err != nil {
				return &http_response.Response{
					Message: "gagal menyimpan log",
					Status:  500,
				}
			}
		}

		err := s.blastingLogRepository.Insert(&blasting_log.BlastingLogDomain{
			Status:            "sent",
			BlastingSessionId: blasting_session_id,
			TopicId:           blasting_session.TopicId,
		})

		if err != nil {
			return &http_response.Response{
				Message: "gagal menyimpan log",
				Status:  500,
			}
		}
	}

	// update blasting session status
	if err := s.blastingSessionRepository.UpdateStatus(blasting_session_id, "revoked"); err != nil {
		return &http_response.Response{
			Message: "gagal update status sesi blasting",
			Status:  500,
		}
	}

	return &http_response.Response{
		Message: "blasting selesai",
		Success: true,
		Status:  200,
	}
}

func (s *blastingSessionService) Count(topic_id string) (response *http_response.Response) {
	// check blasting session id
	if topic_id == "" {
		return &http_response.Response{
			Message: "id topik tidak boleh kosong",
			Status:  400,
		}
	}

	BScount := s.blastingSessionRepository.Count(topic_id)
	Kcount := s.keywordRepository.Cound(topic_id)
	BSLCount := s.blastingLogRepository.Count(topic_id)
	apiToken := s.twitterAPITokenRepository.FindOneByTopicId(topic_id)
	var isTokenSet bool

	if apiToken.Id == "" {
		isTokenSet = false
	} else {
		isTokenSet = true
	}

	return &http_response.Response{
		Message: "perhitungan berhasil diambil",
		Success: true,
		Status:  200,
		Data: map[string]interface{}{
			"blasting_session_count":     BScount,
			"keyword_count":              Kcount,
			"blasting_session_log_count": BSLCount,
			"is_token_set":               isTokenSet,
		},
	}
}

func (s *blastingSessionService) Monitoring(blasting_session_id string) (response *http_response.Response) {
	// check blasting session id
	if blasting_session_id == "" {
		return &http_response.Response{
			Message: "id tidak boleh kosong",
			Status:  400,
		}
	}

	total_message, success_count, failed_count, success_percentage, fail_percentage := s.blastingLogRepository.LogPercentage(blasting_session_id)
	return &http_response.Response{
		Message: "data monitoring berhasil diambil",
		Success: true,
		Status:  200,
		Data: map[string]interface{}{
			"total_message":      total_message,
			"success_percentage": success_percentage,
			"fail_percentage":    fail_percentage,
			"success_count":      success_count,
			"failed_count":       failed_count,
		},
	}
}
