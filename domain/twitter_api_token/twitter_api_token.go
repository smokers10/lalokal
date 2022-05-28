package twitter_api_token

import (
	"lalokal/domain/http_response"
)

type TwitterAPIToken struct {
	Id      string `json:"id" form:"id" bson:"_id"`
	Token   string `json:"token" form:"token" bson:"token"`
	Secret  string `json:"secret" form:"secret" bson:"secret"`
	TopicId string `json:"topic_id" form:"topic_id" bson:"topic_id"`
}

type Repository interface {
	Upsert(data *TwitterAPIToken) (failure error)

	FindOneById(topic_id string) (result *TwitterAPIToken)
}

type Service interface {
	Store(input *TwitterAPIToken) (response *http_response.Response)

	Read(topic_id string) (response *http_response.Response)
}
