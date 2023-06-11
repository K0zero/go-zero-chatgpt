package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-chatgpt/user/internal/config"
	"go-zero-chatgpt/user/internal/model"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
