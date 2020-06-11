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
	return runGoListOrDie("{{.Name}}")
}

// GetImportPathOrDie returns the import path of the current package. Any
// failure will cause the program to exit.
func GetImportPathOrDie() string {
	return runGoListOrDie("{{.ImportPath}}")
}

// runGoListOrDie runs the go list command, capturing and returning the
// output. If the command fails for any reason, the output is printed and the
// program exits. Any white space at the start or end is removed.
func runGoListOrDie(format string) string {
	out, err := runGoList(format)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		fmt.Fprintln(os.Stderr, out)
		fmt.Fprintln(os.Stderr, "The 'go list' command failed: ", err)
		os.Exit(1)
	}
	return out
}

// GetPackage returns the name of the current package.
func GetPackage() (string, error) {
	return runGoList("{{.Name}}")
}

// GetImportPath returns the import path of the current package.
func GetImportPath() (string, error) {
	return runGoList("{{.ImportPath}}")
}

// runGoList runs the go list command, capturing and returning the output. If
// the command fails for any reason, the error is returned. Any white space
// at the start or end is removed.
func runGoList(format string) (string, error) {
	out, err := exec.Command("go", "list", "-f", format).Output()
	if err != nil {
		return string(out), err
	}
	out = bytes.TrimSpace(out)
	return string(out), nil
}
