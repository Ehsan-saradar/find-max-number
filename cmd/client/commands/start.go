package commands

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
	"strconv"
	mathGRPC "github.com/Ehsan-saradar/find-max-number/math/delivery/grpc"
	"github.com/Ehsan-saradar/find-max-number/math/usecase"
	"github.com/Ehsan-saradar/find-max-number/math/repository"
)
var (
	port int
	host string
	protocol string
	inputType string
)

var StartServerCmd=&cobra.Command{
	Use:"start",
	Short:"Start server",
	RunE:startClient,
}
//Initialize client command flags
func init()  {
	StartServerCmd.Flags().IntVarP(&port,"port","p",3000 , "The port for the server to listen on")
	StartServerCmd.Flags().StringVarP(&host,"server","s", "127.0.0.1", "Server address")
	StartServerCmd.Flags().StringVarP(&protocol,"protocol","", "tcp", "Protocol to use (tcp|tcp4|tcp6|unix)")
	StartServerCmd.Flags().StringVarP(&inputType,"input","i", "file", "Client input type to use (console|file ("+repository.InputFileName+"located in home directory))")
}
//Start client app and connect to gRPC server
func startClient(cmd *cobra.Command,args []string) error  {
	var err error
	//load key from home directory
	var key usecase.Key
	var mathRepository repository.MathRepository
	cryptoUseCase,err:=key.Load(HomeDir,true)
	if err != nil {
		return err
	}
	switch inputType {
	case "console":
		mathRepository = repository.NewMathConsoleRepository()
		break
	case "file":
		mathRepository, err = repository.NewMathFileRepository(HomeDir)
		if err != nil {
			return err
		}
	default:
		return errors.New("Invalid client input type")

	}
	fmt.Println(protocol +"://"+host+":"+strconv.Itoa(port))
	conn, err := grpc.Dial(host+":"+strconv.Itoa(port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}
	mathGRPC.NewMathClientGrpc(conn,usecase.NewSimpleInputUseCase(mathRepository),cryptoUseCase,Verbose)
	return nil
}
