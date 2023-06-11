package openai

import (
	"context"
	"go-zero-chatgpt/common/ctxdata"
	"go-zero-chatgpt/openai/internal/model"

	"go-zero-chatgpt/openai/internal/svc"
	"go-zero-chatgpt/openai/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertLogic {
	return &InsertLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsertLogic) Insert(req *types.InsertAPIReq) (resp *types.InsertAPIResp, err error) {
	// todo: add your logic here and delete this line

	OpenAi := &model.Openai{}
	OpenAi.UserId = ctxdata.GetUidFromCtx(l.ctx)
	OpenAi.Api = req.API

	SqlResult, err := l.svcCtx.OpenaiModel.Insert(l.ctx, OpenAi)

	if err != nil {
		return nil, err
	}

	ID, _ := SqlResult.LastInsertId()
	return &types.InsertAPIResp{
		Id:     ID,
		UserId: ctxdata.GetUidFromCtx(l.ctx),
		API:    req.API,
	}, nil
}
