package endpoints

import (
	"dafny-server/compiler"
	"encoding/json"
	"io"
	"net/http"
)

type RequestBody struct {
	Requester string `json:"requester"`
	Files     []compiler.DafnyFile
}

func HandleCompile(c compiler.CompilerService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rb := RequestBody{}
		data, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = json.Unmarshal(data, &rb)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid JSON"))
		}

		c.AddCodeInstanceToQueue(compiler.CodeInstance{
			Requester: rb.Requester,
			Files:     rb.Files,
		})
	}
}
