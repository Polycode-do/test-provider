package user

type GetUserResponse struct {
	Metadata struct{}            `json:"metadata"`
	Data     GetUserResponseData `json:"data"`
}

type GetUserResponseData struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	Description string `json:"description"`
	Points      int64    `json:"points"`
	Rank        int64    `json:"rank"`
}
