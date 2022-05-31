package twitter_api_token

import "lalokal/domain/twitter_api_token"

func storeValidation(i *twitter_api_token.TwitterAPIToken) (string, bool) {
	if i.Secret == "" {
		return "secret token tidak boleh kosong", true
	}

	if i.Token == "" {
		return "token tidak boleh kosong", true
	}

	if i.TopicId == "" {
		return "id topik tidak boleh kosong", true
	}

	return "", false
}
