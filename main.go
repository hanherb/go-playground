package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	u "github.com/hanherb/go-playground/grpc-gen/user_grpc.pb.go"

	"google.golang.org/grpc"
)

type server struct {
	u.UnimplementedUserServiceServer
}

func main() {

}

func startGRPCServer(errChan chan error) {
	port := flag.Int("port", 8080, "port for gRPC server listen")
	flag.Parse()

	// Listen on port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen grpc server: %s", err)
	}
	log.Printf("Listen gRPC server on %d port \n", *port)

	// Define GRPC Server
	s := grpc.NewServer()

	u.RegisterUserServiceServer(s, &server{})

	// Running GRPC Server
	if err := s.Serve(lis); err != nil {
		errChan <- err
		log.Fatalf("failed to serve: %s", err)
	}
}
