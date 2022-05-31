package twitter_api_token

import (
	"lalokal/domain/http_response"
	"lalokal/domain/twitter_api_token"
)

type twitterAPIService struct {
	twitterAPIRepository twitter_api_token.Repository
}

func TwitterAPIService(tat *twitter_api_token.Repository) twitter_api_token.Service {
	return &twitterAPIService{twitterAPIRepository: *tat}
}

func (s *twitterAPIService) Store(input *twitter_api_token.TwitterAPIToken) (response *http_response.Response) {
	if msg, isfail := storeValidation(input); isfail {
		return &http_response.Response{Message: msg, Status: 400}
	}

	if err := s.twitterAPIRepository.Upsert(input); err != nil {
		return &http_response.Response{Message: "kesalahan ketika menyimpan token twitter", Status: 500}
	}

	return &http_response.Response{
		Message: "token twitter tersimpan",
		Success: true,
		Status:  200,
	}
}

func (s *twitterAPIService) Read(topic_id string) (response *http_response.Response) {
	if topic_id == "" {
		return &http_response.Response{
			Message: "id topik tidak boleh kosong",
			Status:  400,
		}
	}

	result := s.twitterAPIRepository.FindOneByTopicId(topic_id)

	return &http_response.Response{
		Message: "token twitter berhasil diambil",
		Success: true,
		Status:  200,
		Data:    result,
	}
}
