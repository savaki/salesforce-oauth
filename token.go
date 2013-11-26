package oauth

type Token struct {
	Id               string `json:"id,omitempty"`
	IssuedAt         string `json:"issued_at,omitempty"`
	Scope            string `json:"scope,omitempty"`
	Error            string `json:"error,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
	RefreshToken     string `json:"refresh_token,omitempty"`
	InstanceUrl      string `json:"instance_url,omitempty"`
	Signature        string `json:"signature,omitempty"`
	AccessToken      string `json:"access_token,omitempty"`
}
