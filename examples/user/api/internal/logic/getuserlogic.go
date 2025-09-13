package logic

import (
	"context"
	"easy-chat/examples/user/rpc/userclient"

	"easy-chat/examples/user/api/internal/svc"
	"easy-chat/examples/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetUser 获取用户信息
func (l *GetUserLogic) GetUser(req *types.UserReq) (resp *types.UserResp, err error) {

	user, err := l.svcCtx.User.GetUser(l.ctx, &userclient.UserReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &types.UserResp{
		Id:    user.Id,
		Name:  user.Name,
		Phone: user.Phone,
	}, nil
}
