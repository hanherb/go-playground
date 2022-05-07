package controllers

import (
	"context"
	"errors"

	grpcService "github.com/hanherb/go-playground/grpc-gen"
	"github.com/hanherb/go-playground/src/config"
	"github.com/hanherb/go-playground/src/repositories"
)

func (g *GrpcController) GetOneUser(ctx context.Context, req *grpcService.UserGetOneRequest) (*grpcService.UserGetOneResponse, error) {
	if &req.Id == nil {
		return nil, errors.New("id cannot be empty")
	}

	user := repositories.NewUserRepository(config.DB)

	if err := user.Get(ctx, req); err != nil {
		return nil, err
	}

	response := &grpcService.UserGetOneResponse{
		Data: user.Data().ToGrpc(),
	}

	return response, nil
}

func (g *GrpcController) GetListUser(ctx context.Context, req *grpcService.UserGetListRequest) (*grpcService.UserGetListResponse, error) {
	users := repositories.NewUserRepositories(config.DB)

	data, err := users.Get(ctx, req)
	if err != nil {
		return nil, err
	}

	response := &grpcService.UserGetListResponse{}
	for _, user := range users.Data() {
		response.Data = append(response.Data, user.ToGrpc())
	}

	response.Count = data.Count

	return response, nil
}
