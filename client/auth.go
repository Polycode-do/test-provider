package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"polycode-provider/client/models/auth"
)

func (client *Client) Login() (*auth.AuthResponse, error) {
	if client.Auth.Username == "" || client.Auth.Password == "" {
		return nil, fmt.Errorf("empty username or password")
	}

	authReq := auth.AuthRequest{
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

	authResponse := auth.AuthResponse{}
	err = json.Unmarshal(body, &authResponse)
	if err != nil {
		return nil, err
	}

	return &authResponse, nil
}
