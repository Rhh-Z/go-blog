package data

import (
	"context"
	"fmt"
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

var DataBase = NewMongoDB().Database("kratos-realworld")

func (r *UserRepo) SaveUser(ctx context.Context, user *biz.User) (*biz.User, error) {
	collection := DataBase.Collection("user")
	fmt.Println(collection)
	if collection == nil {
		DataBase.CreateCollection(ctx, "user")
		newCollection := DataBase.Collection("user")
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

func (r *UserRepo) UpdateUser(ctx context.Context, id int64, user *biz.User) error {
	collection := DataBase.Collection("user")
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
	return err
}

func (r *UserRepo) GetUser(context.Context, int64) (*biz.User, error) {
	// var user = biz.User
	return nil, nil
}

func (r *UserRepo) ListByHello(context.Context, string) ([]*biz.User, error) {
	return nil, nil
}

func (r *UserRepo) UserList(context.Context) ([]*biz.User, error) {
	return nil, nil
}
