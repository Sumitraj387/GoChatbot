package core

import (
	"context"

	"GoChatbot/chatbot/v1/repository"

	"github.com/sirupsen/logrus"
)

type ICore interface {
	GetChats(ctx context.Context) error
}
type Core struct {
	Logger *logrus.Entry
	RepoV1 repository.IRepository
}

func (c Core) GetChats(ctx context.Context) error {
	c.Logger.Info("core GetChats() entry")
	c.Logger.Info("core GetChats() exit")
	return nil
}
