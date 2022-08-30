package client

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"polycode-provider/client/models/auth"
)

type Client struct {
	Host        string
	HTTPClient  *http.Client
	AccessToken string
	Auth        auth.Credentials
}

func NewClient(host, username, password *string) (*Client, error) {
	defaultHost := "http://localhost:3000"
	if os.Getenv("POLYCODE_HOST") != "" {
		defaultHost = os.Getenv("POLYCODE_HOST")
	}

	client := Client{
		Host:       defaultHost,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
	}

	if host != nil {
		client.Host = *host
	}

	if username == nil || password == nil {
		return &client, nil
	}

	client.Auth = auth.Credentials{
		Username: *username,
		Password: *password,
	}

	authResponse, err := client.Login()
	if err != nil {
		return nil, err
	}

	client.AccessToken = authResponse.Data.AccessToken

	return &client, nil
}

func (client *Client) fetchAPI(req *http.Request, authToken *string) ([]byte, error) {
	token := client.AccessToken

	if authToken != nil {
		token = *authToken
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Content-Type", "application/json")
	res, err := client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated && res.StatusCode != http.StatusNoContent {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, string(body))
	}

	return body, err
}
