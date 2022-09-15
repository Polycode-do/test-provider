package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	models "polycode-provider/client/models/item"
)

type GetItemResponse struct {
	Metadata struct{}               `json:"metadata"`
	Data     models.GetItemResponse `json:"data"`
}

// `GetItem` gets an item from the API.
// @param {string} ID - The ID of the item to get.
// @returns {Item} - The item that was retrieved.
// @returns {error} - An error if there was a problem getting the item.
func (client *Client) GetItem(ID string) (*models.Item, error) {
	if ID == "" {
		return nil, fmt.Errorf("empty ID")
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/item/%s", client.Host, ID), nil)
	if err != nil {
		return nil, err
	}

	body, err := client.fetchAPI(req, nil)
	if err != nil {
		return nil, err
	}

	itemResponse := GetItemResponse{}
	err = json.Unmarshal(body, &itemResponse)
	if err != nil {
		return nil, err
	}

	return itemResponse.Data.IntoItem(), nil
}

type CreateItemResponse struct {
	GetItemResponse
}

// `CreateItem` creates an item in the API.
// @param {Item} item - The item to create.
// @returns {Item} - The item that was created.
// @returns {error} - An error if there was a problem creating the item.
func (client *Client) CreateItem(item models.Item) (*models.Item, error) {
	body, err := json.Marshal(item.IntoCreateItemRequest())
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/item", client.Host), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	body, err = client.fetchAPI(req, nil)
	if err != nil {
		return nil, err
	}

	itemResponse := CreateItemResponse{}
	err = json.Unmarshal(body, &itemResponse)
	if err != nil {
		return nil, err
	}

	return itemResponse.Data.IntoItem(), nil
}

type UpdateItemResponse struct {
	GetItemResponse
}

// `UpdateItem` updates an item in the API.
// @param {Item} item - The item to update.
// @returns {Item} - The item that was updated.
// @returns {error} - An error if there was a problem updating the item.
func (client *Client) UpdateItem(item models.Item) (*models.Item, error) {
	if item.ID == "" {
		return nil, fmt.Errorf("empty ID")
	}

	body, err := json.Marshal(item.IntoUpdateItemRequest())
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/item/%s", client.Host, item.ID), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	body, err = client.fetchAPI(req, nil)
	if err != nil {
		return nil, err
	}

	itemResponse := UpdateItemResponse{}
	err = json.Unmarshal(body, &itemResponse)
	if err != nil {
		return nil, err
	}

	return itemResponse.Data.IntoItem(), nil
}

// `DeleteItem` deletes an item from the API.
// @param {string} ID - The ID of the item to delete.
// @returns {error} - An error if there was a problem deleting the item.
func (client *Client) DeleteItem(ID string) error {
	if ID == "" {
		return fmt.Errorf("empty ID")
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/item/%s", client.Host, ID), nil)
	if err != nil {
		return err
	}

	_, err = client.fetchAPI(req, nil)
	if err != nil {
		return err
	}

	return nil
}
