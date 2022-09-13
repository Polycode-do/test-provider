package item

type Item struct {
	ID   string
	Type string
	Data ItemData
	Cost int64
}

func (i *Item) IntoCreateItemRequest() CreateItemRequest {
	return CreateItemRequest{
		Type: i.Type,
		Data: CreateItemRequestData{
			Text: convertNilString(i.Data.Text),
		},
		Cost: i.Cost,
	}
}

func (i *Item) IntoUpdateItemRequest() UpdateItemRequest {
	return UpdateItemRequest{
		CreateItemRequest: CreateItemRequest{
			Type: i.Type,
			Cost: i.Cost,
		},
		Data: UpdateItemRequestData{
			CreateItemRequestData{
				Text: convertNilString(i.Data.Text),
			},
		},
	}
}

type ItemData struct {
	Text string
}

func convertNilString(str string) *string {
	if str == "" {
		return nil
	}

	return &str
}
