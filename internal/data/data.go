package data

import (
	"context"
	"fmt"
	"kratos-realworld/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/qiniu/qmgo"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewMongoDB, NewData, NewRealWorldRepo, NewUserRepo)

// Data .
type Data struct {
	// TODO wrapped database client
}

// 连接数据库
func NewMongoDB() (client *qmgo.Client) {
	ctx := context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: "mongodb://localhost:27017"})
	if err != nil {
		fmt.Printf("数据库客户端连接失败!\n")
		return
	}

	// 关闭连接
	// defer func() {
	// 	if err = client.Close(ctx); err != nil {
	// 		panic(err)
	// 	}
	// }()
	return
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{}, cleanup, nil
}
