package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	models "polycode-provider/client/models/item"
)

type GetItemResponse struct {
	Metadata struct{}
	Data     models.GetItemResponse
}

func (client *Client) GetItem(ID string) (*GetItemResponse, error) {
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

	return &itemResponse, nil
}

type GetItemsResponse struct {
	Metadata struct{}
	Data     []models.GetItemResponse
}

func (client *Client) GetItems(limit *int64, page *int64) (*GetItemsResponse, error) {
	url, err := url.Parse(fmt.Sprintf("%s/item", client.Host))
	if err != nil {
		return nil, err
	}

	q := url.Query()

	if limit != nil {
		q.Set("limit", fmt.Sprintf("%d", *limit))
	}
	if page != nil {
		q.Set("page", fmt.Sprintf("%d", *page))
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}

	body, err := client.fetchAPI(req, nil)
	if err != nil {
		return nil, err
	}

	itemsResponse := GetItemsResponse{}
	err = json.Unmarshal(body, &itemsResponse)
	if err != nil {
		return nil, err
	}

	return &itemsResponse, nil
}

type CreateItemResponse struct {
	GetItemResponse
}

func (client *Client) CreateItem(item models.Item) (*CreateItemResponse, error) {
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

	return &itemResponse, nil
}

type UpdateItemResponse struct {
	GetItemResponse
}

func (client *Client) UpdateItem(item models.Item) (*UpdateItemResponse, error) {
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

	return &itemResponse, nil
}

func (client *Client) DeleteItem(ID string) error {
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
