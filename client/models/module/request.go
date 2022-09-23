package module

// `CreateModuleRequest` is the request body for creating a module.
// @property {string} Name - The name of the module.
// @property {string} Description - A description of the module
// @property {string} Type - The type of module. This can be one of the following:
// `challenge`, `practice`, `certification`, `submodule`
// @property {int64} Reward - The amount of points the user will receive for completing the module.
// @property {[]string} Tags - A list of tags that describe the module.
// @property {CreateModuleRequestData} Data - This is the data that will be used to create the module.
// @property {[]string} Modules - A list of submodule IDs that this module holds.
// @property {[]string} Contents - A list of content IDs that this module holds.
type CreateModuleRequest struct {
	Name        string                  `json:"name"`
	Description string                  `json:"description"`
	Type        string                  `json:"type"`
	Reward      int64                   `json:"reward"`
	Tags        []string                `json:"tags"`
	Data        CreateModuleRequestData `json:"data"`
	Modules     []string                `json:"modules"`
	Contents    []string                `json:"contents"`
}

// `CreateModuleRequestData` is the data that will be used to create the module,
// for now it is empty but it will be used in the future
type CreateModuleRequestData struct{}

// `UpdateModuleRequest` is the request body for creating a module.
// All properties are optional.
// @property {string} Name - The name of the module.
// @property {string} Description - A description of the module
// @property {string} Type - The type of module. This can be one of the following:
// `challenge`, `practice`, `certification`, `submodule`
// @property {int64} Reward - The amount of points the user will receive for completing the module.
// @property {[]string} Tags - A list of tags that describe the module.
// @property {UpdateModuleRequestData} Data - This is the data that will be used to update the module.
// @property {[]string} Modules - A list of submodule IDs that this module holds.
// @property {[]string} Contents - A list of content IDs that this module holds.
type UpdateModuleRequest struct {
	Name        *string                  `json:"name,omitempty"`
	Description *string                  `json:"description,omitempty"`
	Type        *string                  `json:"type,omitempty"`
	Reward      *int64                   `json:"reward,omitempty"`
	Tags        *[]string                `json:"tags,omitempty"`
	Data        *UpdateModuleRequestData `json:"data,omitempty"`
	Modules     *[]string                `json:"modules,omitempty"`
	Contents    *[]string                `json:"contents,omitempty"`
}

// `UpdateModuleRequestData` is the data that will be used to update the module,
// for now it is empty but it will be used in the future
type UpdateModuleRequestData struct{}
