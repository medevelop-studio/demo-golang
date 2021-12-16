package dto

type ChatMessageWTNRequest struct {
	RequestGeneral
	Data *ChatMessageWTNData `json:"data"`
}

type ChatMessageWTNData struct {
	Id       string `json:"id"`
	Content  string `json:"content"`
	UserId   string `json:"userId"`
	Username string `json:"username"`
	Date     int64  `json:"date"`
}

type ChatMessageNTWResponse struct {
	ResponseGeneral
	Data ChatMessageNTWData `json:"data"`
}

type ChatMessageNTWData struct {
	Id       string `json:"id"`
	Content  string `json:"content"`
	UserId   string `json:"userId"`
	Username string `json:"username"`
	Date     int64  `json:"date"`
}
