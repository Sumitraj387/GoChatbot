package model

import "time"

type ChatRequest struct {
	SenderId    int    `json:"sender_id"`
	ReceiverId  int    `json:"receiver_id"`
	MessageText string `json:"message_text"`
}

func CreateMessagePayloadFromRequestPayload(chat ChatRequest) ChatMessages {
	return ChatMessages{
		SenderId:    chat.SenderId,
		ReceiverId:  chat.ReceiverId,
		MessageText: chat.MessageText,
		SentAt:      time.Now(),
	}
}
func CreateImagePayloadFromRequestPayload(chat ChatRequest, image string) ChatImages {
	return ChatImages{
		SenderId:   chat.SenderId,
		ReceiverId: chat.ReceiverId,
		ImageData:  image,
		SentAt:     time.Now(),
	}
}
