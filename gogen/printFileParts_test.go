package gogen_test

import (
	"bytes"
	"testing"

	"github.com/nickwells/gogen.mod/gogen"
	"github.com/nickwells/testhelper.mod/v2/testhelper"
)

const (
	testDataDir       = "testdata"
	ImportsValsSubDir = "imports"
)

var gfc = testhelper.GoldenFileCfg{
	DirNames:    []string{testDataDir, ImportsValsSubDir},
	Sfx:         "txt",
	UpdFlagName: "upd-imports",
}

func init() {
	gfc.AddUpdateFlag()
}

func TestPrintImports(t *testing.T) {
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
		gfc.Check(t, tc.IDStr(), tc.Name, buf.Bytes())
	}
}
