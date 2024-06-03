package endpoints

import (
	"dafny-server/compiler"
	"encoding/json"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RequestBody struct {
	Requester string `json:"requester"`
	Files     []compiler.DafnyFile
}

func HandleCompile(c compiler.CompilerService) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		rb := RequestBody{}
		data, err := io.ReadAll(ctx.Request().Body)
		if err != nil {
			return ctx.String(http.StatusBadRequest, "")
		}

		err = json.Unmarshal(data, &rb)
		if err != nil {
			return ctx.String(http.StatusBadRequest, "Invalid JSON")
		}

		resultChan := make(chan compiler.CompilationResult)
		c.AddCodeInstanceToQueue(compiler.CodeInstance{
			Requester: rb.Requester,
			Files:     rb.Files,
			Result:    resultChan,
		})

		res := <-resultChan
		return ctx.String(res.Status, res.Content)

	}
}
