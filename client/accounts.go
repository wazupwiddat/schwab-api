package client

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/wazupwiddat/schwab-api/models"
)

// accounts/accountNumbers
func (c SchwabAPIClient) GetAccountNumbers(accessToken string) (models.AccountNumbers, error) {
	url := fmt.Sprintf("%s/%s", c.TraderBaseURL, "accounts/accountNumbers")
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Set the Content-Type header to application/json
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	request.Header.Set("Accept", "application/json")

	resp, err := c.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 400 {
		log.Println("SchwabAPIClient:", "Failed to get accounts/accountNumbers")
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println(string(body))
			return nil, err
		}
		return nil, err
	}

	var acctNumbers models.AccountNumbers
	if err := json.NewDecoder(resp.Body).Decode(&acctNumbers); err != nil {
		return nil, err
	}

	return acctNumbers, nil
}

// accounts
func (c SchwabAPIClient) GetAccounts(accessToken string, includePositions bool) (models.SecuritiesAccounts, error) {
	url := fmt.Sprintf("%s/%s", c.TraderBaseURL, "accounts")
	if includePositions {
		url = fmt.Sprintf("%s/%s?%s", c.TraderBaseURL, "accounts", "fields=positions")
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// Set the Content-Type header to application/json
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	request.Header.Set("Accept", "application/json")

	resp, err := c.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 400 {
		log.Println("SchwabAPIClient:", "Failed to get accounts")
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println(string(body))
			return nil, err
		}
		return nil, err
	}

	var accounts models.SecuritiesAccounts
	if err := json.NewDecoder(resp.Body).Decode(&accounts); err != nil {
		return nil, err
	}

	return accounts, nil
}
