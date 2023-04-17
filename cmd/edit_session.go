package cmd

import (
	"flag"
	"fmt"

	"github.com/tsivinsky/fss/session"
)

func EditSession() error {
	name := flag.Arg(1)
	if name == "" {
		return fmt.Errorf("No session name provided")
	}

	s, err := session.FindByName(name)
	if err != nil {
		return err
	}

	err = s.OpenInEditor()
	if err != nil {
		return err
	}

	return nil
}
