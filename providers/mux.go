package providers

import (
	"net/http"

	"github.com/justinas/alice"
)

func GetMux(config AppConfig) http.Handler {
	mux := alice.New().Then(nil)
	return mux
}
