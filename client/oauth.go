package client

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/wazupwiddat/schwab-api/models"
)

// oauth/token
func (c SchwabAPIClient) accessToken(payload string) (*models.SchwabAccess, error) {
	// make call to Schwab OAuth server to get new Access Token
	url := fmt.Sprintf("%s/%s", c.OAuthBaseURL, "oauth/token")
	request, err := http.NewRequest("POST", url, strings.NewReader(payload))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Set the Content-Type header to application/json
	request.Header.Set("Authorization", fmt.Sprintf("Basic %s", c.BasisAuth()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 400 {
		log.Println("SchwabAPIClient:", "Failed to get a new access token")
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println(string(body))
			return nil, err
		}
		return nil, err
	}

	var schwabAccess models.SchwabAccess
	if err := json.NewDecoder(resp.Body).Decode(&schwabAccess); err != nil {
		return nil, err
	}

	return &schwabAccess, nil
}
func (c SchwabAPIClient) CreateAccessToken(authCode string) (*models.SchwabAccess, error) {
	payload := fmt.Sprintf("grant_type=%s&code=%s&client_id=%s&redirect_uri=%s",
		"authorization_code", authCode, c.ClientID, c.AuthRedirect)

	return c.accessToken(payload)
}

func (c SchwabAPIClient) RefreshAccessToken(refreshToken *models.SchwabAccess) error {
	payload := fmt.Sprintf("grant_type=refresh_token&refresh_token=%s",
		refreshToken.RefreshToken)

	response, err := c.accessToken(payload)
	if err != nil {
		return err
	}
	if response == nil {
		return fmt.Errorf("SchwabAPIClient: RefreshAccessToken got a bad response")
	}
	refreshToken.AccessToken = response.AccessToken
	refreshToken.ExpiresIn = response.ExpiresIn
	refreshToken.RefreshToken = response.RefreshToken
	refreshToken.IDToken = response.IDToken
	return nil
}
