package selected_tweet

import (
	"time"
)

type SelectedTweet struct {
	Id                string    `json:"id,omitempty" form:"id" bson:"_id"`
	AuthorId          string    `json:"author_id,omitempty" form:"author_id" bson:"author_id"`
	TweetId           string    `json:"tweet_id,omitempty" form:"tweet_id" bson:"tweet_id"`
	Text              string    `json:"text,omitempty" form:"text" bson:"text"`
	CreatedAt         time.Time `json:"created_at,omitempty" form:"created_at" bson:"created_at"`
	BlastingSessionId string    `json:"blasting_session_id,omitempty" form:"blasting_session_id" bson:"blasting_session_id"`
	Author            Author    `json:"author,omitempty"`
}

type Author struct {
	Data struct {
		ProfileImageURL string `json:"profile_image_url"`
		Username        string `json:"username"`
		URL             string `json:"url"`
		Name            string `json:"name"`
	} `json:"data"`
}
