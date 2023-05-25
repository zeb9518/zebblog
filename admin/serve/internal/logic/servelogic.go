package logic

import (
	"context"

	"serve/internal/svc"
	"serve/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ServeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewServeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ServeLogic {
	return &ServeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ServeLogic) Serve(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
   return &types.Response{
        Message: "Hello go-zero",
    }, nil
}
