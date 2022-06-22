package twitter_http_request

import (
	"lalokal/domain/twitter_api_token"
	"time"
)

type EOMap map[string]interface{}

type ScrapedTweet struct {
	Id        string     `json:"id"`
	Text      string     `json:"text"`
	CreatedAt time.Time  `json:"created_at"`
	AuthorId  string     `json:"author_id"`
	Author    UserDetail `json:"author"`
}

type RetrunValue struct {
	Data []ScrapedTweet `json:"data"`
}

type DMErrorResponse struct {
	Errors []struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"errors"`
}

type DMSuccessResponse struct {
	Event struct {
		Type             string `json:"type"`
		Id               string `json:"id"`
		CreatedTimestamp string `json:"created_timestamp"`
		MessageCreate    struct {
			Target struct {
				RecipientId string `json:"recipient_id"`
			} `json:"target"`
			SenderId    string `json:"sender_id"`
			MessageData struct {
				Text string `json:"text"`
			}
		} `json:"message_create"`
	} `json:"event"`
}

type UserDetail struct {
	Data struct {
		ProfileImageURL string `json:"profile_image_url"`
		Username        string `json:"username"`
		URL             string `json:"url"`
		Name            string `json:"name"`
	} `json:"data"`
}

type Contract interface {
	Search(keyword string, token string) (scraped_tweet *RetrunValue, failure error)

	DirectMessage(token twitter_api_token.TwitterAPIToken, event_object EOMap) (DSR *DMSuccessResponse, DER *DMErrorResponse)
}
