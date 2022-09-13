package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	models "polycode-provider/client/models/content"
)

type GetContentResponse struct {
	Metadata struct{}                  `json:"metadata"`
	Data     models.GetContentResponse `json:"data"`
}

func (client *Client) GetContent(ID string) (*models.Content, error) {
	if ID == "" {
		return nil, fmt.Errorf("empty ID")
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/content/%s", client.Host, ID), nil)
	if err != nil {
		return nil, err
	}

	body, err := client.fetchAPI(req, nil)
	if err != nil {
		return nil, err
	}

	contentResponse := GetContentResponse{}
	err = json.Unmarshal(body, &contentResponse)
	if err != nil {
		return nil, err
	}

	return contentResponse.Data.IntoContent(), nil
}

type CreateContentResponse struct {
	GetContentResponse
}

func (client *Client) CreateContent(content models.Content) (*models.Content, error) {
	body, err := json.Marshal(content.IntoCreateContentRequest())
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/content", client.Host), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	body, err = client.fetchAPI(req, nil)
	if err != nil {
		return nil, err
	}

	contentResponse := CreateContentResponse{}
	err = json.Unmarshal(body, &contentResponse)
	if err != nil {
		return nil, err
	}

	return contentResponse.Data.IntoContent(), nil
}

type UpdateContentResponse struct {
	GetContentResponse
}

func (client *Client) UpdateContent(content models.Content) (*models.Content, error) {
	if content.ID == "" {
		return nil, fmt.Errorf("empty ID")
	}

	body, err := json.Marshal(content.IntoUpdateContentRequest())
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/content/%s", client.Host, content.ID), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	body, err = client.fetchAPI(req, nil)
	if err != nil {
		return nil, err
	}

	contentResponse := UpdateContentResponse{}
	err = json.Unmarshal(body, &contentResponse)
	if err != nil {
		return nil, err
	}

	return contentResponse.Data.IntoContent(), nil
}

func (client *Client) DeleteContent(ID string) error {
	if ID == "" {
		return fmt.Errorf("empty ID")
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/content/%s", client.Host, ID), nil)
	if err != nil {
		return err
	}

	_, err = client.fetchAPI(req, nil)
	if err != nil {
		return err
	}

	return nil
}
