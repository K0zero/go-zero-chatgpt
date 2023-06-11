package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-chatgpt/openai/internal/config"
	"go-zero-chatgpt/openai/internal/model"
)

type ServiceContext struct {
	Config      config.Config
	OpenaiModel model.OpenaiModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		OpenaiModel: model.NewOpenaiModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
