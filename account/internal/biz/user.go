package biz

import (
	v1 "account/api/account/v1"
	"account/internal/conf"
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type CreateAccount struct {
	Phone string
	Code  string
}

type User struct {
	ID        int64
	Uid       string
	UserName  string
	Address   string
	Phone     string
	Gender    int64
	Signature string
	Avatar    string
	Follow    int64
	Watch     int64
	Up        int64
	LoginType int64
}

type UserRepo interface {
	GetUserByPhone(ctx context.Context, phone string) (*User, error)
	CreateUser(ctx context.Context, m *User) (*User, error)
	GetUserById(ctx context.Context, id string) (*User, error)
	UpdateUser(ctx context.Context, m *User) (bool, error)
}

type UserUseCase struct {
	key  string
	repo UserRepo
	log  *log.Helper
}

func NewUserCase(conf *conf.Auth, repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		key:  conf.ApiKey,
		repo: repo,
		log:  log.NewHelper(log.With(logger, "account", "usecase/user")),
	}
}

func (u *UserUseCase) Login(ctx context.Context, req *CreateAccount) (*v1.ResponseModel, error) {
	result, err := u.repo.GetUserByPhone(ctx, req.Phone)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		result, err = u.repo.CreateUser(ctx, &User{Phone: req.Phone})
	}
	if err != nil {
		log.Infof("account login insert new user err %v", err)
		return nil, v1.ErrorLoginFailed("Login failed: %s", err.Error())
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       result.ID,
		"uid":      result.Uid,
		"userName": result.UserName,
		"address":  result.Address,
		"avatar":   result.Avatar,
	})
	signStr, err := claims.SignedString([]byte(u.key))
	if err != nil {
		log.Infof("account login generate sign err %v", err)
		return nil, v1.ErrorLoginFailed("Login failed: %s", err.Error())
	}
	return &v1.ResponseModel{
		Code:     200,
		Reason:   "",
		Message:  "OK",
		Metadata: map[string]string{"sign": signStr},
	}, nil
}

func (u *UserUseCase) UpdateUser(ctx context.Context, model *User) (bool, error) {
	//claims, ok := kjwt.FromContext(ctx);
	//ok{}
	return u.repo.UpdateUser(ctx, model)
}

func (u *UserUseCase) GetUserById(ctx context.Context, id string) (*User, error) {
	return u.repo.GetUserById(ctx, id)
}
