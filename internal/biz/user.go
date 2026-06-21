package biz

import (
	"context"
	v1 "kratos-realworld/api/realworld/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// User is a User model.
type User struct {
	Email    string `bson:"name"`
	Username string `bson:"username"`
	Bio      string `bson:"bio"`
	Image    string `bson:"image"`
	Password string `bson:"password"`
}

type UserLogin struct {
	Email    string `bson:"name"`
	Token    string `bson:"token"`
	Username string `bson:"username"`
	Bio      string `bson:"bio"`
	Image    string `bson:"image"`
}

type UserRegister struct {
	Email        string `bson:"name"`
	Token        string `bson:"token"`
	Username     string `bson:"username"`
	Bio          string `bson:"bio"`
	Image        string `bson:"image"`
	Password     string `bson:"password"`
	PasswordHash string `bson:"passwordHash"`
}

// UserRepo is a user repo.
type UserRepo interface {
	CreateUser(context.Context, *User) (*User, error)
	Register(ctx context.Context, email string, username string, password string) (*UserLogin, error)
	Login(ctx context.Context, email string, passwrod string) (*UserLogin, error)
	UpdateUser(context.Context, int64, *User) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	UserList(context.Context) ([]*User, error)
}

// UserUsecase is a User usecase.  加日志
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func hashPassword(pwd string) string {
	b, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func verifyPassword(hashPwd string, pwd string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(pwd)); err != nil {
		return false
	}
	return true
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) Register(ctx context.Context, email string, username string, password string) (*UserLogin, error) {
	u := &User{
		Email:    email,
		Username: username,
		Password: password,
	}

	if _, err := uc.repo.CreateUser(ctx, u); err != nil {
		return nil, err
	}

	return &UserLogin{
		Username: username,
		Token:    "xXX",
	}, nil
}

func (uc *UserUsecase) Login(ctx context.Context, email string, password string) (*User, error) {
	user, err := uc.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	// if !verifyPassword() {}
	return user, err
}

func (uc *UserUsecase) GetUseByEmail(ctx context.Context, email string) (*User, error) {
	uc.log.WithContext(ctx).Infof("GetUserEmail: %v", email)
	return uc.repo.GetUserByEmail(ctx, email)
}

func (uc *UserUsecase) UpdateUser(ctx context.Context, id int64, user *User) (*User, error) {
	uc.log.WithContext(ctx).Infof("UpdateUser: %v", user.Username)
	return uc.repo.UpdateUser(ctx, id, user)
}
