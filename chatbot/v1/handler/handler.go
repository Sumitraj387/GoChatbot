package v1

import (
	httpserver "GoChatbot/chatbot/v1/httpServer"

	coreV1 "GoChatbot/chatbot/v1/core"
	repoV1 "GoChatbot/chatbot/v1/repository"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type HandlerV1 struct {
	Router *mux.Router
	Logger *logrus.Entry
	Db     *gorm.DB
}

func (h HandlerV1) Init() {
	httpServerV1 := httpserver.HttpServer{
		Router: h.Router,
		Logger: h.Logger,
		Core: coreV1.Core{
			Logger: h.Logger,
			RepoV1: repoV1.Repository{
				Db:     h.Db,
				Logger: h.Logger,
			},
		},
	}
	err := httpServerV1.Init()
	if err != nil {
		panic(err)
	}
}
