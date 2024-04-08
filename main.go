package main

import (
	"GoChatbot/providers"
	"net/http"

	mux "github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	//openAi "github.com/sashabaranov/go-openai"
	handlerV1 "GoChatbot/chatbot/v1/handler"
)

func main() {
	apiRouter := mux.NewRouter()
	http.Handle("/", apiRouter)
	config, err := providers.GetConfig("GoChatbot-Configuration.yml")
	if err != nil {
		panic(err)
	}
	logger := logrus.New().WithFields(logrus.Fields{
		"hostname": "chatbot",
	})
	gormDb, err := providers.GetGormDbClient(config, logger)
	if err != nil {
		logger.Info("gorm client error = ", err)
		panic(err)
	}
	handlerV1 := handlerV1.HandlerV1{
		Router: apiRouter,
		Logger: logger,
		Db:     gormDb,
	}
	handlerV1.Init()
	muxHttpHandler := providers.GetMux(config)
	if err := http.ListenAndServe("127.0.0.1:8080", muxHttpHandler); err != nil {
		panic(err)
	}
	// openAiClient := openAi.NewClient(config.OpenAiConfig.SecretKey)
	// openAiClient.CreateChatCompletion()

}
