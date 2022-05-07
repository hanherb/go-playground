package models

import (
	userGrpc "github.com/hanherb/go-playground/grpc-gen"
)

func (User) TableName() string {
	return "user"
}

type User struct {
	ID    int32  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	City  string `json:"city"`
}

func (m *User) ToGrpc() *userGrpc.User {
	return &userGrpc.User{
		Id:    m.ID,
		Name:  m.Name,
		Email: m.Email,
		City:  m.City,
	}
}
