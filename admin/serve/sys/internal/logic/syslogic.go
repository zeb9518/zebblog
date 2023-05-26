package logic

import (
	"context"

	"sys/internal/svc"
	"sys/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysLogic {
	return &SysLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysLogic) Sys(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
