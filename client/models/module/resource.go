package module

// `Module` is a folder containing content and other modules.
// @property {string} ID - The unique identifier for the module.
// @property {string} Name - The name of the module.
// @property {string} Description - A short description of the module.
// @property {string} Type - The type of module. This can be one of the following:
// `challenge`, `practice`, `certification`, `submodule`
// @property {int64} Reward - The amount of points the user will receive for completing this module.
// @property {[]string} Tags - A list of tags that can be used to search for this module.
// @property {ModuleData} Data - This is the data that the module will use to generate the content.
// @property {[]ModuleIdentifier} Modules - A list of modules that are required to be completed before
// this module can be completed.
// @property {[]ContentIdentifier} Contents - A list of ContentIdentifier objects. These are the
// contents that are required to complete the module.
type Module struct {
	ID          string
	Name        string
	Description string
	Type        string
	Reward      int64
	Tags        []string
	Data        ModuleData
	Modules     []ModuleIdentifier
	Contents    []ContentIdentifier
}

// `IntoCreateModuleRequest` converts a module into a create module request.
// @returns {CreateModuleRequest} The create module request.
func (m *Module) IntoCreateModuleRequest() CreateModuleRequest {
	return CreateModuleRequest{
		Name:        m.Name,
		Description: m.Description,
		Type:        m.Type,
		Reward:      m.Reward,
		Tags:        m.Tags,
		Data:        CreateModuleRequestData{},
		Modules:     *m.FlattenModuleIdentifiers(),
		Contents:    *m.FlattenContentIdentifiers(),
	}
}

// `IntoUpdateModuleRequest` converts a module into a update module request.
// @returns {UpdateModuleRequest} The update module request.
func (m *Module) IntoUpdateModuleRequest() UpdateModuleRequest {
	return UpdateModuleRequest{
		Name:        &m.Name,
		Description: &m.Description,
		Type:        &m.Type,
		Reward:      &m.Reward,
		Tags:        &m.Tags,
		Data:        &UpdateModuleRequestData{},
		Modules:     m.FlattenModuleIdentifiers(),
		Contents:    m.FlattenContentIdentifiers(),
	}
}

// `FlattenModuleIdentifiers` flattens the module identifiers into a list of IDs.
// @returns {[]string} The list of IDs.
func (m *Module) FlattenModuleIdentifiers() *[]string {
	result := make([]string, 0)

	if m.Modules == nil {
		return nil
	}

	for _, module := range m.Modules {
		result = append(result, module.ID)
	}

	return &result
}

// `FlattenContentIdentifiers` flattens the content identifiers into a list of IDs.
// @returns {[]string} The list of IDs.
func (m *Module) FlattenContentIdentifiers() *[]string {
	result := make([]string, 0)

	if m.Contents == nil {
		return nil
	}

	for _, content := range m.Contents {
		result = append(result, content.ID)
	}

	return &result
}

// `ModuleData` is the data that the module will use to generate the content,
// for now it is empty but it will be used in the future
type ModuleData struct{}

// `ModuleIdentifier` is a module identifier.
// @property {string} ID - The unique identifier for the module.
type ModuleIdentifier struct {
	ID string
}

// `ContentIdentifier` is a content identifier.
// @property {string} ID - The unique identifier for the content.
type ContentIdentifier struct {
	ID string
}
