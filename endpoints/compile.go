package endpoints

import (
	"dafny-server/compiler"
	"net/http"
)

func HandleCompile(c compiler.CompilerService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		c.AddCodeInstanceToQueue(compiler.CodeInstance{
			Requester: "web",
			Files: []compiler.DafnyFile{
				{
					Name:    "test.dfy",
					Content: "hi",
				},
			},
		})
	}
}
