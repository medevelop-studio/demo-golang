package dto

type ChatMessageRequest struct {
	RequestGeneral
	Data *ChatMessageData `json:"data"`
}

type ChatMessageData struct {
	UserId  string `json:"userId"`
	Content string `json:"content"`
}

type ChatMessageResponse struct {
	ResponseGeneral
	Data *MessageToChat `json:"data"`
}

type MessageToChat struct {
	Id       string `json:"id"`
	Content  string `json:"content"`
	UserId   string `json:"userId"`
	Username string `json:"username"`
	Date     int64  `json:"date"`
}
