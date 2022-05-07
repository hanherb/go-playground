package controllers

import grpcService "github.com/hanherb/go-playground/grpc-gen"

type GrpcController struct {
	grpcService.UnimplementedMainServiceServer
}
