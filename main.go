package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"dafny-server/compiler"
	"dafny-server/endpoints"
)

func main() {
	c, err := compiler.StartCompilerService()
	if err != nil {
		panic(fmt.Sprintf("Error starting compiler service: %s", err.Error()))
	}

	http.HandleFunc("/health", endpoints.HandleHealth(c))
	http.HandleFunc("/compile", endpoints.HandleCompile(c))

	http.ListenAndServe(":80", nil)

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
