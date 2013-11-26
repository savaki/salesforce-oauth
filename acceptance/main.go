package main

import (
	oauth "../."
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var sfdc *oauth.OAuth = nil

func Callback(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	token, err := sfdc.FetchToken(code)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("nope, not working"))

	} else {
		buffer, _ := json.Marshal(token)
		log.Printf("received token, %s", string(buffer))
		w.WriteHeader(200)
		w.Write(buffer)
	}
}

func main() {
	sfdc, _ = oauth.ReadFile("oauth.json")

	fmt.Println(sfdc.Url("mobile"))

	router := mux.NewRouter()
	router.Path("/oauth/sfdc").HandlerFunc(Callback)
	http.ListenAndServe(":8080", router)
}
