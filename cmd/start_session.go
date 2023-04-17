package cmd

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	"github.com/tsivinsky/fss/session"
)

func StartSession() error {
	name := flag.Arg(0)
	if name == "" {
		return fmt.Errorf("No cmd provided")
	}

	scriptPath, err := session.GetScriptByName(name)
	if err != nil {
		return err
	}

	cmd := exec.Command(scriptPath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
