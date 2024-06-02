package endpoints

import (
	"dafny-server/compiler"
	"encoding/json"
	"fmt"
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

		resultChan := make(chan compiler.CompilationResult)
		c.AddCodeInstanceToQueue(compiler.CodeInstance{
			Requester: rb.Requester,
			Files:     rb.Files,
			Result:    resultChan,
		})

		go func() {
			res := <-resultChan
			//TODO: return the values out of this back to the user
			fmt.Printf("compiled and got http status %d with content %s\n", res.Status, res.Content)
		}()

	}
}
