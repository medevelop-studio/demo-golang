package codes

type MessageTypeCode byte

const (
	MESSAGE_ERROR_TYPE_START = 1

	MESSAGE_INFO_TYPE_START = 30

	MESSAGE_REQ_RES_TYPE_START     = 90
	MESSAGE_INTERNAL_REQ_RES_START = 200
)

const (
	LoginRequest MessageTypeCode = iota + MESSAGE_REQ_RES_TYPE_START
	LoginResponse

	ChatMessageRequest
	ChatMessageResponse
)
