package gogen

import "os/exec"

var goCmdName = "go"

// SetGoCmdName sets the Go command name to the explicit path. The value
// given must be an executable program in the PATH or the pathname to an
// executable program. An error will be returned and the setting left
// unchanged if no such executable exists.
func SetGoCmdName(newPath string) error {
	_, err := exec.LookPath(newPath)
	if err == nil {
		goCmdName = newPath
	}
	return err
}

// GetGoCmdName returns the current Go command name
func GetGoCmdName() string {
	return goCmdName
}
