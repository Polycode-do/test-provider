package content

import "polycode-provider/client/shared"

// `GetContentResponse` is the response body of the get content endpoint.
// @property {string} ID - The unique identifier for the content.
// @property {string} Name - The name of the content.
// @property {string} Description - A description of the content.
// @property {string} Type - The type of content. Only `exercise` is available at the moment.
// @property {int64} Reward - The amount of points the user will receive for completing this content.
// @property {GetComponentResponse} RootComponent - This is the root component of the content.
// @property Data - This is the data of the content.
type GetContentResponse struct {
	ID            string               `json:"id"`
	Name          string               `json:"name"`
	Description   string               `json:"description"`
	Type          string               `json:"type"`
	Reward        int64                `json:"reward"`
	RootComponent GetComponentResponse `json:"rootComponent"`
	Data          struct{}             `json:"data"`
}

// `IntoContent` converts the response body into a pointer of a `Content` struct.
// @returns {Content} The content.
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

// `GetComponentResponse` is the response body holding information about a component.
// @property {string} ID - The unique identifier for the component.
// @property {string} Type - The type of component. Can be one of the following:
// `markdown`, `editor`, `container`
// @property {GetComponentResponseData} Data - This is the data of the component.
type GetComponentResponse struct {
	ID   string                   `json:"id"`
	Type string                   `json:"type"`
	Data GetComponentResponseData `json:"data"`
}

// `GetComponentResponseData` is the data of the component,
// this structure holds pointers because all the properties are optional.
// @property {*GetComponentResponse[]} Components - An array of nested components. This is only used if the type is `container`.
// @property {*string} Markdown - The markdown of the component. This is only used if the type is `markdown`.
// @property {*string} Items - The items of the component. This is only used if the type is `editor`.
// @property {*GetValidatorResponse[]} Validators - The validators of the component. This is only used if the type is `editor`.
// @property {*GetEditorSettingsResponse} EditorSettings - The editor settings of the component. This is only used if the type is `editor`.
// @property {*string} Orientation - The orientation of the component. This is only used if the type is `container`.
type GetComponentResponseData struct {
	Components     *[]GetComponentResponse    `json:"components"`
	Markdown       *string                    `json:"markdown"`
	Items          *[]string                  `json:"items"`
	Validators     *[]GetValidatorResponse    `json:"validators"`
	EditorSettings *GetEditorSettingsResponse `json:"editorSettings"`
	Orientation    *string                    `json:"orientation"`
}

// `IntoComponents` converts the response body into an array of `Component` struct.
// @returns {Component} The components.
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

// `IntoItemsIdentifier` converts the response body into an array of `ItemIdentifier` struct.
// @returns {ItemIdentifier} The item identifiers.
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

// `IntoValidators` converts the response body into an array of `Validator` struct.
// @returns {Validator} The validators.
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

// `IntoEditorSettings` converts the response body into an `EditorSettings` struct.
// @returns {EditorSettings} The editor settings.
func (rd *GetComponentResponseData) IntoEditorSettings() EditorSettings {
	result := EditorSettings{}

	if rd.EditorSettings != nil {
		for _, language := range rd.EditorSettings.Languages {
			result.Languages = append(result.Languages, Language(language))
		}
	}

	return result
}

// `GetItemResponse` is the response body holding information about an item.
// @property {string} ID - The ID of the validator.
// @property {bool} IsHidden - If true, the validator will not be shown to the user.
// @property {GetValidatorInputResponse} Input - The input that the validator will receive.
// @property {GetValidatorOutputResponse} Expected - The expected output of the validator.
type GetValidatorResponse struct {
	ID       string                     `json:"id"`
	IsHidden bool                       `json:"isHidden"`
	Input    GetValidatorInputResponse  `json:"input"`
	Expected GetValidatorOutputResponse `json:"expected"`
}

// `GetValidatorInputResponse` is the input that the validator will receive.
// @property {[]string} Stdin - An array of stdin.
type GetValidatorInputResponse struct {
	Stdin []string `json:"stdin"`
}

// `GetValidatorOutputResponse` is the expected output of the validator.
// @property {[]string} Stdout - An array of stdout.
type GetValidatorOutputResponse struct {
	Stdout []string `json:"stdout"`
}

// `GetEditorSettingsResponse` is the response body holding information about the editor settings of the component.
// @property {string[]} Languages - An array of available languages for the editor.
type GetEditorSettingsResponse struct {
	Languages []GetLanguageResponse `json:"languages"`
}

// `GetLanguageResponse` is the response body holding information about a language settings.
// @property {string} DefaultCode - The default language code for the language.
// @property {string} Language - The language name.
// @property {string} Version - The version of the language.
type GetLanguageResponse struct {
	DefaultCode string `json:"defaultCode"`
	Language    string `json:"language"`
	Version     string `json:"version"`
}

// `CreateContentResponse` is the response body of the create content endpoint.
// @property {string} ID - The unique identifier for the content.
// @property {string} Name - The name of the content.
// @property {string} Description - A description of the content.
// @property {string} Type - The type of content. Only `exercise` is available at the moment.
// @property {int64} Reward - The amount of points the user will receive for completing this content.
// @property {GetComponentResponse} RootComponent - This is the root component of the content.
// @property Data - This is the data of the content.
type CreateContentResponse struct {
	GetContentResponse
}

// `UpdateContentResponse` is the response body of the update content endpoint.
// @property {string} ID - The unique identifier for the content.
// @property {string} Name - The name of the content.
// @property {string} Description - A description of the content.
// @property {string} Type - The type of content. Only `exercise` is available at the moment.
// @property {int64} Reward - The amount of points the user will receive for completing this content.
// @property {GetComponentResponse} RootComponent - This is the root component of the content.
// @property Data - This is the data of the content.
type UpdateContentResponse struct {
	GetContentResponse
}
