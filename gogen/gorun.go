package gogen

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

// runGoOrDie runs the command and exits if it fails
func runGoOrDie(cmd *exec.Cmd) {
	err := cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't exec the command")
		fmt.Fprintf(os.Stderr, "\t%s %s\n",
			cmd.Path, strings.Join(cmd.Args, " "))
		fmt.Fprintln(os.Stderr, "\tError:", err)
		os.Exit(1)
	}
}

// ExecGoCmd will exec the go program with the supplied arguments. If it
// detects an error it will report it and exit
func ExecGoCmd(ioMode CmdIOType, args ...string) {
	cmd := exec.Command("go", args...)
	if ioMode == ShowCmdIO {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	runGoOrDie(cmd)
}

// ExecGoCmdCaptureOutput will exec the go program with the supplied
// arguments. The command's Stdout is connected to the supplied
// writer if it is not nil.
func ExecGoCmdCaptureOutput(w io.Writer, args ...string) {
	cmd := exec.Command("go", args...)
	if w != nil {
		cmd.Stdout = w
	}
	runGoOrDie(cmd)
}
