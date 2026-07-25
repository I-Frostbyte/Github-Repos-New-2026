package users

import (
	"context"

	"github.com/I-Frostbyte/Github-Repos-New-2026/Projects/2-gRPC_Simple_Inventory_Tracker/grpc/go/usersgrpc"
)

type Impl struct {
	err error

	usersgrpc.UnimplementedUsersServiceServer
}

func NewUsersService() *Impl {
	var err error
	return &Impl{
		err: err,
	}
}


func (u *Impl) CreateUser(ctx context.Context, req *usersgrpc.SignUpRequest) (*usersgrpc.SignUpResponse, error) {
	panic ("unimplemented")
}