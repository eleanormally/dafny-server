package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"dafny-server/compiler"
	"dafny-server/endpoints"

	"github.com/labstack/echo/v4"
)

func main() {
	c, err := compiler.StartCompilerService()
	if err != nil {
		panic(fmt.Sprintf("Error starting compiler service: %s", err.Error()))
	}

	e := echo.New()

	e.GET("/health", endpoints.HandleHealth(c))
	e.POST("/compile", endpoints.HandleCompile(c))

	e.Logger.Fatal(e.Start(":80"))

	sigChan := make(chan os.Signal)
	endChan := make(chan int)
	signal.Notify(sigChan)
	go func() {
		for {
			s := <-sigChan
			if s == syscall.SIGTERM || s == syscall.SIGINT {
				endChan <- 0
			}
		}
	}()
	<-endChan
	close(sigChan)
	close(endChan)
	os.Exit(0)
}
