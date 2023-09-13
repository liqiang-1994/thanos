package data

import (
	"account/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"time"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/user")),
	}
}

func (r userRepo) GetUserByPhone(ctx context.Context, phone string) (*biz.User, error) {
	user := &TUser{}
	result := r.data.db.WithContext(ctx).Where("phone = ?", phone).First(user)
	model := &biz.User{
		ID: user.ID,
	}
	return model, result.Error
}

func (r userRepo) CreateUser(ctx context.Context, m *biz.User) (*biz.User, error) {
	model := TUser{
		UserName:   "诗友" + uuid.NewString(),
		Phone:      m.Phone,
		Status:     1,
		Follow:     0,
		Watch:      0,
		Up:         0,
		LoginTime:  time.Now(),
		LoginType:  1,
		CreateTime: time.Now(),
	}
	result := r.data.db.WithContext(ctx).Create(&model)
	return m, result.Error
}

func (u userRepo) GetUserById(ctx context.Context, id string) (*biz.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepo) UpdateUser(ctx context.Context, m *biz.User) (bool, error) {
	//TODO implement me
	panic("implement me")
}

type TUser struct {
	ID         int64     `orm:"column(id)" db:"id" json:"id"`
	UserName   string    `orm:"column(user_name)" db:"user_name" json:"user_name"`
	Address    string    `orm:"column(address)" db:"address" json:"address"`
	Phone      string    `orm:"column(phone)" db:"phone" json:"phone"`
	Gender     int64     `orm:"column(gender)" db:"gender" json:"gender"`
	Signature  string    `orm:"column(signature)" db:"signature" json:"signature"`
	Avatar     string    `orm:"column(avatar)" db:"avatar" json:"avatar"`
	Status     int64     `orm:"column(status)" db:"status" json:"status"`
	Follow     int64     `orm:"column(follow)" db:"follow" json:"follow"`
	Watch      int64     `orm:"column(watch)" db:"watch" json:"watch"`
	Up         int64     `orm:"column(up)" db:"up" json:"up"`
	LoginTime  time.Time `orm:"column(login_time)" db:"login_time" json:"login_time"`
	LoginType  int64     `orm:"column(login_type)" db:"login_type" json:"login_type"`
	CreateTime time.Time `orm:"column(create_time)" db:"create_time" json:"create_time"`
	UpdateTime time.Time `orm:"column(update_time)" db:"update_time" json:"update_time"`
}
