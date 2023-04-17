package session

import (
	"fmt"
	"os"
	"path"
)

func GetScriptByName(name string) (string, error) {
	scriptName := fmt.Sprintf("%s.sh", name)
	p := path.Join(SessionsDir, scriptName)

	if _, err := os.Stat(p); os.IsNotExist(err) {
		return "", fmt.Errorf("Session with name \"%s\" doesn't exist", name)
	}

	return p, nil
}
