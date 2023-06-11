package user

import (
	"context"

	"go-zero-chatgpt/user/internal/svc"
	"go-zero-chatgpt/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	// todo: add your logic here and delete this line
	UserInfo, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, req.Mobile)
	if err != nil {
		return nil, err
	}
	return &types.UserInfoResp{
		Id:         UserInfo.Id,
		Mobile:     UserInfo.Mobile,
		UserName:   UserInfo.Username,
		CreateTime: UserInfo.CreateTime.GoString(),
	}, nil

}
