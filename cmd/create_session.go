package cmd

import (
	"github.com/tsivinsky/fss/cli"
	"github.com/tsivinsky/fss/session"
)

func CreateSession() error {
	name, err := cli.Prompt("Session name: ")
	if err != nil {
		return err
	}

	s, err := session.New(name)
	if err != nil {
		return err
	}

	err = s.OpenInEditor()
	if err != nil {
		return err
	}

	return nil
}
