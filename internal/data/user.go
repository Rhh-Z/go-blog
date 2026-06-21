package data

import (
	"context"
	"kratos-realworld/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"go.mongodb.org/mongo-driver/bson"
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

func (r *UserRepo) CreateUser(ctx context.Context, user *biz.User) (*biz.User, error) {
	database := r.data.Client.Database("kratos-realworld")
	collection := database.Collection("user")
	// 没有集合  创建集合再插入数据
	if collection == nil {
		database.CreateCollection(ctx, "user")
		newCollection := database.Collection("user")
		_, err := newCollection.InsertOne(ctx, biz.User{
			Username: user.Username,
			Password: user.Password,
			Email:    user.Email,
		})
		return user, err
	}

	_, err := collection.InsertOne(ctx, biz.User{
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
	})

	return user, err
}

func (r *UserRepo) Register(ctx context.Context, email string, username string, password string) (*biz.UserLogin, error) {
	return nil, nil
}

func (r *UserRepo) Login(ctx context.Context, email string, passwrod string) (*biz.UserLogin, error) {
	return nil, nil
}

func (r *UserRepo) UpdateUser(ctx context.Context, id int64, user *biz.User) (*biz.User, error) {
	database := r.data.Client.Database("kratos-realworld")
	collection := database.Collection("user")

	filter := bson.M{"id": id}
	update := bson.M{
		"$set": bson.M{
			"username": user.Username,
			"email":    user.Email,
			"image":    user.Image,
			"bio":      user.Bio,
		},
	}
	err := collection.UpdateOne(ctx, filter, update)
	return user, err
}

func (r *UserRepo) GetUserByEmail(ctx context.Context, email string) (*biz.User, error) {
	// var user = biz.User
	return nil, nil
}

func (r *UserRepo) UserList(context.Context) ([]*biz.User, error) {
	return nil, nil
}
