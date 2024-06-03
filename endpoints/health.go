package endpoints

import (
	"fmt"
	"net/http"

	"dafny-server/compiler"

	"github.com/labstack/echo/v4"
)

func HandleHealth(c compiler.CompilerService) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		response := fmt.Sprintf("%d", c.GetQueueSize())
		return ctx.String(http.StatusOK, response)
	}
}
