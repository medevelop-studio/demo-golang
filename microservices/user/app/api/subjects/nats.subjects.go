package subjects

const (
	WS_RESPONSE   = "wsResponse"
	USER_RESPONSE = "userResponse"

	CHAT_REQUEST       = "chatRequest"
	CHAT_REQUEST_QUERY = "chatRequestQuery"

	MESSAGE_RESPONSE = "messageResponse"

	LOGIN_REQUEST = "loginRequest"
	LOGIN_QUEUE   = "loginQueue"
)

func GetUserResponseSub(userId string) string {
	return USER_RESPONSE + "-" + userId
}

func GetChatResponseSub(chatId string) string {
	return MESSAGE_RESPONSE + "-" + chatId
}
