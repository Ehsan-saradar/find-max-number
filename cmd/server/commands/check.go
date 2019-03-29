package commands

import (
	"github.com/Ehsan-saradar/find-max-number/math/usecase"
	"github.com/spf13/cobra"
	"log"
)

var CheckServerCmd=&cobra.Command{
	Use:"check",
	Short:"Check server",
	RunE:checkServer,
}
//Check public key file in home directory
func checkServer(cmd *cobra.Command,args []string)error  {
	var key usecase.Key
	err:=key.LoadPublicKey(HomeDir)
	if err!=nil{
		return err

	}
	log.Printf("found %s public key",key.KeyType)
	return nil
}
