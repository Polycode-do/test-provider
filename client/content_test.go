package client

import (
	"os"
	"polycode-provider/client/models/content"
	"polycode-provider/client/models/item"
	"testing"
)

func TestExerciseLifecycle(t *testing.T) {
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

	item, err := c.CreateItem(i)
	if err != nil {
		t.Errorf("Error creating hint: %s", err)
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
				Components: []content.Component{
					{
						Type: "markdown",
						Data: content.ComponentData{
							Markdown: "# This is a markdown component",
						},
					},
					{
						Type:        "container",
						Orientation: "vertical",
						Data: content.ComponentData{
							Components: []content.Component{
								{
									Type: "markdown",
									Data: content.ComponentData{
										Markdown: "# This is a nested markdown component",
									},
								},
								{
									Type: "editor",
									Data: content.ComponentData{
										EditorSettings: content.EditorSettings{
											Languages: []content.Language{
												{
													DefaultCode: "print('Hello world')",
													Language:    "PYTHON",
												},
											},
										},
										Validators: []content.Validator{
											{
												Input: content.ValidatorInput{
													Stdin: []string{},
												},
												Output: content.ValidatorOutput{
													Stdout: []string{"test output"},
												},
											},
										},
										Items: []content.ItemIdentifier{
											{
												ID: item.ID,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	res, err := c.CreateContent(co)
	if err != nil {
		t.Errorf("Error creating content: %s", err)
	}

	createdContent, err := c.GetContent(res.ID)
	if err != nil {
		t.Errorf("Error reading created content: %s", err)
	}

	newContent := content.Content{
		ID:          createdContent.ID,
		Name:        "Test content",
		Description: "This is an updated test content",
		Type:        "exercise",
		Reward:      100,
		Data:        content.ContentData{},
		RootComponent: content.Component{
			ID:          createdContent.RootComponent.ID,
			Type:        "container",
			Orientation: "vertical",
			Data: content.ComponentData{
				Components: []content.Component{
					{
						ID:   createdContent.RootComponent.Data.Components[0].ID,
						Type: "markdown",
						Data: content.ComponentData{
							Markdown: "# This is an updated markdown component",
						},
					},
					{
						ID:          createdContent.RootComponent.Data.Components[1].ID,
						Type:        "container",
						Orientation: "vertical",
						Data: content.ComponentData{
							Components: []content.Component{
								{
									ID:   createdContent.RootComponent.Data.Components[1].Data.Components[0].ID,
									Type: "markdown",
									Data: content.ComponentData{
										Markdown: "# This is an updated nested markdown component",
									},
								},
								{
									ID:   createdContent.RootComponent.Data.Components[1].Data.Components[1].ID,
									Type: "editor",
									Data: content.ComponentData{
										EditorSettings: content.EditorSettings{
											Languages: []content.Language{
												{
													DefaultCode: "print('Hello world')",
													Language:    "PYTHON",
												},
											},
										},
										Validators: []content.Validator{
											{
												ID: createdContent.RootComponent.Data.Components[1].Data.Components[1].Data.Validators[0].ID,
												Input: content.ValidatorInput{
													Stdin: []string{},
												},
												Output: content.ValidatorOutput{
													Stdout: []string{"test output"},
												},
											},
										},
										Items: []content.ItemIdentifier{
											{
												ID: item.ID,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	res, err = c.UpdateContent(newContent)
	if err != nil {
		t.Errorf("Error updating content: %s", err)
	}

	updatedContent, err := c.GetContent(res.ID)
	if err != nil {
		t.Errorf("Error reading updated content: %s", err)
	}

	if updatedContent.Description != newContent.Description {
		t.Errorf("Error checking updated content: field Description expected '%s' got '%s'", newContent.Description, updatedContent.Description)
	}
	if updatedContent.Reward != newContent.Reward {
		t.Errorf("Error checking updated content: field Reward expected %d got %d", newContent.Reward, updatedContent.Reward)
	}
	if updatedContent.RootComponent.Data.Components[0].Data.Markdown != newContent.RootComponent.Data.Components[0].Data.Markdown {
		t.Errorf("Error checking updated content: field RootComponent.Data.Components[0].Data.Markdown expected %s got %s", newContent.RootComponent.Data.Components[0].Data.Markdown, updatedContent.RootComponent.Data.Components[0].Data.Markdown)
	}
	if updatedContent.RootComponent.Data.Components[1].Data.Components[0].Data.Markdown != newContent.RootComponent.Data.Components[1].Data.Components[0].Data.Markdown {
		t.Errorf("Error checking updated content: field RootComponent.Data.Components[1].Data.Components[0].Data.Markdown expected %s got %s", newContent.RootComponent.Data.Components[1].Data.Components[0].Data.Markdown, updatedContent.RootComponent.Data.Components[1].Data.Components[0].Data.Markdown)
	}

	err = c.DeleteContent(updatedContent.ID)
	if err != nil {
		t.Errorf("Error deleting content: %s", err)
	}
}
