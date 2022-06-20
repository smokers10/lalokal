package twitter_api_token

import "lalokal/domain/twitter_api_token"

func storeValidation(i *twitter_api_token.TwitterAPIToken) (string, bool) {
	if i.APIToken == "" {
		return "API token tidak boleh kosong", true
	}

	if i.AccessSecret == "" {
		return "access secret tidak boleh kosong", true
	}

	if i.AccessToken == "" {
		return "access token tidak boleh kosong", true
	}

	if i.ConsumerKey == "" {
		return "consumer key tidak boleh kosong", true
	}

	if i.ConsumerSecret == "" {
		return "consumer secret tidak boleh kosong", true
	}

	if i.TopicId == "" {
		return "id topik tidak boleh kosong", true
	}

	return "", false
}
