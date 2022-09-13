package item

type CreateItemRequest struct {
	Type string                `json:"type"`
	Data CreateItemRequestData `json:"data"`
	Cost int64                 `json:"cost"`
}

type CreateItemRequestData struct {
	Text *string `json:"text"`
}

type UpdateItemRequest struct {
	CreateItemRequest
	Data UpdateItemRequestData `json:"data"`
}

type UpdateItemRequestData struct {
	CreateItemRequestData
}
