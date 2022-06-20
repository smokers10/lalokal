package twitter_http_request

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type twitterHTTP struct{}

func TwitterHTTP() Contract {
	return &twitterHTTP{}
}

// DirectMessage implements Contract
func (*twitterHTTP) DirectMessage(author_id string, message string, token string) (failure error) {
	panic("unimplemented")
}

// Search implements Contract
func (*twitterHTTP) Search(keyword string, token string) (scraped_tweet *RetrunValue, failure error) {
	client := &http.Client{}
	request, err := http.NewRequest("get", "https://api.twitter.com/2/tweets/search/recent", nil)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	request.Header.Set("user-agent", "golang application")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	params := request.URL.Query()
	params.Add("query", keyword)
	params.Add("tweet.fields", "attachments,author_id,created_at,geo,id,possibly_sensitive,text,withheld")
	params.Add("user.fields", "created_at,description,entities,id,location,name,pinned_tweet_id,profile_image_url,protected,public_metrics,url,username,verified,withheld")
	params.Add("max_results", "10")

	request.URL.RawQuery = params.Encode()

	response, err := client.Do(request)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(&scraped_tweet); err != nil {
		fmt.Print(err)
		return nil, err
	}

	for i := 0; i < len(scraped_tweet.Data); i++ {
		ud := getUser(scraped_tweet.Data[i].AuthorId, token)
		scraped_tweet.Data[i].Author.Data = ud.Data
	}

	return scraped_tweet, nil
}

func getUser(userid string, token string) (result UserDetail) {
	client := &http.Client{}
	path := fmt.Sprintf("https://api.twitter.com/2/users/%s", userid)

	request, err := http.NewRequest("get", path, nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("user-agent", "golang application")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	params := request.URL.Query()
	params.Add("user.fields", "description,id,location,name,profile_image_url,url,username")
	request.URL.RawQuery = params.Encode()

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return result
}
