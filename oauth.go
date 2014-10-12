package oauth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type OAuth struct {
	ConsumerKey    string `json:"consumer-key"`
	ConsumerSecret string `json:"consumer-secret"`
	RedirectUri    string `json:"redirect-uri"`
}

func doWithJsonResponse(request *http.Request, obj interface{}) error {
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	log.Println("")
	log.Printf("response code => %d", response.StatusCode)
	for header := range response.Header {
		log.Printf("%s: %s", header, response.Header.Get(header))
	}
	log.Println("")

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	json.Unmarshal(bytes, obj)
	return nil
}

func ReadFile(filename string) (*OAuth, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	value := new(OAuth)
	err = json.Unmarshal(bytes, value)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (self *OAuth) Url(state string) string {
	params := url.Values{
		"response_type": {"code"},
		"client_id":     {self.ConsumerKey},
		"redirect_uri":  {self.RedirectUri},
		"state":         {state},
	}
	return fmt.Sprintf("https://login.salesforce.com/services/oauth2/authorize?%s", params.Encode())
}

func (self *OAuth) FetchToken(code string) (*Token, error) {
	params := url.Values{
		"format":        {"json"},
		"grant_type":    {"authorization_code"},
		"client_id":     {self.ConsumerKey},
		"client_secret": {self.ConsumerSecret},
		"redirect_uri":  {self.RedirectUri},
		"code":          {code},
	}

	log.Printf("FetchToken(%s)", code)
	log.Printf("params => %s", params.Encode())

	request, _ := http.NewRequest("POST", "https://login.salesforce.com/services/oauth2/token", strings.NewReader(params.Encode()))
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	token := new(Token)
	if err := doWithJsonResponse(request, token); err != nil {
		log.Printf("unable to extract json content => %s", err.Error())
		return nil, err
	}

	return token, nil
}
