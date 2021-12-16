package message_domain

type Message struct {
	Id      string `json:"id"`
	UserId  string `json:"userId"`
	Content string `json:"content"`
	Date    int64  `json:"date"`
}
