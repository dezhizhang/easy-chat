package logic

import (
	"context"
	"database/sql"
	"easy-chat/apps/user/model"
	"easy-chat/apps/user/rpc/internal/svc"
	"easy-chat/apps/user/rpc/user"
	"easy-chat/pkg/auth"
	"easy-chat/pkg/encrypt"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Register 注册
func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterResp, error) {

	resp, err := l.svcCtx.UserModel.FindOneByPhone(l.ctx, in.Phone)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, err
	}

	if resp != nil {
		return nil, errors.New("手机号已注册")
	}

	// 定义数据
	users := &model.Users{
		Id:       "121",
		Avatar:   in.Avatar,
		Phone:    in.Phone,
		Username: in.Username,
		Password: sql.NullString{
			String: encrypt.Md5(in.Password),
			Valid:  true,
		},
		Gender: sql.NullInt64{
			Int64: int64(in.Gender),
			Valid: true,
		},
	}

	_, err = l.svcCtx.UserModel.Insert(l.ctx, users)
	if err != nil {
		return nil, err
	}

	secret := l.svcCtx.Config.Jwt.AccessSecret
	expire := l.svcCtx.Config.Jwt.AccessExpire

	token, err := auth.GenerateToken(secret, expire, users.Id)

	if err != nil {
		return nil, err
	}

	now := time.Now().Unix()

	return &user.RegisterResp{Token: token, Expire: int32(now + expire)}, nil
}
