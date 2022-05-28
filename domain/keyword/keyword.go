package keyword

import "lalokal/domain/http_response"

type Keyword struct {
	Id      string `json:"id,omitempty" form:"id" bson:"_id"`
	Keyword string `json:"keyword,omitempty" form:"keyword" bson:"keyword"`
	TopicId string `json:"topic_id,omitempty" form:"topic_id" bson:"topic_id"`
}

type Repository interface {
	Insert(data *Keyword) (failure error)

	Delete(keyword_id string) (failure error)

	FindByTopicId(topic_id string) (result []Keyword)
}

type Service interface {
	Store(data *Keyword) (response *http_response.Response)

	Delete(keyword_id string) (response *http_response.Response)

	ReadAll(topic_id string) (response *http_response.Response)
}
