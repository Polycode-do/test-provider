package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	models "polycode-provider/client/models/content"
)

type GetContentResponse struct {
	Metadata struct{}
	Data     models.GetContentResponse
}

func (client *Client) GetContent(ID string) (*GetContentResponse, error) {
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

	return &contentResponse, nil
}

type GetContentsResponse struct {
	Metadata struct{}
	Data     []models.GetContentResponse
}

type ContentOrderBy struct {
	Name   *int8
	Reward *int8
}

func (client *Client) GetContents(limit *int64, page *int64) (*GetContentsResponse, error) {
	url, err := url.Parse(fmt.Sprintf("%s/content", client.Host))
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

	contentResponse := GetContentsResponse{}
	err = json.Unmarshal(body, &contentResponse)
	if err != nil {
		return nil, err
	}

	return &contentResponse, nil
}

type CreateContentResponse struct {
	GetContentResponse
}

func (client *Client) CreateContent(content models.Content) (*CreateContentResponse, error) {
	fmt.Printf("%+v\n", content.RootComponent.Data.Components)

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

	return &contentResponse, nil
}

type UpdateContentResponse struct {
	GetContentResponse
}

func (client *Client) UpdateContent(content models.Content) (*UpdateContentResponse, error) {
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

	return &contentResponse, nil
}

func (client *Client) DeleteContent(ID string) error {
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
