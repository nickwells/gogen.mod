package gogen

import (
	"testing"

	"github.com/nickwells/testhelper.mod/testhelper"
)

func TestRunGoList(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		format      string
		expectedVal string
	}{
		{
			ID:          testhelper.MkID("Name"),
			format:      "{{.Name}}",
			expectedVal: "gogen",
		},
		{
			ID:          testhelper.MkID("Name with spaces around"),
			format:      "  \t{{.Name}} \n",
			expectedVal: "gogen",
		},
	}

	for _, tc := range testCases {
		val := runGoListOrDie(tc.format)
		if val != tc.expectedVal {
			t.Log(tc.IDStr())
			t.Errorf("\t: unexpected value returned from RunGoList\n")
		}
	}
}
