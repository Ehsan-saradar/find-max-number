package grpc
import (
	"context"
	"github.com/Ehsan-saradar/find-max-number/math/delivery/grpc/math_grpc"
	"google.golang.org/grpc"
	"io"
	"log"
	"github.com/Ehsan-saradar/find-max-number/math/usecase"
	"strconv"
)
type client struct {
	inputUseCase  usecase.InputUseCase
	cryptoUseCase usecase.CryptoUseCase
	done          chan bool
	max           int32
	stream        math_grpc.MathHandler_FindMaxNumberClient
	verbos bool
}
// Connect to gRpc server
func NewMathClientGrpc(clientConnect *grpc.ClientConn, inputUseCase usecase.InputUseCase,cryptoUseCase usecase.CryptoUseCase,verbos bool) {
	var mathClient client
	var err error
	mathClient.cryptoUseCase=cryptoUseCase
	mathClient.inputUseCase=inputUseCase
	mathClient.verbos=verbos
	mathClient.done=make(chan bool)
	clientHandler := math_grpc.NewMathHandlerClient(clientConnect)
	mathClient.stream, err = clientHandler.FindMaxNumber(context.Background())
	if err != nil {
		log.Fatalf("open stream error %v", err)
	}
	mathClient.Start()
}
//Start sending and receiving data
func (mathClient *client)Start(){
	go mathClient.sendData()
	go mathClient.receiveData()
	ctx := mathClient.stream.Context()
	go func() {
		<-ctx.Done()
		if err := ctx.Err(); err != nil {
			log.Println(err)
		}
		close(mathClient.done)
	}()
	<-mathClient.done
}
//handel send data
func (mathClient *client)sendData(){
	var err error
	var request math_grpc.FindMaxNumberRequest
	var number int
	for{
		number,err=mathClient.inputUseCase.GetNextNumber()
		if err!=nil{
			if err==io.EOF{
				mathClient.stream.CloseSend()
				return
			}
			log.Fatal("Failed to get next number")
		}
		request.Signature,err=mathClient.cryptoUseCase.Sign(strconv.Itoa(number))
		if err!=nil{
			log.Fatalf("Failed to sign %d, err=%v",number,err)
		}
		request.Number=int32(number)

		if err := mathClient.stream.Send(&request); err != nil {
			log.Fatalf("Failed to send number %d, err=%v",request.Number,err)
		}
		if mathClient.verbos{
			log.Printf("send number %d", request.Number)
		}
	}
	if err := mathClient.stream.CloseSend(); err != nil {
		log.Println(err)
	}
}
//handel data received from server
func (mathClient *client) receiveData(){
	for {
		resp, err := mathClient.stream.Recv()
		if err == io.EOF {
			close(mathClient.done)
			return
		}
		if err != nil {
			log.Fatalf("failed to receive %v", err)
		}
		mathClient.max = resp.MaxNumber
		if mathClient.verbos{
			log.Printf("new max number %d received", mathClient.max)
		}
	}
}