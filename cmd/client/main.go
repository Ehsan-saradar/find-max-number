package main

import (
	"github.com/Ehsan-saradar/find-max-number/cmd/client/commands"
	"os"
)

func main() {
	cmd:=commands.RootCmd
	cmd.AddCommand(commands.InitFilesCmd)
	cmd.AddCommand(commands.VersionCmd)
	cmd.AddCommand(commands.StartServerCmd)
	cmd.AddCommand(commands.CheckServerCmd)
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
