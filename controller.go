package oauth

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func doWithJsonResponse(request *http.Request, obj interface{}) error {
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	json.Unmarshal(bytes, obj)
	return nil
}

func (self *OAuth) Callback(w http.ResponseWriter, r *http.Request) {
	code := r.PostFormValue("code")
	self.FetchToken(code)
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

	request, _ := http.NewRequest("POST", "https://login.salesforce.com/services/oauth2/token", strings.NewReader(params.Encode()))
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	token := new(Token)
	if err := doWithJsonResponse(request, token); err != nil {
		log.Printf("unable to extract json content => %s", err.Error())
		return nil, err
	} else {
		log.Printf("unable to extract json content => %#v", token)
		return token, nil
	}
}
