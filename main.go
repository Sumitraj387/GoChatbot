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
	handlerV1 := handlerV1.HandlerV1{
		Router: apiRouter,
		Logger: logger,
	}
	handlerV1.Init()
	muxHttpHandler := providers.GetMux(config)
	if err := http.ListenAndServe(config.HttpConfig.Address, muxHttpHandler); err != nil {
		panic(err)
	}
	// openAiClient := openAi.NewClient(config.OpenAiConfig.SecretKey)
	// openAiClient.CreateChatCompletion()

}
