package twitter_api_token

import (
	"lalokal/domain/http_response"
)

type TwitterAPIToken struct {
	Id             string `json:"id" form:"id" bson:"_id"`
	APIToken       string `json:"api_token" form:"api_token" bson:"api_token"`
	ConsumerKey    string `json:"consumer_key" form:"consumer_key" bson:"consumer_key"`
	ConsumerSecret string `json:"consumer_secret" form:"consumer_secret" bson:"consumer_secret"`
	AccessToken    string `json:"access_token" form:"access_token" bson:"access_token"`
	AccessSecret   string `json:"access_secret" form:"access_secret" bson:"access_secret"`
	TopicId        string `json:"topic_id" form:"topic_id" bson:"topic_id"`
}

type Repository interface {
	Upsert(data *TwitterAPIToken) (failure error)

	FindOneByTopicId(topic_id string) (result *TwitterAPIToken)
}

type Service interface {
	Store(input *TwitterAPIToken) (response *http_response.Response)

	Read(topic_id string) (response *http_response.Response)
}
