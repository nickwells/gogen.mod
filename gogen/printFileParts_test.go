package gogen_test

import (
	"bytes"
	"flag"
	"testing"

	"github.com/nickwells/gogen.mod/gogen"
	"github.com/nickwells/testhelper.mod/testhelper"
)

const (
	testDataDir       = "testdata"
	ImportsValsSubDir = "imports"
)

var updateImports = flag.Bool("upd-imports", false,
	"update the files holding the import statements")

func TestPrintImports(t *testing.T) {
	gfc := testhelper.GoldenFileCfg{
		DirNames: []string{testDataDir, ImportsValsSubDir},
		Sfx:      "txt",
	}

	testCases := []struct {
		testhelper.ID
		imports []string
	}{
		{
			ID:      testhelper.MkID("oneImport-stdlib"),
			imports: []string{"xx"},
		},
		{
			ID:      testhelper.MkID("oneImport-other"),
			imports: []string{"xx.y"},
		},
		{
			ID:      testhelper.MkID("multiImport-stdlib"),
			imports: []string{"xx", "aa", "mm", "bb"},
		},
		{
			ID:      testhelper.MkID("multiImport-other"),
			imports: []string{"xx.y", "aa.com", "mm.z", "bb.x"},
		},
		{
			ID: testhelper.MkID("multiImport-both"),
			imports: []string{
				"xx.y", "xx", "aa.com", "aa", "mm", "bb", "mm.z", "bb.x",
			},
		},
	}

	for _, tc := range testCases {
		buf := new(bytes.Buffer)
		gogen.PrintImports(buf, tc.imports...)
		testhelper.CheckAgainstGoldenFile(t, tc.IDStr(), buf.Bytes(),
			gfc.PathName(tc.ID.Name), *updateImports)
	}
}
