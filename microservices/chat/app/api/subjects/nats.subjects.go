package subjects

const (
	USER_RESPONSE = "userResponse"

	// Subject: WebSocket User -> Chat microservice
	CHAT_REQUEST = "chatRequest"
	// Query: WebSocket User -> Chat microservice
	CHAT_REQUEST_QUERY = "chatRequestQuery"

	MESSAGE_RESPONSE = "messageResponse"
)

func GetUserResponseSub(userId string) string {
	return USER_RESPONSE + "-" + userId
}

func GetChatResponseSub(chatId string) string {
	return MESSAGE_RESPONSE + "-" + chatId
}
