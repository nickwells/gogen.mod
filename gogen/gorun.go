package gogen

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

// exitOnErrorType controls whether or not to exit on program failure
type exitOnErrorType int

const (
	exitOnErr exitOnErrorType = iota
	dontExitOnErr
)

// runGo runs the command and reports any errors. If ctrl is exitOnErr it
// exits if the program fails. It returns true if it succeeds, false
// otherwise.
func runGo(cmd *exec.Cmd, ctrl exitOnErrorType) bool {
	var b bytes.Buffer
	if cmd.Stdout == nil {
		cmd.Stdout = &b
	}
	if cmd.Stderr == nil {
		cmd.Stderr = &b
	}

	err := cmd.Run()
	if err != nil {
		command := cmd.Path + " " + strings.Join(cmd.Args[1:], " ")
		fmt.Fprintln(os.Stderr, "Command failed:", command)
		fmt.Fprintln(os.Stderr, "         Error:", err)
		fmt.Fprintln(os.Stderr, b.String())
		if ctrl == exitOnErr {
			os.Exit(1)
		}
		return false
	}
	return true
}

// ExecGoCmd will exec the go program with the supplied arguments. If it
// detects an error it will report it and exit.
func ExecGoCmd(ioMode CmdIOType, args ...string) {
	cmd := exec.Command(goCmdName, args...)
	if ioMode == ShowCmdIO {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	runGo(cmd, exitOnErr)
}

// ExecGoCmdCaptureOutput will exec the go program with the supplied
// arguments. The command's Stdout is connected to the supplied writer if it
// is not nil. If it detects an error it will report it and exit.
func ExecGoCmdCaptureOutput(w io.Writer, args ...string) {
	cmd := exec.Command(goCmdName, args...)
	if w != nil {
		cmd.Stdout = w
	}
	runGo(cmd, exitOnErr)
}

// ExecGoCmdNoExit will exec the go program with the supplied arguments. If
// it detects an error it will report it and return false. Otherwise it
// returns true
func ExecGoCmdNoExit(ioMode CmdIOType, args ...string) bool {
	cmd := exec.Command(goCmdName, args...)
	if ioMode == ShowCmdIO {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	return runGo(cmd, dontExitOnErr)
}

// ExecGoCmdCaptureOutputNoExit will exec the go program with the supplied
// arguments. The command's Stdout is connected to the supplied
// writer if it is not nil. If
// it detects an error it will report it and return false. Otherwise it
// returns true
func ExecGoCmdCaptureOutputNoExit(w io.Writer, args ...string) bool {
	cmd := exec.Command(goCmdName, args...)
	if w != nil {
		cmd.Stdout = w
	}
	return runGo(cmd, dontExitOnErr)
}
