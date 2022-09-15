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

// `GetContent` gets a content from the API.
// @param {string} ID - The ID of the content to get.
// @returns {Content} - The content that was retrieved.
// @returns {error} - An error if there was a problem getting the content.
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

// `CreateContent` creates a content in the API.
// @param {Content} content - The content to create.
// @returns {Content} - The content that was created.
// @returns {error} - An error if there was a problem creating the content.
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

// `UpdateContent` updates a content in the API.
// @param {Content} content - The content to update.
// @returns {Content} - The content that was updated.
// @returns {error} - An error if there was a problem updating the content.
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

// `DeleteContent` deletes a content from the API.
// @param {string} ID - The ID of the content to delete.
// @returns {error} - An error if there was a problem deleting the content.
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
