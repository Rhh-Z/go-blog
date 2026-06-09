package service

import (
	"context"
	v1 "kratos-realworld/api/realworld/v1"
)

func (s *RealWorldService) Login(ctx context.Context, res *v1.LoginRequest) (req *v1.LoginResponse, err error) {
	return &v1.LoginResponse{
		User: &v1.LoginResponse_User{
			Username: "jack",
		},
	}, nil
}
