//go:build generate

package gogen

//go:generate mkfunccontrolparamtype -d "This determines whether or not a go subcommand should be run with its output displayed. You can also choose to not show IO on failure" -t CmdIO -v NoCmdIO -v ShowCmdIO -v NoCmdFailIO
