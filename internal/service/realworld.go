package service

import (
	v1 "kratos-realworld/api/realworld/v1"
	"kratos-realworld/internal/biz"
)

// RealWorld is a realworld service.
type RealWorldService struct {
	v1.UnimplementedRealWorldServer

	uc *biz.RealWorldUsecase
}

// NewRealWorld new a realworld service.
func NewRealWorld(uc *biz.RealWorldUsecase) *RealWorldService {
	return &RealWorldService{uc: uc}
}
