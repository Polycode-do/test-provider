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
