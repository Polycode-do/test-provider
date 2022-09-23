package module

// `GetModuleResponse` is the response body for getting a module.
// @property {string} ID - The ID of the module.
// @property {string} Name - The name of the module
// @property {string} Description - A description of the module.
// @property {string} Type - The type of module. This can be one of the following:
// `challenge`, `practice`, `certification`, `submodule`
// @property {int64} Reward - The amount of points the user will receive for completing the module.
// @property {[]string} Tags - A list of tags that describe the module.
// @property {GetModuleResponseData} Data - This is the data that the module will use to generate the
// content.
// @property {[]string} Modules - A list of module IDs that are required to complete this module.
// @property {[]string} Contents - A list of content IDs that are part of this module.
type GetModuleResponse struct {
	ID          string                               `json:"id"`
	Name        string                               `json:"name"`
	Description string                               `json:"description"`
	Type        string                               `json:"type"`
	Reward      int64                                `json:"reward"`
	Tags        []string                             `json:"tags"`
	Data        GetModuleResponseData                `json:"data"`
	Modules     []GetModuleResponseModuleIdentifier  `json:"modules"`
	Contents    []GetModuleResponseContentIdentifier `json:"contents"`
}

func (mr *GetModuleResponse) IntoModule() *Module {
	return &Module{
		ID:          mr.ID,
		Name:        mr.Name,
		Description: mr.Description,
		Type:        mr.Type,
		Reward:      mr.Reward,
		Tags:        mr.Tags,
		Data:        ModuleData(mr.Data),
		Modules:     mr.IntoModuleIdentifier(),
		Contents:    mr.IntoContentIdentifier(),
	}
}

// `IntoModuleIdentifier` converts the module IDs into a list of `ModuleIdentifier`
// @returns {[]ModuleIdentifier} The list of `ModuleIdentifier`
func (mr *GetModuleResponse) IntoModuleIdentifier() []ModuleIdentifier {
	modules := make([]ModuleIdentifier, 0)

	if mr.Modules != nil {
		for _, module := range mr.Modules {
			modules = append(modules, ModuleIdentifier(module))
		}
	}

	return modules
}

// `IntoContentIdentifier` converts the content IDs into a list of `ContentIdentifier`
// @returns {[]ContentIdentifier} The list of `ContentIdentifier`
func (mr *GetModuleResponse) IntoContentIdentifier() []ContentIdentifier {
	contents := make([]ContentIdentifier, 0)

	if mr.Contents != nil {
		for _, content := range mr.Contents {
			contents = append(contents, ContentIdentifier(content))
		}
	}

	return contents
}

// `GetModuleResponseData` is the data that the module will use to generate the content,
// for now it is empty but it will be used in the future.
type GetModuleResponseData struct{}

// `GetModuleResponseModuleIdentifier` is the module identifier that is used in the response body
// for getting a module.
type GetModuleResponseModuleIdentifier struct {
	ID string `json:"id"`
}

// `GetModuleResponseContentIdentifier` is the content identifier that is used in the response body
// for getting a module.
type GetModuleResponseContentIdentifier struct {
	ID string `json:"id"`
}

// `CreateModuleResponse` is the response body for creating a module.
type CreateModuleResponse struct {
	GetModuleResponse
	Modules  []string `json:"modules"`
	Contents []string `json:"contents"`
}

// `IntoModule` converts the `CreateModuleResponse` into a `Module`
// @returns {Module} The `Module`
func (mr *CreateModuleResponse) IntoModule() *Module {
	return &Module{
		ID:          mr.ID,
		Name:        mr.Name,
		Description: mr.Description,
		Type:        mr.Type,
		Reward:      mr.Reward,
		Tags:        mr.Tags,
		Data:        ModuleData(mr.Data),
		Modules:     mr.IntoModuleIdentifier(),
		Contents:    mr.IntoContentIdentifier(),
	}
}

// `IntoContentIdentifier` converts the content IDs into a list of `ContentIdentifier`
// @returns {[]ContentIdentifier} The list of `ContentIdentifier`
func (mr *CreateModuleResponse) IntoContentIdentifier() []ContentIdentifier {
	contents := make([]ContentIdentifier, 0)

	if mr.Contents != nil {
		for _, content := range mr.Contents {
			contents = append(contents, ContentIdentifier{ID: content})
		}
	}

	return contents
}

// `IntoModuleIdentifier` converts the module IDs into a list of `ModuleIdentifier`
// @returns {[]ModuleIdentifier} The list of `ModuleIdentifier`
func (mr *CreateModuleResponse) IntoModuleIdentifier() []ModuleIdentifier {
	modules := make([]ModuleIdentifier, 0)

	if mr.Modules != nil {
		for _, module := range mr.Modules {
			modules = append(modules, ModuleIdentifier{ID: module})
		}
	}

	return modules
}
