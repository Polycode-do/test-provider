package content

type CreateContentRequest struct {
	Name          string                   `json:"name"`
	Description   string                   `json:"description"`
	Type          string                   `json:"type"`
	Reward        int64                    `json:"reward"`
	RootComponent CreateComponentRequest   `json:"rootComponent"`
	Data          CreateContentRequestData `json:"data"`
}

type CreateContentRequestData struct{}

type CreateComponentRequest struct {
	Type string                     `json:"type"`
	Data CreateComponentRequestData `json:"data"`
}

type CreateComponentRequestData struct {
	Components     *[]CreateComponentRequest    `json:"components,omitempty"`
	Markdown       *string                      `json:"markdown,omitempty"`
	Items          *[]string                    `json:"items,omitempty"` //An array of UUID pointing to the items
	Validators     *[]CreateValidatorRequest    `json:"validators,omitempty"`
	EditorSettings *CreateEditorSettingsRequest `json:"editorSettings,omitempty"`
	Orientation    *string                      `json:"orientation,omitempty"`
}

type CreateValidatorRequest struct {
	IsHidden bool                         `json:"isHidden"`
	Input    CreateValidatorInputRequest  `json:"input"`
	Expected CreateValidatorOutputRequest `json:"expected"`
}

type CreateValidatorInputRequest struct {
	Stdin []string `json:"stdin"`
}

type CreateValidatorOutputRequest struct {
	Stdout []string `json:"stdout"`
}

type CreateEditorSettingsRequest struct {
	Languages []CreateLanguageRequest `json:"languages"`
}

type CreateLanguageRequest struct {
	DefaultCode string `json:"defaultCode"`
	Language    string `json:"language"`
	Version     string `json:"version"`
}

type UpdateContentRequest struct {
	CreateContentRequest
	RootComponent UpdateComponentRequest   `json:"rootComponent"`
	Data          UpdateContentRequestData `json:"data"`
}

type UpdateContentRequestData struct{}

type UpdateComponentRequest struct {
	ID string `json:"id"`
	CreateComponentRequest
	Data UpdateComponentRequestData `json:"data"`
}

type UpdateComponentRequestData struct {
	CreateComponentRequestData
	Components *[]UpdateComponentRequest `json:"components,omitempty"`
	Validators *[]UpdateValidatorRequest `json:"validators,omitempty"`
}

type UpdateValidatorRequest struct {
	ID string `json:"id"`
	CreateValidatorRequest
}
