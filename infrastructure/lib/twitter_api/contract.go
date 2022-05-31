package twitter_http_request

type Contract interface {
	Search(keyword string) (scraped_tweet []map[string]interface{}, failure error)

	DirectMessage(author_id string, message string) (failure error)
}
