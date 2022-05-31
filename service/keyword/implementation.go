package keyword

import (
	"lalokal/domain/http_response"
	"lalokal/domain/keyword"
)

type keywordService struct {
	keywordRepo keyword.Repository
}

func KeywordService(kr *keyword.Repository) keyword.Service {
	return &keywordService{keywordRepo: *kr}
}

func (s *keywordService) Store(input *keyword.Keyword) (response *http_response.Response) {
	if msg, isfail := storeValidation(input); isfail {
		return &http_response.Response{
			Message: msg,
			Status:  400,
		}
	}

	inserted_id, err := s.keywordRepo.Insert(input)
	if err != nil {
		return &http_response.Response{
			Message: "kesalahan ketika menyimpan kata kunci",
			Status:  500,
		}
	}

	input.Id = inserted_id

	return &http_response.Response{
		Message: "kata kunci tersimpan",
		Success: true,
		Status:  200,
		Data:    input,
	}
}

func (s *keywordService) Delete(keyword_id string) (response *http_response.Response) {
	if keyword_id == "" {
		return &http_response.Response{
			Message: "id tidak boleh kosong",
			Status:  400,
		}
	}

	if err := s.keywordRepo.Delete(keyword_id); err != nil {
		return &http_response.Response{
			Message: "kesalahan saat menghapus kata kunci",
			Status:  500,
		}
	}

	return &http_response.Response{
		Message: "kata kunci terhpaus",
		Success: true,
		Status:  200,
	}
}

func (s *keywordService) ReadAll(topic_id string) (response *http_response.Response) {
	if topic_id == "" {
		return &http_response.Response{
			Message: "id topik tidak boleh kosong",
			Status:  400,
		}
	}

	result := s.keywordRepo.FindByTopicId(topic_id)

	return &http_response.Response{
		Message: "kata kunci berhasil diambil",
		Success: true,
		Status:  200,
		Data:    result,
	}
}
