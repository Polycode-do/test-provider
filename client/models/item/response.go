package item

import "polycode-provider/client/shared"

// `GetItemResponse` is the response body for the get item endpoint.
// @property {string} ID - The ID of the item.
// @property {string} Type - The type of the item. Only `hint` is available at the moment.
// @property {GetItemResponseData} Data - The data that is stored in the item.
// @property {int64} Cost - The cost of the request.
type GetItemResponse struct {
	ID   string              `json:"id"`
	Type string              `json:"type"`
	Data GetItemResponseData `json:"data"`
	Cost int64               `json:"cost"`
}

// `IntoItem` converts a `GetItemResponse` into a pointer of an `Item`.
// @returns {Item} The `Item` that was created.
func (i *GetItemResponse) IntoItem() *Item {
	return &Item{
		ID:   i.ID,
		Type: i.Type,
		Data: ItemData{
			Text: shared.ConvertNilStringPointer(i.Data.Text),
		},
		Cost: i.Cost,
	}
}

// `GetItemResponseData` is the data that is stored in the item.
// @property Text - The text of the item.
type GetItemResponseData struct {
	Text *string `json:"text"`
}

// `CreateItemResponse` is the response body for the create item endpoint.
// @property {string} ID - The ID of the item.
// @property {string} Type - The type of the item.
// @property {GetItemResponseData} Data - The data that is stored in the item.
// @property {int64} Cost - The cost of the request.
type CreateItemResponse struct {
	GetItemResponse
}

// `UpdateItemResponse` is the response body for the update item endpoint.
// @property {string} ID - The ID of the item.
// @property {string} Type - The type of the item.
// @property {GetItemResponseData} Data - The data that is stored in the item.
// @property {int64} Cost - The cost of the request.
type UpdateItemResponse struct {
	GetItemResponse
}
