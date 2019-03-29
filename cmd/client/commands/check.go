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
//Check client public/private key pairs in home directory
func checkServer(cmd *cobra.Command,args []string)error  {
	var key usecase.Key
	err:=key.LoadPublicKey(HomeDir)
	if err!=nil{
		return err

	}
	log.Printf("found %s public key",key.KeyType)
	err=key.LoadPrivateKey(HomeDir)
	if err!=nil{
		return err

	}
	log.Printf("found %s private key",key.KeyType)
	return nil
}
