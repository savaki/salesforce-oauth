salesforce-oauth
================

Simple go library to handle OAuth with Salesforce.

### Sample oauth.json

```
{
	"consumer-key":"YOUR-CONSUMER-KEY",
	"consumer-secret":"YOUR-CONSUMER-SECRET",
	"redirect-uri":"YOUR-CALLBACK-URI"
}


```

### Minimalist Sample Application

```
package main

import (
	oauth "github.com/savaki/salesforce-oauth"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var sfdc *oauth.OAuth = nil

func Callback(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	token, _ := sfdc.FetchToken(code)
	buffer, _ := json.Marshal(token)
	w.WriteHeader(200)
	w.Write(buffer)
}

func main() {
	// read our configuration from oauth.json
	sfdc, _ = oauth.ReadFile("oauth.json")

	// displays the oauth url.  paste this into your browser
	fmt.Println(sfdc.Url("mobile"))

	// start up the web server to listen for the response
	router := mux.NewRouter()
	router.Path("/oauth/sfdc").HandlerFunc(Callback)
	http.ListenAndServe(":8080", router)
}
```
