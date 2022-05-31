package topic

import (
	"lalokal/domain/http_response"
	"lalokal/domain/topic"
)

type topicService struct {
	topicRepository topic.Repository
}

func TopicService(tr *topic.Repository) topic.Service {
	return &topicService{topicRepository: *tr}
}

func (s *topicService) Store(input *topic.Topic) (response *http_response.Response) {
	if msg, isfail := storeValidation(input); isfail {
		return &http_response.Response{
			Message: msg,
			Status:  400,
		}
	}

	inserted_id, err := s.topicRepository.Insert(input)
	if err != nil {
		return &http_response.Response{
			Message: "kesalahan saat menyimpan topic",
			Status:  500,
		}
	}

	input.Id = inserted_id

	return &http_response.Response{
		Message: "topic tersimpan",
		Success: true,
		Status:  200,
		Data:    input,
	}
}

func (s *topicService) Update(input *topic.Topic) (response *http_response.Response) {
	if msg, isfail := updateValidation(input); isfail {
		return &http_response.Response{
			Message: msg,
			Status:  400,
		}
	}

	if err := s.topicRepository.Update(input); err != nil {
		return &http_response.Response{
			Message: "kesalahan saat update topic",
			Status:  500,
		}
	}

	return &http_response.Response{
		Message: "topic terupdate",
		Success: true,
		Status:  200,
	}
}

func (s *topicService) ReadAll(user_id string) (response *http_response.Response) {
	if user_id == "" {
		return &http_response.Response{
			Message: "id pengguna tidak boleh kosong",
			Status:  400,
		}
	}

	result := s.topicRepository.FindByUserId(user_id)

	return &http_response.Response{
		Message: "topik berhasil diambil",
		Success: true,
		Status:  200,
		Data:    result,
	}
}

func (s *topicService) Detail(topic_id string) (response *http_response.Response) {
	if topic_id == "" {
		return &http_response.Response{
			Message: "id topik tidak boleh kosong",
			Status:  400,
		}
	}

	result := s.topicRepository.FindOneById(topic_id)

	return &http_response.Response{
		Message: "topik berhasil diambil",
		Success: true,
		Status:  200,
		Data:    result,
	}
}
