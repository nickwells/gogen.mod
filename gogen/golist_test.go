package gogen

import (
	"testing"

	"github.com/nickwells/testhelper.mod/v2/testhelper"
)

func TestRunGoList(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		format string
		expVal string
	}{
		{
			ID:     testhelper.MkID("Name"),
			format: "{{.Name}}",
			expVal: "gogen",
		},
		{
			ID:     testhelper.MkID("Name with spaces around"),
			format: "  \t{{.Name}} \n",
			expVal: "gogen",
		},
	}

	for _, tc := range testCases {
		val := runGoListOrDie(tc.format)
		testhelper.DiffString(t, tc.IDStr(), "package name", val, tc.expVal)
	}
}
