package configuration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Application struct {
	BaseURL string `json:"base_url"`
	Port    string `json:"port"`
	Secret  string `json:"secret"`
	Mode    string `json:"mode"`
}

type SMTP struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Sender   string `json:"sender"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Database struct {
	MongoURI          string `json:"mongo_uri"`
	DatabaseName      string `json:"database_name"`
	SessionStorage    string `json:"session_storage"`
	MaxPool           int    `json:"max_pool"`
	MinPool           int    `json:"min_pool"`
	MaxIdleConnection int    `json:"max_idle_connection"`
}

type Configuration struct {
	Application Application `json:"application"`
	SMTP        SMTP        `json:"smtp"`
	Database    Database    `json:"database"`
}

func ReadConfiguration() (configuration *Configuration) {
	r, err := ioutil.ReadFile("infrastructure/configuration/configuration.json")
	if err != nil {
		fmt.Println(err)
		r, err := ioutil.ReadFile("../../infrastructure/configuration/configuration.json")
		if err != nil {
			fmt.Println(err)
		}

		return unmarshalData(r)
	}

	return unmarshalData(r)
}

func unmarshalData(r []byte) (configuration *Configuration) {
	cfg := Configuration{}
	json.Unmarshal(r, &cfg)

	return &cfg
}
