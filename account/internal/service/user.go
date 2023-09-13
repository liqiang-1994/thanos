package service

import (
	api1 "account/api/account/v1"
	"account/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type UserService struct {
	api1.UnimplementedUserServer
	uc  *biz.UserUseCase
	log *log.Helper
}

func NewUserService(uc *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "account", "service/user")),
	}
}

func (u UserService) Login(ctx context.Context, model *api1.LoginRequest) (*api1.ResponseModel, error) {
	m := &biz.CreateAccount{
		Code:  model.Code,
		Phone: model.Phone,
	}
	result, err := u.uc.Login(ctx, m)
	if err != nil {
		return nil, err
	}
	return result, nil
}
