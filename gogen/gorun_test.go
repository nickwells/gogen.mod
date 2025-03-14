package gogen_test

import (
	"os"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/nickwells/gogen.mod/gogen"
	"github.com/nickwells/testhelper.mod/v2/testhelper"
)

func TestExecGoCmdNoExit(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		dir       string
		expResult bool
		expStdout string
		expStderr string
	}{
		{
			ID:        testhelper.MkID("Bad code"),
			dir:       "testdata/code/badCode",
			expResult: false,
			expStderr: "Command failed: go build\n" +
				"         Error: exit status 1\n" +
				"# badCode\n" +
				"./bad.go:7:7: syntax error:" +
				" unexpected name is at end of statement\n" +
				"\n",
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

	editStderr := regexp.MustCompile("Command failed: .*go build")

	for _, tc := range testCases {
		tmpDir, cleanup, err := testhelper.MakeTempDirCopy(tc.dir)
		if err != nil {
			if err := cleanup(); err != nil {
				t.Log("cleanup failed")
				t.Fatal(err)
			}

			t.Error(err)

			continue
		}

		cdOrFatal(t, tmpDir)

		fakeIO, err := testhelper.NewStdioFromString("")
		if err != nil {
			t.Fatal(err)
		}

		buildSucceeded := gogen.ExecGoCmdNoExit(gogen.NoCmdIO, "build")
		testhelper.DiffBool(t, tc.IDStr(), "build status",
			buildSucceeded, tc.expResult)

		if buildSucceeded {
			_ = os.Remove(filepath.Base(tc.dir))
		}

		stdout, stderr, err := fakeIO.Done()
		if err != nil {
			t.Fatal(err)
		}

		stderr = editStderr.ReplaceAll(stderr,
			[]byte("Command failed: go build"))

		testhelper.DiffString(t, tc.IDStr(), "stdout",
			string(stdout), tc.expStdout)
		testhelper.DiffString(t, tc.IDStr(), "stderr",
			string(stderr), tc.expStderr)

		cdOrFatal(t, initialDir)

		if cleanup != nil {
			if err := cleanup(); err != nil {
				t.Log("cleanup failed")
				t.Fatal(err)
			}
		}
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
