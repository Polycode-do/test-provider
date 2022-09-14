package client

import (
	"os"
	"polycode-provider/client/models/item"
	"testing"
)

func TestHintLifecycle(t *testing.T) {
	username := "admin@gmail.com"
	password := "12345678"
	if os.Getenv("POLYCODE_USERNAME") != "" {
		username = os.Getenv("POLYCODE_USERNAME")
	}
	if os.Getenv("POLYCODE_PASSWORD") != "" {
		password = os.Getenv("POLYCODE_PASSWORD")
	}

	c, err := NewClient(nil, &username, &password)
	if err != nil {
		t.Errorf("Error creating client: %s", err)
	}

	i := item.Item{
		Type: "hint",
		Cost: 10,
		Data: item.ItemData{
			Text: "This is a test hint",
		},
	}

	res, err := c.CreateItem(i)
	if err != nil {
		t.Errorf("Error creating hint: %s", err)
	}

	createdItem, err := c.GetItem(res.ID)
	if err != nil {
		t.Errorf("Error reading created hint: %s", err)
	}

	newItem := item.Item{
		ID:   createdItem.ID,
		Type: "hint",
		Cost: 100,
		Data: item.ItemData{
			Text: "This is an updated test hint",
		},
	}

	res, err = c.UpdateItem(newItem)
	if err != nil {
		t.Errorf("Error updating hint: %s", err)
	}

	updatedItem, err := c.GetItem(res.ID)
	if err != nil {
		t.Errorf("Error reading updated hint: %s", err)
	}

	if updatedItem.Cost != newItem.Cost {
		t.Errorf("Error checking updated hint: field Cost expected %d got %d", newItem.Cost, updatedItem.Cost)
	}
	if updatedItem.Data.Text != newItem.Data.Text {
		t.Errorf("Error checking updated hint: field Text expected '%s' got '%s'", newItem.Data.Text, updatedItem.Data.Text)
	}

	err = c.DeleteItem(updatedItem.ID)
	if err != nil {
		t.Errorf("Error deleting hint: %s", err)
	}
}
