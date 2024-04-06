package httpserver

import (
	httpresponse "GoChatbot/utils/httpResponse"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type HttpServer struct {
	Router *mux.Router
	Logger *logrus.Entry
}

func (h HttpServer) Init() error {
	h.registerRoutes()
	return nil
}
func (h HttpServer) registerRoutes() {
	h.Router.HandleFunc("/", h.Ping()).Methods("GET")
}
func (h HttpServer) Ping() func(rw http.ResponseWriter, rq *http.Request) {
	return func(rw http.ResponseWriter, rq *http.Request) {
		httpresponse.HTTPSuccessWith200("/ping", rw, h.Logger)
	}
}
