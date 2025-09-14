package logic

import (
	"context"

	"easy-chat/examples/mongo/internal/svc"
	"easy-chat/examples/mongo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MongoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMongoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MongoLogic {
	return &MongoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MongoLogic) Mongo(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
