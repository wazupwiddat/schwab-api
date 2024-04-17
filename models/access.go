package models

type SchwabAccess struct {
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refresh_token"` // valid for 7 days
	AccessToken  string `json:"access_token"`  // valid for 30 minutes
	IDToken      string `json:"id_token"`      // JWT
}
