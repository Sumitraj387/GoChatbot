package core

import (
	"context"

	"GoChatbot/chatbot/v1/model"
	"GoChatbot/chatbot/v1/repository"
	"GoChatbot/globals"

	"github.com/sirupsen/logrus"
)

type ICore interface {
	Chat(ctx context.Context, chat model.ChatRequest) (model.ChatResponse, error)
}
type Core struct {
	Logger *logrus.Entry
	RepoV1 repository.Repository
}

func (c Core) Chat(ctx context.Context, chat model.ChatRequest) (model.ChatResponse, error) {
	c.Logger.Info("core GetChats() entry")
	chatMessage := model.CreateMessagePayloadFromRequestPayload(chat)
	//transaction begin
	c.RepoV1.Db.Begin()
	defer c.RepoV1.Db.Rollback()
	createChatMessageResponse, err := c.RepoV1.SaveMessage(ctx, chatMessage)
	if err != nil {
		return model.ChatResponse{}, err
	}

	res := model.CreateChatResponsePayload(createChatMessageResponse, chat.SenderId, chat.ReceiverId)
	image := globals.IMG
	img := model.CreateImagePayloadFromRequestPayload(chat, image)
	createImageMessageResponse, err := c.RepoV1.SaveImage(ctx, img)
	if err != nil {
		return model.ChatResponse{}, err
	}
	res.ImageData = createImageMessageResponse.ImageData

	c.RepoV1.Db.Commit()
	c.Logger.Info("core GetChats() exit")
	return res, nil
}
