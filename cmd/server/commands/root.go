package commands

import (
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var (
	Verbose bool
	HomeDir string
	RootCmd = &cobra.Command{
		Use:   "server",
		Short: "server",
		Long:  `GRPC server`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
)
//Initialize flags for all commands
func init(){
	RootCmd.PersistentFlags().BoolVarP(&Verbose,"verbos","v",true,"verbose output")
	RootCmd.PersistentFlags().StringVarP(&HomeDir,"home","",os.ExpandEnv(filepath.Join("$HOME","GrpcKeys")),"home directory for client key file")
}