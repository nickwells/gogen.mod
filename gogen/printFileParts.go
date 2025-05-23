package gogen

import (
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/nickwells/param.mod/v6/param"
	"github.com/nickwells/param.mod/v6/phelp"
)

// PrintPreambleOrDie prints the introductory comments for the file to be
// generated.
//
// Deprecated: Use PrintPreamble. This function will not call os.Exit and so
// is misnamed.
func PrintPreambleOrDie(f io.Writer, ps *param.PSet) {
	PrintPreamble(f, ps)
}

// PrintPreamble prints the introductory comments for the file to be
// generated. Note that the package name is not written by this but when the
// file is created; this is because the go list command fails if any go file
// is empty when it runs.
func PrintPreamble(f io.Writer, ps *param.PSet) {
	fmt.Fprintln(f,
		"\n// Code generated by "+ps.ProgBaseName()+"; DO NOT EDIT.")

	paramsSetAt := ""

	for _, g := range ps.GetGroups() {
		if strings.HasPrefix(g.Name(), phelp.CommonParamsGroupNamePrefix()) {
			continue
		}

		for _, p := range g.Params() {
			whereSet := p.WhereSet()
			for _, ws := range whereSet {
				paramsSetAt += "//\t" + ws + "\n"
			}
		}
	}

	if paramsSetAt != "" {
		fmt.Fprint(f, "// with parameters set at:\n"+paramsSetAt)
	}
}

// PrintImports adds the import statements, if any, one per line to the file
func PrintImports(f io.Writer, imps ...string) {
	if len(imps) == 0 {
		return
	}

	if len(imps) == 1 {
		fmt.Fprintf(f, "import %q\n", imps[0])
		return
	}

	stdLib, other := partitionImports(imps)

	fmt.Fprintln(f, "import (")
	addUniqueImport(f, stdLib)

	if len(stdLib) > 0 {
		fmt.Fprintln(f)
	}

	addUniqueImport(f, other)
	fmt.Fprintln(f, ")")
}

// partitionImports splits the slice of imports into std libs and others and
// sorts them read to be written out.
func partitionImports(imps []string) ([]string, []string) {
	var stdLib, other []string

	for _, v := range imps {
		if v == "" {
			continue
		}

		_, imp, ok := strings.Cut(v, "=")
		if !ok {
			imp = v
		}

		first, _, _ := strings.Cut(imp, "/")
		if strings.ContainsRune(first, '.') {
			other = append(other, v)
		} else {
			stdLib = append(stdLib, v)
		}
	}

	sort.Strings(stdLib)
	sort.Strings(other)

	return stdLib, other
}

// addUniqueImport writes the import entries. It suppresses duplicate
// entries.
func addUniqueImport(f io.Writer, imps []string) {
	prev := ""

	for _, v := range imps {
		if v == prev {
			continue
		}

		id, imp, ok := strings.Cut(v, "=")
		if ok {
			fmt.Fprintf(f, "\t%s %q\n", id, imp)
		} else {
			fmt.Fprintf(f, "\t%q\n", v)
		}

		prev = v
	}
}
