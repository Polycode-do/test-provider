package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"polycode-provider/client/models/user"
)

func (client *Client) GetUser(ID string) (*user.GetUserResponse, error) {
	if ID == "" {
		return nil, fmt.Errorf("empty ID")
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/user/%s", client.Host, ID), nil)
	if err != nil {
		return nil, err
	}

	body, err := client.fetchAPI(req, nil)
	if err != nil {
		return nil, err
	}

	userResponse := user.GetUserResponse{}
	err = json.Unmarshal(body, &userResponse)
	if err != nil {
		return nil, err
	}

	return &userResponse, nil
}
