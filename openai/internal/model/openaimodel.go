package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OpenaiModel = (*customOpenaiModel)(nil)

type (
	// OpenaiModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOpenaiModel.
	OpenaiModel interface {
		openaiModel
	}

	customOpenaiModel struct {
		*defaultOpenaiModel
	}
)

// NewOpenaiModel returns a model for the database table.
func NewOpenaiModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) OpenaiModel {
	return &customOpenaiModel{
		defaultOpenaiModel: newOpenaiModel(conn, c, opts...),
	}
}
