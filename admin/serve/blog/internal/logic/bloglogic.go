package logic

import (
	"context"

	"serve/blog/internal/svc"
	"serve/blog/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BlogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBlogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BlogLogic {
	return &BlogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BlogLogic) Blog(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
