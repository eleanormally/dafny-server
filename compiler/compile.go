package compiler

import (
	"os"
	"os/exec"
)

func compileAtTmp() (string, error) {
	files, err := os.ReadDir("/tmp/dafny-server")
	if err != nil {
		return "", err
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
	return string(stdout), err
}
