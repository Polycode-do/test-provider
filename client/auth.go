package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"polycode-provider/client/models/auth"
)

type LoginResponse struct {
	Metadata struct{}           `json:"metadata"`
	Data     auth.LoginResponse `json:"data"`
}

// `Login` authenticates the client with the API and returns an access token.
// @returns {string} - The access token that will be used to authenticate the client.
// @returns {error} - An error if the client could not be authenticated.
func (client *Client) Login() (*string, error) {
	if client.Auth.Username == "" || client.Auth.Password == "" {
		return nil, fmt.Errorf("empty username or password")
	}

	authReq := auth.LoginRequest{
		Username:  client.Auth.Username,
		Password:  client.Auth.Password,
		GrantType: "implicit",
	}

	reqBody, err := json.Marshal(authReq)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/auth/token", client.Host), strings.NewReader(string(reqBody)))
	if err != nil {
		return nil, err
	}

	body, err := client.fetchAPI(req, nil)
	if err != nil {
		return nil, err
	}

	authResponse := LoginResponse{}
	err = json.Unmarshal(body, &authResponse)
	if err != nil {
		return nil, err
	}

	return &authResponse.Data.AccessToken, nil
}
