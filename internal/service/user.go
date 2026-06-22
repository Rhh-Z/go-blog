package service

import (
	"context"

	pb "kratos-realworld/api/realworld/v1"
	v1 "kratos-realworld/api/realworld/v1"
	"kratos-realworld/internal/biz"
	"kratos-realworld/internal/errors"
)

type UserService struct {
	pb.UnimplementedUserServer
	uc *biz.UserUsecase
}

func NewUserService(uc *biz.UserUsecase) *UserService {
	return &UserService{uc: uc}
}

func (userService *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	if len(req.User.Email) == 0 {
		return nil, errors.NewHttpError(422, "email", "email can't not be empty")
	}
	return &v1.LoginResponse{
		User: &v1.LoginResponse_User{
			Username: "asd",
			Email:    "1",
			Password: "111",
			Bio:      "qwe",
			Image:    "123",
		},
	}, nil
}

func (userService *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	user, err := userService.uc.Register(ctx, req.User.Email, req.User.Username, req.User.Password)
	if err != nil {
		return nil, err
	}

	return &pb.CreateUserReply{
		Username: user.Username,
		Token:    user.Token,
	}, nil
}
func (userService *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	return &pb.UpdateUserReply{}, nil
}
func (userService *UserService) GetUserByEmail(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	return &pb.GetUserReply{}, nil
}
func (userService *UserService) UserList(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	return &pb.ListUserReply{}, nil
}
