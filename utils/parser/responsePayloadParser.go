package parser

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

func ResponsePayloadParser(data []byte, requestLogger *logrus.Entry) []byte {
	interfaceMap := make(map[string]interface{})
	err := json.Unmarshal(data, &interfaceMap)
	if err != nil {
		requestLogger.WithField("err", err).Error("Error while unmarshal data from ResponsePayloadParser")
		panic(err)
	}
	response, err := json.Marshal(interfaceMap)
	if err != nil {
		requestLogger.WithField("err", err).Error("Error while marshal data from ResponsePayloadParser")
		panic(err)
	}
	return response
}
