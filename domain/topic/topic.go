package topic

import "lalokal/domain/http_response"

type Topic struct {
	Id          string `json:"id" form:"id" bson:"_id"`
	Title       string `json:"title" form:"title" bson:"title"`
	Description string `json:"description" form:"description" bson:"description"`
	UserId      string `json:"user_id" form:"user_id" bson:"user_id"`
}

type Repository interface {
	Insert(data *Topic) (failure error)

	Update(data *Topic) (failure error)

	FindByUserId(user_id string) (result []Topic)

	FindOneById(topic_id string) (result *Topic)
}

type Service interface {
	Store(input *Topic) (response *http_response.Response)

	Update(input *Topic) (response *http_response.Response)

	ReadAll(user_id string) (response *http_response.Response)

	Detail(topic_id string) (response *http_response.Response)
}
