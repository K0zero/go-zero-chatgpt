package user

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-chatgpt/common/ctxdata"
	"go-zero-chatgpt/user/internal/model"
	"go-zero-chatgpt/user/internal/svc"
	"go-zero-chatgpt/user/internal/types"
	"time"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// todo: add your logic here and delete this line
	//var user model.User

	User := &model.User{}
	User.Mobile = req.Mobile
	User.Username = req.UserName
	User.Password = req.Password

	SqlResult, err := l.svcCtx.UserModel.Insert(l.ctx, User)

	Res, err := SqlResult.LastInsertId()

	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	accessToken, err := l.getJwtToken(l.svcCtx.Config.JwtAuth.AccessSecret, now, accessExpire, Res)
	if err != nil {
		return nil, err
	}
	return &types.RegisterResp{
		AccessToken:  accessToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}

func (l *RegisterLogic) getJwtToken(secretKey string, iat, seconds, Id int64) (string, error) {

	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims[ctxdata.CtxKeyJwtUserId] = Id
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
