package content

import "polycode-provider/client/shared"

type GetContentResponse struct {
	ID            string               `json:"id"`
	Name          string               `json:"name"`
	Description   string               `json:"description"`
	Type          string               `json:"type"`
	Reward        int64                `json:"reward"`
	RootComponent GetComponentResponse `json:"rootComponent"`
	Data          struct{}             `json:"data"`
}

func (cr *GetContentResponse) IntoContent() *Content {
	return &Content{
		ID:          cr.ID,
		Name:        cr.Name,
		Description: cr.Description,
		Type:        cr.Type,
		Reward:      cr.Reward,
		RootComponent: Component{
			ID:   cr.RootComponent.ID,
			Type: cr.RootComponent.Type,
			Data: ComponentData{
				Components:     cr.RootComponent.Data.IntoComponents(),
				Markdown:       shared.ConvertNilStringPointer(cr.RootComponent.Data.Markdown),
				Items:          cr.RootComponent.Data.IntoItemsIdentifier(),
				Validators:     cr.RootComponent.Data.IntoValidators(),
				EditorSettings: cr.RootComponent.Data.IntoEditorSettings(),
			},
			Orientation: shared.ConvertNilStringPointer(cr.RootComponent.Data.Orientation),
		},
		Data: ContentData{},
	}
}

type GetComponentResponse struct {
	ID   string                   `json:"id"`
	Type string                   `json:"type"`
	Data GetComponentResponseData `json:"data"`
}

type GetComponentResponseData struct {
	Components     *[]GetComponentResponse    `json:"components"`
	Markdown       *string                    `json:"markdown"`
	Items          *[]string                  `json:"items"`
	Validators     *[]GetValidatorResponse    `json:"validators"`
	EditorSettings *GetEditorSettingsResponse `json:"editorSettings"`
	Orientation    *string                    `json:"orientation"`
}

func (rd *GetComponentResponseData) IntoComponents() []Component {
	components := make([]Component, 0)

	if rd.Components != nil {
		for _, component := range *rd.Components {
			components = append(components, Component{
				ID:   component.ID,
				Type: component.Type,
				Data: ComponentData{
					Components:     component.Data.IntoComponents(),
					Markdown:       shared.ConvertNilStringPointer(component.Data.Markdown),
					Items:          component.Data.IntoItemsIdentifier(),
					Validators:     component.Data.IntoValidators(),
					EditorSettings: component.Data.IntoEditorSettings(),
				},
				Orientation: shared.ConvertNilStringPointer(component.Data.Orientation),
			})
		}
	}

	return components
}

func (rd *GetComponentResponseData) IntoItemsIdentifier() []ItemIdentifier {
	items := make([]ItemIdentifier, 0)

	if rd.Items != nil {
		for _, item := range *rd.Items {
			items = append(items, ItemIdentifier{
				ID: item,
			})
		}
	}

	return items
}

func (rd *GetComponentResponseData) IntoValidators() []Validator {
	validators := make([]Validator, 0)

	if rd.Validators != nil {
		for _, validator := range *rd.Validators {
			validators = append(validators, Validator{
				ID:       validator.ID,
				IsHidden: validator.IsHidden,
				Input: ValidatorInput{
					Stdin: validator.Input.Stdin,
				},
				Output: ValidatorOutput{
					Stdout: validator.Expected.Stdout,
				},
			})
		}
	}

	return validators
}

func (rd *GetComponentResponseData) IntoEditorSettings() EditorSettings {
	result := EditorSettings{}

	if rd.EditorSettings != nil {
		for _, language := range rd.EditorSettings.Languages {
			result.Languages = append(result.Languages, Language(language))
		}
	}

	return result
}

type GetValidatorResponse struct {
	ID       string                     `json:"id"`
	IsHidden bool                       `json:"isHidden"`
	Input    GetValidatorInputResponse  `json:"input"`
	Expected GetValidatorOutputResponse `json:"expected"`
}

type GetValidatorInputResponse struct {
	Stdin []string `json:"stdin"`
}

type GetValidatorOutputResponse struct {
	Stdout []string `json:"stdout"`
}

type GetEditorSettingsResponse struct {
	Languages []GetLanguageResponse `json:"languages"`
}

type GetLanguageResponse struct {
	DefaultCode string `json:"defaultCode"`
	Language    string `json:"language"`
	Version     string `json:"version"`
}

type CreateContentResponse struct {
	GetContentResponse
}

type UpdateContentResponse struct {
	GetContentResponse
}
