package message_domain

import (
	"time"

	"github.com/google/uuid"
)

type CreateMessageDto struct {
	UserId  string
	Content string
}

func Create(data *CreateMessageDto) (*Message, error) {
	id := uuid.New().String()
	date := time.Now().Unix()

	message := &Message{
		Id:      id,
		UserId:  data.UserId,
		Content: data.Content,
		Date:    date,
	}

	return message, nil
}
