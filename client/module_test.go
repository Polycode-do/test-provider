package client

import (
	"fmt"
	"os"
	"polycode-provider/client/models/content"
	"polycode-provider/client/models/module"
	"testing"
)

func TestModuleLifecycle(t *testing.T) {
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

	m := module.Module{
		Name:        "Test",
		Description: "This is a test module.",
		Reward:      10,
		Type:        "challenge",
		Tags:        []string{"test"},
		Data:        module.ModuleData{},
		Modules:     []module.ModuleIdentifier{},
		Contents:    []module.ContentIdentifier{},
	}

	res, err := c.CreateModule(m)
	if err != nil {
		t.Errorf("Error creating module: %s", err)
	}

	t.Logf("%+v", res)

	createdModule, err := c.GetModule(res.ID)
	if err != nil {
		t.Errorf("Error reading created module: %s", err)
	}

	newModule := module.Module{
		ID:          createdModule.ID,
		Name:        "Test",
		Description: "This is an updated test module.",
		Reward:      20,
		Type:        "challenge",
		Tags:        []string{"test", "updated"},
		Data:        module.ModuleData{},
		Modules:     []module.ModuleIdentifier{},
		Contents:    []module.ContentIdentifier{},
	}

	res, err = c.UpdateModule(newModule)
	if err != nil {
		t.Errorf("Error updating module: %s", err)
	}

	updatedModule, err := c.GetModule(res.ID)
	if err != nil {
		t.Errorf("Error reading updated module: %s", err)
	}

	if updatedModule.Description != newModule.Description {
		t.Errorf("Error checking updated module: field Description expected %s got %s", newModule.Description, updatedModule.Description)
	}
	if updatedModule.Reward != newModule.Reward {
		t.Errorf("Error checking updated module: field Reward expected %d got %d", newModule.Reward, updatedModule.Reward)
	}
	if updatedModule.Tags[1] != newModule.Tags[1] {
		t.Errorf("Error checking updated module: field Tags expected %s got %s", newModule.Tags[1], updatedModule.Tags[1])
	}

	err = c.DeleteModule(updatedModule.ID)
	if err != nil {
		t.Errorf("Error deleting module: %s", err)
	}
}

func TestNestedModules(t *testing.T) {
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

	m := module.Module{
		Name:        "Test",
		Description: "This is a nested test module.",
		Reward:      10,
		Type:        "submodule",
		Tags:        []string{"test"},
		Data:        module.ModuleData{},
		Modules:     []module.ModuleIdentifier{},
		Contents:    []module.ContentIdentifier{},
	}

	res1, err := c.CreateModule(m)
	if err != nil {
		t.Errorf("Error creating module: %s", err)
	}

	createdModule1, err := c.GetModule(res1.ID)
	if err != nil {
		t.Errorf("Error reading created module: %s", err)
	}

	co := content.Content{
		Name:        "Test content",
		Description: "This is a test content",
		Type:        "exercise",
		Reward:      10,
		Data:        content.ContentData{},
		RootComponent: content.Component{
			Type:        "container",
			Orientation: "vertical",
			Data: content.ComponentData{
				Components: []content.Component{},
			},
		},
	}

	res2, err := c.CreateContent(co)
	if err != nil {
		t.Errorf("Error creating content: %s", err)
	}

	createdContent, err := c.GetContent(res2.ID)
	if err != nil {
		t.Errorf("Error reading created content: %s", err)
	}

	m = module.Module{
		Name:        "Test",
		Description: "This is a test module.",
		Reward:      10,
		Type:        "challenge",
		Tags:        []string{"test"},
		Data:        module.ModuleData{},
		Modules:     []module.ModuleIdentifier{{ID: createdModule1.ID}},
		Contents:    []module.ContentIdentifier{{ID: createdContent.ID}},
	}

	res1, err = c.CreateModule(m)
	if err != nil {
		t.Errorf("Error creating module: %s", err)
	}

	fmt.Sprintln(res1)

	createdModule2, err := c.GetModule(res1.ID)
	if err != nil {
		t.Errorf("Error reading created module: %s", err)
	}

	if createdModule2.Modules[0].ID != createdModule1.ID {
		t.Errorf("Error checking created module: field Modules expected %s got %s", createdModule1.ID, createdModule2.Modules[0].ID)
	}

	newModule := module.Module{
		ID:          createdModule2.ID,
		Name:        "Test",
		Description: "This is an updated test module.",
		Reward:      20,
		Type:        "challenge",
		Tags:        []string{"test", "updated"},
		Data:        module.ModuleData{},
		Modules:     []module.ModuleIdentifier{{ID: createdModule1.ID}},
		Contents:    []module.ContentIdentifier{},
	}

	res1, err = c.UpdateModule(newModule)
	if err != nil {
		t.Errorf("Error updating module: %s", err)
	}

	updatedModule, err := c.GetModule(res1.ID)
	if err != nil {
		t.Errorf("Error reading updated module: %s", err)
	}

	if updatedModule.Description != newModule.Description {
		t.Errorf("Error checking updated module: field Description expected %s got %s", newModule.Description, updatedModule.Description)
	}
	if updatedModule.Reward != newModule.Reward {
		t.Errorf("Error checking updated module: field Reward expected %d got %d", newModule.Reward, updatedModule.Reward)
	}
	if updatedModule.Tags[1] != newModule.Tags[1] {
		t.Errorf("Error checking updated module: field Tags expected %s got %s", newModule.Tags[1], updatedModule.Tags[1])
	}
	if len(updatedModule.Contents) > 0 {
		t.Errorf("Error checking updated module: field Contents expected %s got %s", "nothing", updatedModule.Contents)
	}

	err = c.DeleteModule(updatedModule.ID)
	if err != nil {
		t.Errorf("Error deleting module: %s", err)
	}

	err = c.DeleteContent(createdContent.ID)
	if err != nil {
		t.Errorf("Error deleting content: %s", err)
	}

	err = c.DeleteModule(createdModule1.ID)
	if err != nil {
		t.Errorf("Error deleting module: %s", err)
	}
}
