package controllers

import (
	"context"
	"errors"

	userGrpc "github.com/hanherb/go-playground/grpc-gen"
	"github.com/hanherb/go-playground/src/config"
	"github.com/hanherb/go-playground/src/repositories"
)

type UserUnimplemented struct {
	userGrpc.UnimplementedUserServiceServer
}

func (u *UserUnimplemented) GetOneUser(ctx context.Context, req *userGrpc.UserGetOneRequest) (*userGrpc.UserGetOneResponse, error) {
	if &req.Id == nil {
		return nil, errors.New("id cannot be empty")
	}

	user := repositories.NewUserRepository(config.DB)

	if err := user.Get(ctx, req); err != nil {
		return nil, err
	}

	response := &userGrpc.UserGetOneResponse{
		Data: user.Data().ToGrpc(),
	}

	return response, nil
}

func (u *UserUnimplemented) GetListUser(ctx context.Context, req *userGrpc.UserGetListRequest) (*userGrpc.UserGetListResponse, error) {
	users := repositories.NewUserRepositories(config.DB)

	data, err := users.Get(ctx, req)
	if err != nil {
		return nil, err
	}

	response := &userGrpc.UserGetListResponse{}
	for _, user := range users.Data() {
		response.Data = append(response.Data, user.ToGrpc())
	}

	response.Count = data.Count

	return response, nil
}
