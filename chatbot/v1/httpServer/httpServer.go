package httpserver

import (
	httpresponse "GoChatbot/utils/httpResponse"
	"net/http"

	coreV1 "GoChatbot/chatbot/v1/core"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type HttpServer struct {
	Router *mux.Router
	Logger *logrus.Entry
	Core   coreV1.ICore
}

func (h HttpServer) Init() error {
	h.registerRoutes()
	return nil
}
func (h HttpServer) registerRoutes() {
	h.Router.HandleFunc("/ping", h.Ping()).Methods("GET")
}
func (h HttpServer) Ping() func(rw http.ResponseWriter, rq *http.Request) {
	return func(rw http.ResponseWriter, rq *http.Request) {
		httpresponse.HTTPSuccessWith200("/ping", rw, h.Logger)
	}
}
