package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	models "polycode-provider/client/models/module"
)

type GetModuleResponse struct {
	Metadata struct{}                 `json:"metadata"`
	Data     models.GetModuleResponse `json:"data"`
}

// `GetModule` gets a module from the API.
// @param {string} ID - The ID of the module to get.
// @returns {Module} - The module that was retrieved.
// @returns {error} - An error if there was a problem getting the module.
func (c *Client) GetModule(ID string) (*models.Module, error) {
	if ID == "" {
		return nil, fmt.Errorf("empty ID")
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/module/%s", c.Host, ID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.fetchAPI(req, nil)
	if err != nil {
		return nil, err
	}

	moduleResponse := GetModuleResponse{}
	err = json.Unmarshal(body, &moduleResponse)
	if err != nil {
		return nil, err
	}

	return moduleResponse.Data.IntoModule(), nil
}

type CreateModuleResponse struct {
	Metadata struct{}                    `json:"metadata"`
	Data     models.CreateModuleResponse `json:"data"`
}

// `CreateModule` creates a module in the API.
// @param {Module} module - The module to create.
// @returns {Module} - The module that was created.
// @returns {error} - An error if there was a problem creating the module.
func (c *Client) CreateModule(module models.Module) (*models.Module, error) {
	body, err := json.Marshal(module.IntoCreateModuleRequest())
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/module", c.Host), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	body, err = c.fetchAPI(req, nil)
	if err != nil {
		return nil, err
	}

	moduleResponse := CreateModuleResponse{}
	err = json.Unmarshal(body, &moduleResponse)
	if err != nil {
		return nil, err
	}

	return moduleResponse.Data.IntoModule(), nil
}

// `UpdateModule` updates a module in the API.
// @param {Module} module - The module to update.
// @returns {Module} - The module that was updated.
// @returns {error} - An error if there was a problem updating the module.
func (c *Client) UpdateModule(module models.Module) (*models.Module, error) {
	body, err := json.Marshal(module.IntoUpdateModuleRequest())
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/module/%s", c.Host, module.ID), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	body, err = c.fetchAPI(req, nil)
	if err != nil {
		return nil, err
	}

	moduleResponse := GetModuleResponse{}
	err = json.Unmarshal(body, &moduleResponse)
	if err != nil {
		return nil, err
	}

	return moduleResponse.Data.IntoModule(), nil
}

// `DeleteModule` deletes a module from the API.
// @param {string} ID - The ID of the module to delete.
// @returns {error} - An error if there was a problem deleting the module.
func (c *Client) DeleteModule(ID string) error {
	if ID == "" {
		return fmt.Errorf("empty ID")
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/module/%s", c.Host, ID), nil)
	if err != nil {
		return err
	}

	_, err = c.fetchAPI(req, nil)
	if err != nil {
		return err
	}

	return nil
}
