package blasting_session

import "lalokal/domain/http_response"

type BlastingSession struct {
	Id      string `json:"id,omitempty" form:"id" bson:"_id"`
	Title   string `json:"title,omitempty" form:"title" bson:"title"`
	Message string `json:"message,omitempty" form:"message" bson:"message"`
	TopicId string `json:"topic_id,omitempty" form:"topic_id" bson:"topic_id"`
}

type Repository interface {
	Insert(data *BlastingSession) (failure error)

	Update(data *BlastingSession) (failure error)

	FindByTopicId(topic_id string) (result []BlastingSession)

	FindById(blasting_session_id string) (result *BlastingSession)
}

type Service interface {
	Store(input *BlastingSession) (response *http_response.Response)

	Update(data *BlastingSession) (response *http_response.Response)

	ReadAll(topic_id string) (response *http_response.Response)

	Detail(blasting_session_id string) (response *http_response.Response)
}
