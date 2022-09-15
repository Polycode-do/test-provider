package item

import "polycode-provider/client/shared"

// `Item` is a buyable item in Polycode.
// @property {string} ID - The ID of the item.
// @property {string} Type - The type of item. Only `hint` is available at the moment.
// @property {ItemData} Data - This is the data that is stored in the item.
// @property {int64} Cost - The cost of the item in the store.
type Item struct {
	ID   string
	Type string
	Data ItemData
	Cost int64
}

// `IntoCreateItemRequest` converts an `Item` into a `CreateItemRequest`.
// @returns {CreateItemRequest} The `CreateItemRequest` that was created.
func (i *Item) IntoCreateItemRequest() CreateItemRequest {
	return CreateItemRequest{
		Type: i.Type,
		Data: CreateItemRequestData{
			Text: shared.ConvertNilString(i.Data.Text),
		},
		Cost: i.Cost,
	}
}

// `IntoUpdateItemRequest` converts an `Item` into a `UpdateItemRequest`.
// @returns {UpdateItemRequest} The `UpdateItemRequest` that was created.
func (i *Item) IntoUpdateItemRequest() UpdateItemRequest {
	return UpdateItemRequest{
		CreateItemRequest: CreateItemRequest{
			Type: i.Type,
			Cost: i.Cost,
			Data: CreateItemRequestData{
				Text: shared.ConvertNilString(i.Data.Text),
			},
		},
	}
}

// `ItemData` is the data that is stored in the item.
// @property {string} Text - The text to display in the item.
type ItemData struct {
	Text string
}
