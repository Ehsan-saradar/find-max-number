package commands

import (
	"github.com/Ehsan-saradar/find-max-number/version"
	"github.com/spf13/cobra"
	"log"
)
//Get version of server
var VersionCmd=&cobra.Command{
	Use:"version",
	Short:"Get version of server",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println(version.Version)
	},
}
