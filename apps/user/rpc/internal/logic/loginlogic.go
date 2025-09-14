package logic

import (
	"context"
	"easy-chat/apps/user/model"
	"easy-chat/apps/user/rpc/internal/svc"
	"easy-chat/apps/user/rpc/user"
	"easy-chat/pkg/auth"
	"errors"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Login 登录
func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginResp, error) {
	// todo: add your logic here and delete this line
	users, err := l.svcCtx.UserModel.FindOneByPhone(l.ctx, in.Phone)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errors.New("手机号未注册")
	}

	// 生成token
	secret := l.svcCtx.Config.Jwt.AccessSecret
	expire := l.svcCtx.Config.Jwt.AccessExpire

	token, err := auth.GenerateToken(secret, expire, users.Id)

	if err != nil {
		return nil, err
	}

	now := time.Now().Unix()

	return &user.LoginResp{Token: token, Expire: int32(now + expire)}, nil
}
