package biz

import (
	"context"
	v1 "kratos-realworld/api/realworld/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// User is a User model.
type User struct {
	ID       int64  `bson:"id"`
	Email    string `bson:"name"`
	Username string `bson:"username"`
	Bio      string `bson:"bio"`
	Image    string `bson:"image"`
	Password string `bson:"password"`
}

// UserRepo is a user repo.
type UserRepo interface {
	SaveUser(context.Context, *User) (*User, error)
	UpdateUser(context.Context, int64, *User) (*User, error)
	GetUser(context.Context, int64) (*User, error)
	ListByHello(context.Context, string) ([]*User, error)
	UserList(context.Context) ([]*User, error)
}

// UserUsecase is a User usecase.  加日志
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) CreateUser(ctx context.Context, user *User) (*User, error) {
	uc.log.WithContext(ctx).Infof("CreateUser: %v", user.Username)
	return uc.repo.SaveUser(ctx, user)
}

func (uc *UserUsecase) GetUser(ctx context.Context, id int64, user *User) (*User, error) {
	uc.log.WithContext(ctx).Infof("GetUser: %v", user.Username)
	return uc.repo.GetUser(ctx, id)
}

func (uc *UserUsecase) UpdateUser(ctx context.Context, id int64, user *User) (*User, error) {
	uc.log.WithContext(ctx).Infof("UpdateUser: %v", user.Username)
	return uc.repo.UpdateUser(ctx, id, user)
}
