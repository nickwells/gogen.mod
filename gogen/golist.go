package gogen

import (
	"fmt"
	"os"
	"os/exec"
)

// GetPackage returns the package name - it does this by running the go
// list command and gathering the output
func GetPackage() string {
	out, err := exec.Command("go", "list", "-f", "{{.Name}}").Output()
	if err != nil {
		fmt.Fprint(os.Stderr, "can't run the go list command", err)
		os.Exit(1)
	}
	return string(out)
}

// GetImportPath returns the importPath name - it does this by running the go
// list command and gathering the output
func GetImportPath() string {
	out, err := exec.Command("go", "list", "-f", "{{.ImportPath}}").Output()
	if err != nil {
		fmt.Fprint(os.Stderr, "can't run the go list command", err)
		os.Exit(1)
	}
	return string(out)
}
