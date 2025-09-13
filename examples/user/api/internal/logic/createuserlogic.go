package logic

import (
	"context"
	"easy-chat/examples/user/api/internal/svc"
	"easy-chat/examples/user/api/internal/types"
	"easy-chat/examples/user/rpc/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req *types.CreateReq) (resp *types.CreateResp, err error) {
	user, err := l.svcCtx.User.CreateUser(l.ctx, &userclient.CreateReq{
		Id:    req.Id,
		Name:  req.Name,
		Phone: req.Phone,
	})
	if err != nil {
		return nil, err
	}

	return &types.CreateResp{Msg: user.Msg}, nil
}
