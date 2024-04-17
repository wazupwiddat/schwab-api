package client

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

type SchwabAPIClient struct {
	HTTPClient    *http.Client
	TraderBaseURL string
	OAuthBaseURL  string
	ClientID      string
	ClientSecret  string
	AuthRedirect  string
}

func NewSchwabClient(clientID string, secret string, authRedirect string) *SchwabAPIClient {
	return &SchwabAPIClient{
		HTTPClient:    &http.Client{Timeout: 10 * time.Second},
		OAuthBaseURL:  "https://api.schwabapi.com/v1/",
		TraderBaseURL: "https://api.schwabapi.com/trader/v1",
		ClientID:      clientID,
		ClientSecret:  secret,
		AuthRedirect:  authRedirect,
	}
}

func (c SchwabAPIClient) BasisAuth() string {
	basicAuth := fmt.Sprintf("%s:%s", c.ClientID, c.ClientSecret)
	encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(basicAuth))

	return encodedBasicAuth
}
