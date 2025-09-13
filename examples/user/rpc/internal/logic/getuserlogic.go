package logic

import (
	"context"
	"easy-chat/examples/user/rpc/internal/svc"
	"easy-chat/examples/user/rpc/user"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetUser rpc 服务处理逻辑
func (l *GetUserLogic) GetUser(in *user.UserReq) (*user.UserResp, error) {
	out, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &user.UserResp{
		Id:    out.Id,
		Name:  out.Name,
		Phone: out.Phone,
	}, nil

}
