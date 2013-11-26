package oauth

type OAuth struct {
	ConsumerKey    string `json:"consumer-key"`
	ConsumerSecret string `json:"consumer-secret"`
	RedirectUri    string `json:"redirect-uri"`
}
