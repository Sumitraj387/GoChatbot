package v1

import (
	httpserver "GoChatbot/chatbot/v1/httpServer"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type HandlerV1 struct {
	Router *mux.Router
	Logger *logrus.Entry
}

func (h HandlerV1) Init() {
	httpServerV1 := httpserver.HttpServer{
		Router: h.Router,
		Logger: h.Logger,
	}
	err := httpServerV1.Init()
	if err != nil {
		panic(err)
	}
}
