package cmd

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/tsivinsky/fss/session"
)

func DeleteSession() error {
	name := flag.Arg(1)
	if name == "" {
		return fmt.Errorf("No session name provided")
	}

	scriptName := fmt.Sprintf("%s.sh", name)
	p := path.Join(session.SessionsDir, scriptName)

	if _, err := os.Stat(p); os.IsNotExist(err) {
		return fmt.Errorf("Session with name \"%s\" doesn't exist", name)
	}

	err := os.Remove(p)
	if err != nil {
		return err
	}

	return nil
}
