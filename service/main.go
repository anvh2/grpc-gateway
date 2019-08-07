// package main

// import (
// 	"context"
// 	"log"
// 	"net"

// 	pb "grpc/service/proto"

// 	"google.golang.org/grpc"
// )

// //Server ...
// type Server struct{}

// //Run ...
// func (s *Server) Run() error {
// 	lis, err := net.Listen("tcp", ":8000")

// 	if err != nil {
// 		log.Fatal("Failed to listen: %v", err)
// 		return err
// 	}
// 	defer lis.Close()

// 	grpcServer := grpc.NewServer()

// 	pb.RegisterGretterServiceServer(grpcServer, s)

// 	if err := grpcServer.Serve(lis); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (s *Server) Say(ctx context.Context, req *pb.Request) (*pb.Response, error) {
// 	return &pb.Response{
// 		Message: "Hello" + req.GetName(),
// 	}, nil
// }

// func main() {
// 	server := Server{}
// 	server.Run()
// }

package main

import (
	"context"
	pb "grpc/service/proto"
	"log"
	"time"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
)

//Say ...
type Say struct{}

//Say ...
func (s *Say) Say(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	log.Print("Received Say.Hello request")
	rsp.Message = "Hello " + req.Name
	return nil
}

func main() {
	service := grpc.NewService(
		micro.Name("go.micro.srv.greeter"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// optionally setup command line usage
	service.Init()

	// Register Handlers
	pb.RegisterGretterServiceHandler(service.Server(), new(Say))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
