package httpresponse

import (
	"encoding/json"
	"net/http"

	"GoChatbot/utils/parser"

	"github.com/sirupsen/logrus"
)

func HTTPFailWith4xx(errorMessage string, httpStatusCode int, rw http.ResponseWriter, requestLogger *logrus.Entry) {
	responseJson, err := json.Marshal(struct {
		Success      bool   `json:"success"`
		ErrorMessage string `json:"error_message"`
	}{
		false,
		errorMessage,
	})
	if err != nil {
		requestLogger.WithField("err", err).Error("Error while marshal data from HTTPFailWith4xx")

		panic(err)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(httpStatusCode)
	rw.Write(responseJson)
}

func HTTPSuccessWith200(data interface{}, rw http.ResponseWriter, requestLogger *logrus.Entry) {
	responseJson, err := json.Marshal(struct {
		Success bool        `json:"success"`
		Data    interface{} `json:"data"`
	}{
		true,
		data,
	})
	if err != nil {
		requestLogger.WithField("err", err).Error("Error while marshal data from HTTPSuccessWith200")
		panic(err)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(parser.ResponsePayloadParser(responseJson, requestLogger))
}

func HTTPSuccessWithPaginated200(data interface{}, pageInfo interface{}, rw http.ResponseWriter, requestLogger *logrus.Entry) {
	responseJson, err := json.Marshal(struct {
		Success  bool        `json:"success"`
		Data     interface{} `json:"data"`
		PageInfo interface{} `json:"page_info"`
	}{
		true,
		data,
		pageInfo,
	})
	if err != nil {
		requestLogger.WithField("err", err).Error("Error while marshal data from HTTPSuccessWithPaginated200")
		panic(err)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(parser.ResponsePayloadParser(responseJson, requestLogger))
}

func HTTPFailWith5xx(errorMessage string, httpStatusCode int, rw http.ResponseWriter, requestLogger *logrus.Entry) {
	responseJson, err := json.Marshal(struct {
		Success      bool   `json:"success"`
		ErrorMessage string `json:"error_message"`
	}{
		false,
		errorMessage,
	})
	if err != nil {
		requestLogger.WithField("err", err).Error("Error while marshal data from HTTPFailWith5xx")
		panic(err)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(httpStatusCode)
	rw.Write(responseJson)
}

func HTTPFailWithCustomErrorCode4xx(errorMessage string, customErrorCode string, httpStatusCode int, rw http.ResponseWriter, requestLogger *logrus.Entry) {
	responseJson, err := json.Marshal(struct {
		Success         bool   `json:"success"`
		ErrorMessage    string `json:"error_message"`
		CustomErrorCOde string `json:"custom_error_code"`
	}{
		false,
		errorMessage,
		customErrorCode,
	})
	if err != nil {
		requestLogger.WithField("err", err).Error("Error while marshal data from HTTPFailWithCustomErrorCode4xx")

		panic(err)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(httpStatusCode)
	rw.Write(responseJson)
}
