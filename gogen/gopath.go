package gogen

import (
	"os"
	"path/filepath"
)

// GetGopath returns the value of the GOPATH variable (strictly the first
// entry in the list if it is set). If the GOPATH env var is not set then the
// default value is returned: $HOME/go.
func GetGopath() string {
	gopath := os.Getenv("GOPATH")
	parts := filepath.SplitList(gopath)
	if len(parts) > 0 {
		return parts[0]
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(home, "go")
}
