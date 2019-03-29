package grpc
import (
	"github.com/Ehsan-saradar/find-max-number/math/delivery/grpc/math_grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io"
	"log"
	"strconv"
	"github.com/Ehsan-saradar/find-max-number/math/usecase"
)


type server struct {
	mathUseCase usecase.MathUseCase
	cryptoUseCase usecase.CryptoUseCase
	verbos bool
}
//Start new gRPC server
func NewMathServerGrpc(grpcServer *grpc.Server, mathUseCase usecase.MathUseCase,cryptoUseCase usecase.CryptoUseCase,verbos bool) {
	mathServer := &server{
		mathUseCase:mathUseCase,
		cryptoUseCase:cryptoUseCase,
		verbos:verbos,
	}
	math_grpc.RegisterMathHandlerServer(grpcServer, mathServer)
	reflection.Register(grpcServer)

}

//Handel client which call  FindMaxNumber function
func (s server) FindMaxNumber(srv math_grpc.MathHandler_FindMaxNumberServer) error {
	if s.verbos{
		log.Println("start serving new client")
	}
	ctx := srv.Context()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			req, err := srv.Recv()
			if err == io.EOF {
				log.Println("client closed the stream.")
				s.mathUseCase.Reset()
				return nil
			}
			if err != nil {
				log.Printf("receive error %v", err)
				continue
			}
			if !s.cryptoUseCase.Verify(strconv.Itoa(int(req.Number)), req.Signature){
				if s.verbos{
					log.Printf("invalid signature for %d", req.Number)
				}
				continue
			}
			if s.mathUseCase.IsMax(req.Number) {
				resp := math_grpc.FindMaxNumberResponse{MaxNumber: req.Number}
				if err := srv.Send(&resp); err != nil {
					if s.verbos{
						log.Printf("failed to send max number %d, err=%v",req.Number, err)
					}
				}
				if s.verbos{
					log.Printf("send new max=%d", req.Number)
				}
			}
		}
	}
}
