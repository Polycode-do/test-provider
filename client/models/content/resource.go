package content

import "polycode-provider/client/shared"

// `Content` is the main model for the content system. It contains all the information about a content.
// @property {string} ID - The unique identifier for the content.
// @property {string} Name - The name of the content.
// @property {string} Description - A description of the content.
// @property {string} Type - The type of content. Only `exercise` is available at the moment.
// @property {int64} Reward - The amount of points the user will receive for completing this content.
// @property {Component} RootComponent - This is the root component of the content.
// @property {ContentData} Data - This is the data of the content.
type Content struct {
	ID            string
	Name          string
	Description   string
	Type          string
	Reward        int64
	RootComponent Component
	Data          ContentData
}

// `IntoCreateContentRequest` converts the content into a `CreateContentRequest`.
// @returns {CreateContentRequest} The converted content.
func (content *Content) IntoCreateContentRequest() CreateContentRequest {
	return CreateContentRequest{
		Name:        content.Name,
		Description: content.Description,
		Type:        content.Type,
		Reward:      content.Reward,
		RootComponent: CreateComponentRequest{
			Type: content.RootComponent.Type,
			Data: CreateComponentRequestData{
				Components:     content.RootComponent.Data.IntoCreateComponentRequest(),
				Markdown:       shared.ConvertNilString(content.RootComponent.Data.Markdown),
				Items:          content.RootComponent.Data.FlattenItemIdentifiers(),
				Validators:     content.RootComponent.Data.IntoCreateValidatorRequest(),
				EditorSettings: content.RootComponent.Data.IntoCreateEditorSettingsRequest(),
				Orientation:    shared.ConvertNilString(content.RootComponent.Orientation),
			},
		},
		Data: CreateContentRequestData{},
	}
}

// `IntoUpdateContentRequest` converts the content into a `UpdateContentRequest`.
// @returns {UpdateContentRequest} The converted content.
func (content Content) IntoUpdateContentRequest() UpdateContentRequest {
	return UpdateContentRequest{
		CreateContentRequest: CreateContentRequest{
			Name:        content.Name,
			Description: content.Description,
			Type:        content.Type,
			Reward:      content.Reward,
		},
		RootComponent: UpdateComponentRequest{
			ID: content.RootComponent.ID,
			CreateComponentRequest: CreateComponentRequest{
				Type: content.RootComponent.Type,
			},
			Data: UpdateComponentRequestData{
				Components: content.RootComponent.Data.IntoUpdateComponentRequest(),
				Validators: content.RootComponent.Data.IntoUpdateValidatorRequest(),
				CreateComponentRequestData: CreateComponentRequestData{
					Markdown:       shared.ConvertNilString(content.RootComponent.Data.Markdown),
					Items:          content.RootComponent.Data.FlattenItemIdentifiers(),
					EditorSettings: content.RootComponent.Data.IntoCreateEditorSettingsRequest(),
					Orientation:    shared.ConvertNilString(content.RootComponent.Orientation),
				},
			},
		},
		Data: UpdateContentRequestData{},
	}
}

// `ContentData` is the data of the content.
// for now it is empty but it will be used in the future
type ContentData struct{}

// `Component` is a component of the content.
// @property {string} ID - The ID of the component.
// @property {string} Type - The type of component. Can be one of the following :
// `markdown`, `editor`, `container`
// @property {string} Orientation - The orientation of the component. This can be either "horizontal"
// or "vertical" (only if type = container).
// @property {ComponentData} Data - This is the data of the component.
type Component struct {
	ID          string
	Type        string
	Orientation string
	Data        ComponentData
}

// `ComponentData` is the data of the component.
// @property {Component[]} Components - An array of nested components. This is only used if the type is `container`.
// @property {string} Markdown - The markdown of the component. This is only used if the type is `markdown`.
// @property {ItemIdentifier[]} Items - An array of related items. This is only used if the type is `editor`.
// @property {Validator[]} Validators - The validators of the component. This is only used if the type is `editor`.
// @property {EditorSettings} EditorSettings - The editor settings of the component. This is only used if the type is `editor`.
type ComponentData struct {
	Components     []Component
	Markdown       string
	Items          []ItemIdentifier
	Validators     []Validator
	EditorSettings EditorSettings
}

// `IntoUpdateComponentRequest` converts the component data into a pointer of an array of `UpdateComponentRequest`.
// @returns {*UpdateComponentRequest[]} The converted component data.
func (componentData *ComponentData) IntoCreateComponentRequest() *[]CreateComponentRequest {
	result := make([]CreateComponentRequest, 0)

	if componentData.Components == nil {
		return nil
	}

	for _, component := range componentData.Components {
		result = append(result, CreateComponentRequest{
			Type: component.Type,
			Data: CreateComponentRequestData{
				Components:     component.Data.IntoCreateComponentRequest(),
				Markdown:       shared.ConvertNilString(component.Data.Markdown),
				Items:          component.Data.FlattenItemIdentifiers(),
				Validators:     component.Data.IntoCreateValidatorRequest(),
				EditorSettings: component.Data.IntoCreateEditorSettingsRequest(),
				Orientation:    shared.ConvertNilString(component.Orientation),
			},
		})
	}

	return &result
}

// `IntoUpdateComponentRequest` converts the component data into a pointer of an array of `UpdateComponentRequest`.
// @returns {*UpdateComponentRequest[]} The converted component data.
func (componentData *ComponentData) IntoUpdateComponentRequest() *[]UpdateComponentRequest {
	result := make([]UpdateComponentRequest, 0)

	if componentData.Components == nil {
		return nil
	}

	for _, component := range componentData.Components {
		result = append(result, UpdateComponentRequest{
			ID: component.ID,
			CreateComponentRequest: CreateComponentRequest{
				Type: component.Type,
			},
			Data: UpdateComponentRequestData{
				Components: component.Data.IntoUpdateComponentRequest(),
				Validators: component.Data.IntoUpdateValidatorRequest(),
				CreateComponentRequestData: CreateComponentRequestData{
					Markdown:       shared.ConvertNilString(component.Data.Markdown),
					Items:          component.Data.FlattenItemIdentifiers(),
					EditorSettings: component.Data.IntoCreateEditorSettingsRequest(),
					Orientation:    shared.ConvertNilString(component.Orientation),
				},
			},
		})
	}

	return &result
}

// `IntoCreateValidatorRequest` converts the component data into a pointer of an array of `CreateValidatorRequest`.
// @returns {*CreateValidatorRequest[]} The converted component data.
func (componentData *ComponentData) IntoCreateValidatorRequest() *[]CreateValidatorRequest {
	result := make([]CreateValidatorRequest, 0)

	if componentData.Validators == nil {
		return nil
	}

	for _, validator := range componentData.Validators {
		result = append(result, CreateValidatorRequest{
			IsHidden: validator.IsHidden,
			Input:    CreateValidatorInputRequest(validator.Input),
			Expected: CreateValidatorOutputRequest(validator.Output),
		})
	}

	return &result
}

// `IntoCreateEditorSettingsRequest` converts the component data into a pointer of an array of `CreateEditorSettingsRequest`.
// @returns {*CreateEditorSettingsRequest[]} The converted component data.
func (componentData *ComponentData) IntoCreateEditorSettingsRequest() *CreateEditorSettingsRequest {
	result := CreateEditorSettingsRequest{}

	if componentData.EditorSettings.Languages == nil {
		return nil
	}

	for _, language := range componentData.EditorSettings.Languages {
		result.Languages = append(result.Languages, CreateLanguageRequest(language))
	}

	return &result
}
// `IntoUpdateValidatorRequest` converts the component data into a pointer of an array of `UpdateValidatorRequest`.
// @returns {*UpdateValidatorRequest[]} The converted component data.
func (componentData *ComponentData) IntoUpdateValidatorRequest() *[]UpdateValidatorRequest {
	result := make([]UpdateValidatorRequest, 0)

	if componentData.Validators == nil {
		return nil
	}

	for _, validator := range componentData.Validators {
		result = append(result, UpdateValidatorRequest{
			ID: validator.ID,
			CreateValidatorRequest: CreateValidatorRequest{
				IsHidden: validator.IsHidden,
				Input:    CreateValidatorInputRequest(validator.Input),
				Expected: CreateValidatorOutputRequest(validator.Output),
			},
		})
	}

	return &result
}

// `FlattenItemIdentifiers` flattens the item identifiers into a pointer of an array of `ItemIdentifier`.
// @returns {*ItemIdentifier[]} The flattened item identifiers.
func (componentData *ComponentData) FlattenItemIdentifiers() *[]string {
	result := make([]string, 0)

	if componentData.Items == nil {
		return nil
	}

	for _, item := range componentData.Items {
		result = append(result, item.ID)
	}

	return &result
}

// `ItemIdentifier` is the identifier of an item.
// @property {string} ID - The unique identifier for the item.
type ItemIdentifier struct {
	ID string
}

// `Validator` is a test to pass for the user in order to validate the content, it is attached to an editor component.
// @property {string} ID - The ID of the validator. This is used to identify the validator.
// @property {bool} IsHidden - If true, the validator will not be displayed in the UI.
// @property {ValidatorInput} Input - The input to the validator.
// @property {ValidatorOutput} Output - The expected output of the validator.
type Validator struct {
	ID       string
	IsHidden bool
	Input    ValidatorInput
	Output   ValidatorOutput
}

// `ValidatorInput` is the input to a validator.
// @property {[]string} Stdin - A list of stdin.
type ValidatorInput struct {
	Stdin []string
}

// `ValidatorOutput` is the expected output of a validator.
// @property {[]string} Stdout - A list of stdout.
type ValidatorOutput struct {
	Stdout []string
}

// `EditorSettings` is the settings for an editor component.
// @property {Language[]} Languages - An array of available language for the editor.
type EditorSettings struct {
	Languages []Language
}

// `Language` is the language settings for an editor component.
// @property {string} DefaultCode - The default language code for the language.
// @property {string} Language - The name language. Can be one of the following: `NODE`, `PYTHON`, `JAVA`, `RUST`.
// @property {string} Version - The version of the language.
type Language struct {
	DefaultCode string
	Language    string
	Version     string
}
