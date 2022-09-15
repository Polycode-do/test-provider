package content

// `CreateContentRequest` is the request body for the create content endpoint.
// @property {string} Name - The name of the content.
// @property {string} Description - The description of the content.
// @property {string} Type - The type of content you want to create. Only `exercise` is available at the moment.
// @property {int64} Reward - The amount of points the user will receive for completing the content.
// @property {CreateComponentRequest} RootComponent - The root component of the content.
// @property {CreateContentRequestData} Data - This is the data that will be used to create the
// content.
type CreateContentRequest struct {
	Name          string                   `json:"name"`
	Description   string                   `json:"description"`
	Type          string                   `json:"type"`
	Reward        int64                    `json:"reward"`
	RootComponent CreateComponentRequest   `json:"rootComponent"`
	Data          CreateContentRequestData `json:"data"`
}

// `CreateContentRequestData` is the data that will be used to create the content,
// for now it is empty but it will be used in the future
type CreateContentRequestData struct{}

// `CreateComponentRequest` is the request body holding information about a component.
// @property {string} Type - The type of component you want to create. Can be one of the following:
// `markdown`, `editor`, `container`.
// @property {CreateComponentRequestData} Data - This is the data that will be used to create the
// component.
type CreateComponentRequest struct {
	Type string                     `json:"type"`
	Data CreateComponentRequestData `json:"data"`
}

// `CreateComponentRequestData` is the data that will be used to create the component,
// this structure holds pointers because all the properties are optional.
// @property {*[]CreateComponentRequest} Components - An array of nested components (only if type = container).
// @property {*string} Markdown - The markdown content of the component (only if type = markdown).
// @property {*[]string} Items - An array of UUID pointing to the items related to the component (only if type = editor).
// @property {*[]CreateValidatorRequest} Validators - An array of validators to be applied to the component (only if type = editor).
// @property {*CreateEditorSettingsRequest} EditorSettings - This is the settings for the editor (only if type = editor).
// @property {*string} Orientation - The orientation of the component. Can be either "horizontal" or "vertical" (only if type = container).
type CreateComponentRequestData struct {
	Components     *[]CreateComponentRequest    `json:"components,omitempty"`
	Markdown       *string                      `json:"markdown,omitempty"`
	Items          *[]string                    `json:"items,omitempty"` //An array of UUID pointing to the items
	Validators     *[]CreateValidatorRequest    `json:"validators,omitempty"`
	EditorSettings *CreateEditorSettingsRequest `json:"editorSettings,omitempty"`
	Orientation    *string                      `json:"orientation,omitempty"`
}

// `CreateValidatorRequest` is the request body holding information about a validator.
// @property {bool} IsHidden - This is a boolean value that determines whether the validator is hidden
// or not.
// @property {CreateValidatorInputRequest} Input - The input to the validator.
// @property {CreateValidatorOutputRequest} Expected - The expected output of the validator.
type CreateValidatorRequest struct {
	IsHidden bool                         `json:"isHidden"`
	Input    CreateValidatorInputRequest  `json:"input"`
	Expected CreateValidatorOutputRequest `json:"expected"`
}

// `CreateValidatorInputRequest` is the request body holding information about the input of a validator.
// @property {[]string} Stdin - A list of stdin.
type CreateValidatorInputRequest struct {
	Stdin []string `json:"stdin"`
}

// `CreateValidatorOutputRequest` is the request body holding information about the expected output of a validator.
// @property {[]string} Stdout - A list of stdout.
type CreateValidatorOutputRequest struct {
	Stdout []string `json:"stdout"`
}

// `CreateEditorSettingsRequest` is the request body holding information about the editor settings.
// @property {[]CreateLanguageRequest} Languages - An array of available language for the editor.
type CreateEditorSettingsRequest struct {
	Languages []CreateLanguageRequest `json:"languages"`
}

// `CreateLanguageRequest` is the request body holding information about a language.
// @property {string} DefaultCode - The default code for the language.
// @property {string} Language - The language name. Can be one of the following: `NODE`, `PYTHON`, `JAVA`, `RUST`.
// @property {string} Version - The version of the language.
type CreateLanguageRequest struct {
	DefaultCode string `json:"defaultCode"`
	Language    string `json:"language"`
	Version     string `json:"version"`
}

// `UpdateContentRequest` is the request body for the update content endpoint.
// @property {string} Name - The name of the content.
// @property {string} Description - The description of the content.
// @property {string} Type - The type of content you want to create. Only `exercise` is available at the moment.
// @property {int64} Reward - The amount of points the user will receive for completing the content.
// @property {UpdateComponentRequest} RootComponent - The root component of the content.
// @property {UpdateContentRequestData} Data - This is the data that will be used to update the
// content.
type UpdateContentRequest struct {
	CreateContentRequest
	RootComponent UpdateComponentRequest   `json:"rootComponent"`
	Data          UpdateContentRequestData `json:"data"`
}

// `UpdateContentRequestData` is the data that will be used to update the content,
// for now it is empty but it will be used in the future
type UpdateContentRequestData struct{}

// `UpdateComponentRequest` is the request body holding information about a component.
// @property {string} ID - The ID of the component.
// @property {string} Type - The type of component you want to create. Can be one of the following:
// `markdown`, `editor`, `container`.
// @property {UpdateComponentRequestData} Data - This is the data that will be used to update the
// component.
type UpdateComponentRequest struct {
	ID string `json:"id"`
	CreateComponentRequest
	Data UpdateComponentRequestData `json:"data"`
}

// `UpdateComponentRequestData` is the data that will be used to update the component,
// this structure holds pointers because all the properties are optional.
// @property {*string} Markdown - The markdown content of the component (only if type = markdown).
// @property {*[]string} Items - An array of UUID pointing to the items related to the component (only if type = editor).
// @property {*CreateEditorSettingsRequest} EditorSettings - This is the settings for the editor (only if type = editor).
// @property {*string} Orientation - The orientation of the component. Can be either "horizontal" or "vertical" (only if type = container).
// @property {*[]UpdateComponentRequest} Components - An array of nested components (only if type = container).
// @property {*[]UpdateValidatorRequest} Validators - An array of validators to be applied to the component (only if type = editor).
type UpdateComponentRequestData struct {
	CreateComponentRequestData
	Components *[]UpdateComponentRequest `json:"components,omitempty"`
	Validators *[]UpdateValidatorRequest `json:"validators,omitempty"`
}

// `UpdateValidatorRequest` is the request body holding information about a validator.
// @property {string} ID - The ID of the validator.
// @property {bool} IsHidden - This is a boolean value that determines whether the validator is hidden
// or not.
// @property {CreateValidatorInputRequest} Input - The input to the validator.
// @property {CreateValidatorOutputRequest} Expected - The expected output of the validator.
type UpdateValidatorRequest struct {
	ID string `json:"id"`
	CreateValidatorRequest
}
