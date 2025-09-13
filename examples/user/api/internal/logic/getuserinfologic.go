package logic

import (
	"context"
	"easy-chat/examples/user/api/internal/svc"
	"easy-chat/examples/user/api/internal/types"
	"easy-chat/examples/user/rpc/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.UserReq) (resp *types.UserResp, err error) {

	l.Infof("GetUserInfo")
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
