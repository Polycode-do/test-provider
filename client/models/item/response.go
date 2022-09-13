package item

type GetItemResponse struct {
	ID   string              `json:"id"`
	Type string              `json:"type"`
	Data GetItemResponseData `json:"data"`
	Cost int64               `json:"cost"`
}

type GetItemResponseData struct {
	Text string `json:"text"`
}

type CreateItemResponse struct {
	GetItemResponse
}

type UpdateItemResponse struct {
	GetItemResponse
}
