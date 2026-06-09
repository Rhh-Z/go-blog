package data

import (
	"context"

	"kratos-realworld/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type RealWorldRepo struct {
	data *Data
	log  *log.Helper
}

// NewRealWorldRepo .
func NewRealWorldRepo(data *Data, logger log.Logger) biz.RealWorldRepo {
	return &RealWorldRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *RealWorldRepo) Save(ctx context.Context, g *biz.RealWorld) (*biz.RealWorld, error) {
	return g, nil
}

func (r *RealWorldRepo) Update(ctx context.Context, g *biz.RealWorld) (*biz.RealWorld, error) {
	return g, nil
}

func (r *RealWorldRepo) FindByID(context.Context, int64) (*biz.RealWorld, error) {
	return nil, nil
}

func (r *RealWorldRepo) ListByHello(context.Context, string) ([]*biz.RealWorld, error) {
	return nil, nil
}

func (r *RealWorldRepo) ListAll(context.Context) ([]*biz.RealWorld, error) {
	return nil, nil
}
