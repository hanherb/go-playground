package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	u "github.com/hanherb/go-playground/grpc-gen"
	"github.com/hanherb/go-playground/src/config"
	"github.com/hanherb/go-playground/src/controllers"

	"google.golang.org/grpc"
)

func main() {
	config.MysqlInitialization()

	errChan := make(chan error)
	defer close(errChan)

	go startGRPCServer(errChan)

	<-errChan
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

	u.RegisterUserServiceServer(s, &controllers.UserUnimplemented{})

	// Running GRPC Server
	if err := s.Serve(lis); err != nil {
		errChan <- err
		log.Fatalf("failed to serve: %s", err)
	}
}
