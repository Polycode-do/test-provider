package content

type GetContentResponse struct {
	ID            string               `json:"id"`
	Name          string               `json:"name"`
	Description   string               `json:"description"`
	Type          string               `json:"type"`
	Reward        int64                `json:"reward"`
	RootComponent GetComponentResponse `json:"rootComponent"`
	Data          struct{}             `json:"data"`
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
