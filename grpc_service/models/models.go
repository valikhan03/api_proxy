package models

type NewAuctionRequestModel struct {
	UserId string `json:"user_id"`
	Title  string `json:"title"`
	Type   string `json:"type"`
	Status string `json:"status"`
	Date   string `json:"date"`
}

type NewAuctionResponseModel struct {
	AuctionId string `json:"auction_id"`
}
