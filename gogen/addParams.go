package gogen

import (
	"errors"
	"fmt"

	"github.com/nickwells/check.mod/check"
	"github.com/nickwells/param.mod/v3/param"
	"github.com/nickwells/param.mod/v3/param/psetter"
)

// AddParams will add standard parameters to the passed ParamSet.
//
// Note that if the filename has no default value then the filename parameter
// must be set if the file is to be created (if, after parameters are set,
// makeFile is true). If it does have a value then the option will not be
// shown in the standard (abridged) help message.
//
// Note also that the makeFile parameter must be set to true. This will panic
// if it is false.
func AddParams(fileName *string, makeFile *bool) func(ps *param.PSet) error {
	if !*makeFile {
		panic(errors.New(
			"The makeFile parameter must point to a value set to true"))
	}

	fileNameOptAttr := param.DontShowInStdUsage
	if *fileName == "" {
		fileNameOptAttr = 0
	}

	return func(ps *param.PSet) error {
		fileNameParam := ps.Add("file-name",
			psetter.Pathname{
				Value: fileName,
				Checks: []check.String{
					check.StringLenGT(3),
					check.StringHasSuffix(".go"),
					check.StringNot(
						check.StringHasSuffix("_test.go"),
						"a test file"),
				},
			},
			"set the name of the output file",
			param.AltName("f"),
			param.Attrs(fileNameOptAttr),
		)

		noFileParam := ps.Add("no-file",
			psetter.Bool{
				Invert: true,
				Value:  makeFile,
			},
			"don't create the go file, instead just print the content to"+
				" standard out. This is useful for debugging or just to "+
				"see what would have been produced",
			param.Attrs(param.DontShowInStdUsage),
		)

		ps.AddFinalCheck(func() error {
			if fileNameParam.HasBeenSet() && noFileParam.HasBeenSet() {
				return fmt.Errorf(
					"only one of %q and %q may be set at the same time",
					fileNameParam.Name(), noFileParam.Name())
			}
			if *makeFile && *fileName == "" {
				return fmt.Errorf(
					"if the file is to be made the name must be set,"+
						" use %q to set it",
					fileNameParam.Name())
			}
			return nil
		})

		return nil
	}
}
