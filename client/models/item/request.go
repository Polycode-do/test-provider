package item

// `CreateItemRequest` is the request body for the create item endpoint.
// @property {string} Type - The type of item you want to create. Only `hint` is available at the moment.
// @property {CreateItemRequestData} Data - This is the data that will be stored in the item.
// @property {int64} Cost - The cost of the item in the currency specified in the request.
type CreateItemRequest struct {
	Type string                `json:"type"`
	Data CreateItemRequestData `json:"data"`
	Cost int64                 `json:"cost"`
}

// `CreateItemRequestData` is the data that will be stored in the item.
// @property Text - The text of the item.
type CreateItemRequestData struct {
	Text *string `json:"text"`
}

// `UpdateItemRequest` is the request body for the update item endpoint.
// @property {string} Type - The type of item you want to create.
// @property {CreateItemRequestData} Data - This is the data that will be stored in the item.
// @property {int64} Cost - The cost of the item in the currency specified in the request.
type UpdateItemRequest struct {
	CreateItemRequest
}
