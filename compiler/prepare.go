package compiler

import (
	"fmt"
	"io/fs"
	"os"
)

func prepareCompilationEnvironment(inst CodeInstance) error {
	err := os.RemoveAll("/tmp/dafny-server")
	if err != nil {
		return err
	}

	err = os.MkdirAll("/tmp/dafny-server", fs.ModePerm)
	if err != nil {
		return err
	}

	for _, f := range inst.Files {
		fmt.Printf("preparing file %s\n", f.Name)
		err := os.WriteFile("/tmp/dafny-server/"+f.Name, []byte(f.Content), fs.ModePerm)
		if err != nil {
			return fmt.Errorf("Error writing file %s, error %s", f.Name, err.Error())
		}
	}

	return nil
}
