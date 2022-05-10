package gogen_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/nickwells/gogen.mod/gogen"
	"github.com/nickwells/testhelper.mod/v2/testhelper"
)

func TestExecGoCmdNoExit(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		dir       string
		expResult bool
	}{
		{
			ID:        testhelper.MkID("Bad code"),
			dir:       "testdata/code/badCode",
			expResult: false,
		},
		{
			ID:        testhelper.MkID("Good code"),
			dir:       "testdata/code/goodCode",
			expResult: true,
		},
	}

	initialDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Couldn't get the initial working directory: %v ", err)
	}

	for _, tc := range testCases {
		cdOrFatal(t, tc.dir)

		buildSucceeded := gogen.ExecGoCmdNoExit(gogen.NoCmdIO, "build")
		testhelper.DiffBool(t, tc.IDStr(), "", buildSucceeded, tc.expResult)
		if buildSucceeded {
			_ = os.Remove(filepath.Base(tc.dir))
		}

		cdOrFatal(t, initialDir)
	}
}

// cdOrFatal tries to chdir to the given directory and will report a fatal
// error if it cannot.
func cdOrFatal(t *testing.T, dir string) {
	t.Helper()

	err := os.Chdir(dir)
	if err != nil {
		t.Fatalf("Couldn't chdir to %q: %v", dir, err)
	}
}
