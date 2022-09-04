package content

type Content struct {
	ID            string
	Name          string
	Description   string
	Type          string
	Reward        int64
	RootComponent Component
	Data          ContentData
}

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
				Markdown:       convertNilString(content.RootComponent.Data.Markdown),
				Items:          content.RootComponent.Data.IntoStringSlice(),
				Validators:     content.RootComponent.Data.IntoCreateValidatorRequest(),
				EditorSettings: content.RootComponent.Data.IntoCreateEditorSettingsRequest(),
				Orientation:    convertNilString(content.RootComponent.Orientation),
			},
		},
		Data: CreateContentRequestData{},
	}
}

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
					Markdown:       convertNilString(content.RootComponent.Data.Markdown),
					Items:          content.RootComponent.Data.IntoStringSlice(),
					EditorSettings: content.RootComponent.Data.IntoCreateEditorSettingsRequest(),
					Orientation:    convertNilString(content.RootComponent.Orientation),
				},
			},
		},
		Data: UpdateContentRequestData{},
	}
}

type ContentData struct{}

type Component struct {
	ID          string
	Type        string
	Orientation string
	Data        ComponentData
}

type ComponentData struct {
	Components     []Component
	Markdown       string
	Items          []Item
	Validators     []Validator
	EditorSettings EditorSettings
}

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
				Markdown:       convertNilString(component.Data.Markdown),
				Items:          component.Data.IntoStringSlice(),
				Validators:     component.Data.IntoCreateValidatorRequest(),
				EditorSettings: component.Data.IntoCreateEditorSettingsRequest(),
				Orientation:    convertNilString(component.Orientation),
			},
		})
	}

	return &result
}

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
					Markdown:       convertNilString(component.Data.Markdown),
					Items:          component.Data.IntoStringSlice(),
					EditorSettings: component.Data.IntoCreateEditorSettingsRequest(),
					Orientation:    convertNilString(component.Orientation),
				},
			},
		})
	}

	return &result
}

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

func (componentData *ComponentData) IntoStringSlice() *[]string {
	result := make([]string, 0)

	if componentData.Items == nil {
		return nil
	}

	for _, item := range componentData.Items {
		result = append(result, item.Data.Text)
	}

	return &result
}

type Item struct {
	ID   string
	Cost int64
	Type string
	Data ItemData
}

type ItemData struct {
	Text string
}

type Validator struct {
	ID       string
	IsHidden bool
	Input    ValidatorInput
	Output   ValidatorOutput
}

type ValidatorInput struct {
	Stdin []string
}

type ValidatorOutput struct {
	Stdout []string
}

type EditorSettings struct {
	Languages []Language
}

type Language struct {
	DefaultCode string
	Language    string
	Version     string
}

func convertNilString(str string) *string {
	if str == "" {
		return nil
	}

	return &str
}
