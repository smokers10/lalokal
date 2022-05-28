package twitter_http_request

type TwitterAPI interface {
	Search(keyword string) (scraped_tweet []map[string]interface{}, failure error)

	DirectMessage(author_id string, message string) (failure error)
}
