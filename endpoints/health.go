package endpoints

import (
	"fmt"
	"net/http"

	"dafny-server/compiler"
)

func HandleHealth(c compiler.CompilerService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		response := fmt.Sprintf("%d", c.GetQueueSize())
		w.Write([]byte(response))
	}
}
