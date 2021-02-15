// +build generate

package gogen

//go:generate mkfunccontrolparamtype -d "This determines whether or not a go subcommand should be run with its output displayed" -t CmdIO -v NoCmdIO -v ShowCmdIO
