package repository

import (
	modelV1 "GoChatbot/chatbot/v1/model"
	"GoChatbot/globals"
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IRepository interface {
	SaveMessage(ctx context.Context, message modelV1.ChatMessages) (modelV1.ChatMessages, error)
	SaveImage(ctx context.Context, message modelV1.ChatImages) (modelV1.ChatImages, error)
	// RetrieveAllImages(ctx context.Context) ([]modelV1.ChatImages, error)
	// RetrieveAllMessages(ctx context.Context) ([]modelV1.ChatMessages, error)
}
type Repository struct {
	Db     *gorm.DB
	Logger *logrus.Entry
}

func (r Repository) SaveMessage(ctx context.Context, message modelV1.ChatMessages) (modelV1.ChatMessages, error) {
	r.Logger.Info("SaveMessage() enter")
	err := r.Db.Debug().Table(globals.CHAT_MESSAGE_TABLE).Create(&message).Error
	if err != nil {
		return modelV1.ChatMessages{}, err
	}
	r.Logger.Info("SaveMessage() exit")
	return message, err
}
func (r Repository) SaveImage(ctx context.Context, message modelV1.ChatImages) (modelV1.ChatImages, error) {
	r.Logger.Info("SaveImage() enter")
	err := r.Db.Debug().Table(globals.CHAT_IMAGE_TABLE).Create(&message).Error
	if err != nil {
		return modelV1.ChatImages{}, err
	}
	r.Logger.Info("SaveImage() exit")
	return message, err
}
