package client

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"polycode-provider/client/models/auth"
)

// `Client` is a struct that holds all information about the client that will interact with the API.
// @property {string} Host - The hostname of the API server.
// @property HTTPClient - The HTTP client to use for making requests.
// @property {string} AccessToken - The access token that will be used to authenticate the client.
// @property Auth - This is the authentication credentials that will be used to authenticate the
// client.
type Client struct {
	Host        string
	HTTPClient  *http.Client
	AccessToken string
	Auth        auth.Credentials
}

// `NewClient` creates a new client for interacting with the API
// @param {string} host - The hostname of the API server.
// @param {string} username - The username to use for authentication.
// @param {string} password - The password to use for authentication.
// @returns {Client} - The client that will be used to interact with the API.
// @returns {error} - An error if the client could not be created.
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

	token, err := client.Login()
	if err != nil {
		return nil, err
	}
	if token == nil {
		return nil, fmt.Errorf("access token nil")
	}

	client.AccessToken = *token

	return &client, nil
}

// `fetchAPI` is a function that is used to make requests to the API adding every needed headers to the request.
// @param {http.Request} req - The request to make.
// @param {string} authToken - The access token to use for authentication.
// @returns {[]byte} - The response body.
// @returns {error} - An error if the request could not be made.
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
