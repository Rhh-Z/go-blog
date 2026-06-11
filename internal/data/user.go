package data

import (
	"context"
	"kratos-realworld/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type UserRepo struct {
	data *Data // *Data是连接数据库客户端
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &UserRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *UserRepo) Save(ctx context.Context, user *biz.User) (*biz.User, error) {
	return user, nil
}

func (r *UserRepo) Update(ctx context.Context, user *biz.User) (*biz.User, error) {
	return user, nil
}

func (r *UserRepo) FindByID(context.Context, int64) (*biz.User, error) {
	// var user = biz.User
	return nil, nil
}

func (r *UserRepo) ListByHello(context.Context, string) ([]*biz.User, error) {
	return nil, nil
}

func (r *UserRepo) ListAll(context.Context) ([]*biz.User, error) {
	return nil, nil
}
