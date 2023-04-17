package cli

import "os"

func GetEditorPath() string {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vim"
	}

	return editor
}
