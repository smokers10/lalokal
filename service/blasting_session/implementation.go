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
	selectedTweetRepository   selected_tweet.Repository
}

func BlastingSessionService(blr *blasting_log.Repository, bsr *blasting_session.Repository, tat *twitter_api_token.Repository,
	kr *keyword.Repository, tw *twitter_http_request.Contract, str *selected_tweet.Repository) blasting_session.Service {
	return &blastingSessionService{
		blastingLogRepository:     *blr,
		blastingSessionRepository: *bsr,
		twitterAPITokenRepository: *tat,
		keywordRepository:         *kr,
		twitter:                   *tw,
		selectedTweetRepository:   *str,
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
	tweets := []map[string]interface{}{}

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
	if twitter_token.Secret == "" || twitter_token.Token == "" {
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
		scraped_tweets, err := s.twitter.Search(k.Keyword)
		if err != nil {
			fmt.Println(err)
			return &http_response.Response{
				Message: "kesalahan ketika mengambil tuitan",
				Status:  500,
			}
		}

		tweets = append(tweets, scraped_tweets...)
	}

	return &http_response.Response{
		Message: "tuitan berhasil diambil",
		Success: true,
		Status:  200,
		Data:    tweets,
	}
}

func (s *blastingSessionService) Blast(blasting_session_id string) (response *http_response.Response) {
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

	// if blasting session status is finish
	if blasting_session.Status == "revoked" {
		return &http_response.Response{
			Message: "sesi blasting sudah selesai",
			Status:  403,
		}
	}

	// retrieve selected tweets
	selected_tweet := s.selectedTweetRepository.FindByBlastingSessionId(blasting_session_id)

	// if there is no selected tweets
	if len(selected_tweet) == 0 {
		return &http_response.Response{
			Message: "tidak ada tuitan yang dipilih",
			Status:  404,
		}
	}

	// start BLASTING!!!
	for _, st := range selected_tweet {
		if err := s.twitter.DirectMessage(st.AuthorId, blasting_session.Message); err != nil {
			s.blastingLogRepository.Insert(&blasting_log.BlastingLogDomain{
				Status:            "gagal",
				BlastingSessionId: blasting_session_id,
				TopicId:           blasting_session.TopicId,
			})
		} else {
			s.blastingLogRepository.Insert(&blasting_log.BlastingLogDomain{
				Status:            "berhasil",
				BlastingSessionId: blasting_session_id,
				TopicId:           blasting_session.TopicId,
			})
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
