package gogen

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

// GetPackageOrDie returns the name of the current package. Any failure will
// cause the program to exit.
func GetPackageOrDie() string {
	return runGoList("{{.Name}}")
}

// GetImportPathOrDie returns the import path of the current package. Any
// failure will cause the program to exit.
func GetImportPathOrDie() string {
	return runGoList("{{.ImportPath}}")
}

// runGoList runs the go list command, capturing and returning the output. If
// the command fails for any reason, the output is printed and the program
// exits. Any white space at the start or end is removed.
func runGoList(format string) string {
	out, err := exec.Command("go", "list", "-f", format).CombinedOutput()
	if err != nil {
		fmt.Fprintln(os.Stderr, string(out))
		fmt.Fprintln(os.Stderr, "The go list command failed: ", err)
		os.Exit(1)
	}
	out = bytes.TrimSpace(out)
	return string(out)
}
