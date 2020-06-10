package gogen

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// ExecGoCmd will exec the go program with the supplied arguments. If it
// detects an error it will report it and exit
func ExecGoCmd(ioMode CmdIOType, args ...string) {
	cmd := exec.Command("go", args...)
	if ioMode == ShowCmdIO {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	err := cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't exec the go command")
		fmt.Fprintln(os.Stderr, "\tgo "+strings.Join(args, " "))
		fmt.Fprintln(os.Stderr, "\tError:", err)
		os.Exit(1)
	}
}
