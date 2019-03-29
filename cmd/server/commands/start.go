package commands

import (
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
	mathGRPC "github.com/Ehsan-saradar/find-max-number/math/delivery/grpc"
	"github.com/Ehsan-saradar/find-max-number/math/usecase"
)
var (
	port int
	host string
	protocol string
)

var StartServerCmd=&cobra.Command{
	Use:"start",
	Short:"Start server",
	RunE:startServer,
}
//Initialize start command flags
func init()  {
	StartServerCmd.Flags().IntVarP(&port,"port","p",3000 , "The port for the server to listen on")
	StartServerCmd.Flags().StringVarP(&host,"server","s", "127.0.0.1", "Server address")
	StartServerCmd.Flags().StringVarP(&protocol,"protocol","", "tcp", "Protocol to use (tcp|tcp4|tcp6|unix)")
}
//Start gRPC server
func startServer(cmd *cobra.Command,args []string) error  {
	var err error
	//load key from home directory
	var key usecase.Key
	cryptoUseCase,err:=key.Load(HomeDir,false)
	if err != nil {
		return err
	}

	//start listening
	lis, err := net.Listen(protocol, host+":"+strconv.Itoa(port))
	if err != nil {
		return err
	}else{
		log.Printf("start listening to %d",port)
	}
	grpcServer := grpc.NewServer()
	mathGRPC.NewMathServerGrpc(grpcServer,usecase.NewSimpleMathUseCase(),cryptoUseCase,Verbose)
	if err := grpcServer.Serve(lis); err != nil {
		return  err
	}
	return nil
}
