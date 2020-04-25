package gogen

import (
	"fmt"
	"os"
	"strings"
)

// MakeFileOrDie creates the file truncating it if it already exists and
// returning the open file. If an error is detected, it is reported and the
// program exits. It also checks that the file is a go file and does not have
// the _test.go suffix.
func MakeFileOrDie(filename string) *os.File {
	mustHaveSuffix(filename, ".go")
	mustNotHaveSuffix(filename, "_test.go")

	packageName := GetPackageOrDie()
	f, err := os.Create(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating the Go file:", err)
		fmt.Fprintln(os.Stderr, "      filename:", filename)
		os.Exit(1)
	}
	fmt.Fprintln(f, "package "+packageName)
	return f
}

// MakeTestFileOrDie creates the file truncating it if it already exists and
// returning the open file. If an error is detected, it is reported and the
// program exits. It also checks that the file has the _test.go suffix.
func MakeTestFileOrDie(filename string) *os.File {
	mustHaveSuffix(filename, "_test.go")

	packageName := GetPackageOrDie()
	f, err := os.Create(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating the Go file:", err)
		fmt.Fprintln(os.Stderr, "      filename:", filename)
		os.Exit(1)
	}
	fmt.Fprintln(f, "package "+packageName+"_test")
	return f
}

// mustHaveSuffix checks that the file name has the given suffix, it reports
// an error and exits if not
func mustHaveSuffix(filename, suffix string) {
	if !strings.HasSuffix(filename, suffix) {
		fmt.Fprintln(os.Stderr, "Error the file must end with:", suffix)
		fmt.Fprintln(os.Stderr, "      filename:", filename)
		os.Exit(1)
	}
}

// mustNotHaveSuffix checks that the file name has not got the given suffix,
// it reports an error and exits if it does
func mustNotHaveSuffix(filename, suffix string) {
	if strings.HasSuffix(filename, suffix) {
		fmt.Fprintln(os.Stderr, "Error the file must not end with:", suffix)
		fmt.Fprintln(os.Stderr, "      filename:", filename)
		os.Exit(1)
	}
}
