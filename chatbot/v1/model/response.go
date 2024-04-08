package model

type ChatResponse struct {
	SenderId    int    `json:"sender_id"`
	ReceiverId  int    `json:"receiver_id"`
	MessageText string `json:"message_text"`
	ImageData   string `json:"image_data"`
}

func CreateChatResponsePayload(chatMessages ChatMessages, senderId int, receiverId int) ChatResponse {
	return ChatResponse{
		SenderId:    senderId,
		ReceiverId:  receiverId,
		MessageText: "I am fine",
	}
}
