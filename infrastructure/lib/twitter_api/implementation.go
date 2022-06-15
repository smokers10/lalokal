package twitter_http_request

type twitterHTTP struct{}

// DirectMessage implements Contract
func (*twitterHTTP) DirectMessage(author_id string, message string) (failure error) {
	panic("unimplemented")
}

// Search implements Contract
func (*twitterHTTP) Search(keyword string) (scraped_tweet []map[string]interface{}, failure error) {
	panic("unimplemented")
}

func TwitterHTTP() Contract {
	return &twitterHTTP{}
}
