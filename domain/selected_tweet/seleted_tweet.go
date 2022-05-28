package selected_tweet

import "lalokal/domain/http_response"

type SelectedTweet struct {
	Id                string `json:"id,omitempty" form:"id" bson:"_id"`
	AuthorId          string `json:"author_id,omitempty" form:"author_id" bson:"author_id"`
	TweetId           string `json:"tweet_id,omitempty" form:"tweet_id" bson:"tweet_id"`
	Text              string `json:"text,omitempty" form:"text" bson:"text"`
	BlastingSessionId string `json:"blasting_session_id,omitempty" form:"blasting_session_id" bson:"blasting_session_id"`
}

type Repository interface {
	Insert(data *SelectedTweet) (failure error)

	Delete(selected_tweet_id string) (failure error)

	FindByBlastingSessionId(blasting_session_id string) (result []SelectedTweet)
}

type Service interface {
	Store(input *SelectedTweet) (response *http_response.Response)

	Delete(blasting_session_id string, selected_tweet_id string) (response *http_response.Response)

	ReadAll(blasting_session_id string) (response *http_response.Response)
}
