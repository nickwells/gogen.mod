package gogen

import (
	"fmt"
	"os"
)

// MakeFileOrDie creates the file truncating it if it already exists and
// returning the open file. If an error is detected, it is reported and the
// program exits
func MakeFileOrDie(filename string) *os.File {
	packageName := GetPackageOrDie()
	f, err := os.Create(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating the Go file:", err)
		fmt.Fprintln(os.Stderr, "filename:", filename)
		os.Exit(1)
	}
	fmt.Fprintln(f, "package ", packageName)
	return f
}
