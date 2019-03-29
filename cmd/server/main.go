package main

import (
	"github.com/Ehsan-saradar/find-max-number/cmd/server/commands"
	"os"
)

func main(){
	cmd:=commands.RootCmd
	cmd.AddCommand(commands.StartServerCmd)
	cmd.AddCommand(commands.CheckServerCmd)
	cmd.AddCommand(commands.VersionCmd)
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}


