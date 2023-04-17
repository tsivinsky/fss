package session

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path"

	"github.com/tsivinsky/fss/cli"
	"github.com/tsivinsky/fss/config"
)

var (
	SessionsDir = ""
)

func init() {
	SessionsDir = path.Join(config.FssConfigDir, "sessions")

	if _, err := os.Stat(SessionsDir); os.IsNotExist(err) {
		os.MkdirAll(SessionsDir, 0777)
	}
}

type Session struct {
	Name string `json:"name"`
}

func (s *Session) OpenInEditor() error {
	scriptName := fmt.Sprintf("%s.sh", s.Name)
	p := path.Join(SessionsDir, scriptName)
	editor := cli.GetEditorPath()

	cmd := exec.Command(editor, p)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func New(name string) (*Session, error) {
	scriptName := fmt.Sprintf("%s.sh", name)
	p := path.Join(SessionsDir, scriptName)

	if _, err := os.Stat(p); os.IsExist(err) {
		return nil, fmt.Errorf("Session with name %s already exists", name)
	}

	f, err := os.OpenFile(p, 0755, fs.FileMode(os.O_CREATE)|fs.FileMode(os.O_RDWR))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	err = f.Chmod(0755) // i didn't get why putting 0755 in `os.OpenFile` doesn't work so chmod file here too
	if err != nil {
		return nil, err
	}

	f.WriteString("#! /bin/bash\n\n") // add shebang in new script

	return &Session{
		Name: name,
	}, nil
}
