package logic

import (
	"context"
	"easy-chat/examples/user/model"

	"easy-chat/examples/user/rpc/internal/svc"
	"easy-chat/examples/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateUser  创建用户信息
func (l *CreateUserLogic) CreateUser(in *user.CreateReq) (*user.CreateResp, error) {
	_, err := l.svcCtx.UserModel.Insert(l.ctx, &model.Users{
		Id:    in.Id,
		Name:  in.Name,
		Phone: in.Phone,
	})

	if err != nil {
		return nil, err
	}

	return &user.CreateResp{Msg: "ok"}, nil
}
