package logic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"serve/sys/internal/svc"
	"serve/sys/internal/types"
	"serve/sys/internal/utils"
)

type SysUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysUserLogic {
	return &SysUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 登录
func (l *SysUserLogic) Login(req types.LoginReq) (*types.UserReply, error) {
	// 查询用户信息
	user, err := l.svcCtx.SysUser.FindUser(req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	// 检查用户是否存在
	if user.Username.Valid {
		accessSecret := l.svcCtx.Config.Auth.AccessSecret
		accessExpire := l.svcCtx.Config.Auth.AccessExpire
		// 生成jwt token
		token, _ := utils.GenerateToke(accessSecret, user.Username.String, accessExpire)
		fmt.Sprintf("jwt token: %v%v", user, token)
	}
	return nil, nil
}

// 注册
func (l *SysUserLogic) Register(req types.RegisterReq) (*types.UserReply, error) {
	// 查询用户信息
	user, err := l.svcCtx.SysUser.FindUser(req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	// 生成jwt token
	fmt.Sprintf("jwt token: %v", user)
	return nil, nil
}
