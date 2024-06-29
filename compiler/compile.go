package compiler

import (
	"net/http"
	"os"
	"os/exec"
)

func compileAtTmp() (string, error, int) {
	files, err := os.ReadDir("/tmp/dafny-server")
	if err != nil {
		return "", err, http.StatusInternalServerError
	}
	filePaths := []string{"verify"}
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		filePaths = append(filePaths, "/tmp/dafny-server/"+f.Name())
	}
	cmd := exec.Command("dafny", filePaths...)
	stdout, err := cmd.Output()
	return string(stdout), err, http.StatusOK //output ok status if dafny fails because it is supposed to fail on bad code
}
