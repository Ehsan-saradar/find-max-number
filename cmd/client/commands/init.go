package commands

import (
	"errors"
	"github.com/Ehsan-saradar/find-max-number/math/repository"
	"github.com/spf13/cobra"
	"strings"
	"github.com/Ehsan-saradar/find-max-number/math/usecase"
)
var encryptionType string
var InitFilesCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize client keys and data",
	RunE:  initFiles,
}
//Initialize command flags
func init()  {
	InitFilesCmd.Flags().StringVarP(&encryptionType,"encryption","e", "ed25519", "Encryption to use (ed25519|rsa)")
}
//Create public/private key pair and generate random numbers
func initFiles(cmd *cobra.Command, args []string) error {
	var cryptoService usecase.CryptoUseCase
	switch strings.ToLower(encryptionType) {
	case "ed25519":
		cryptoService=new(usecase.Ed25519)
		break
	case "rsa":
		cryptoService=new(usecase.Rsa)
		break
	default:
		return errors.New("Invalid encryption type")
	}
	key,err:=cryptoService.GenerateKey()
	if err!=nil{
		return err
	}
	key.Save(HomeDir)
	repository.GenerateFileRepository(HomeDir)
	return nil
}