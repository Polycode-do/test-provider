package item

import "polycode-provider/client/shared"

type GetItemResponse struct {
	ID   string              `json:"id"`
	Type string              `json:"type"`
	Data GetItemResponseData `json:"data"`
	Cost int64               `json:"cost"`
}

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

type GetItemResponseData struct {
	Text *string `json:"text"`
}

type CreateItemResponse struct {
	GetItemResponse
}

type UpdateItemResponse struct {
	GetItemResponse
}
