package logic

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"serve/sys/internal/svc"
	"serve/sys/internal/types"
	"serve/sys/internal/utils"
	"serve/sys/model"
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
func (l *SysUserLogic) Login(req types.LoginReq) (*types.LoginRep, error) {
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

// 注册 内部系统注册不开放api 需要在test中测试
func (l *SysUserLogic) Register(ctx context.Context, req types.RegisterReq) (*types.RegisterRep, error) {
	// 查询用户信息
	user, err := l.svcCtx.SysUser.FindOneByUsername(ctx, req.Username)
	if err != nil && err != sql.ErrNoRows {
		// 检查用户是否存在
		if user != nil && user.Username.Valid {
			return nil, fmt.Errorf("用户已存在")
		}
		return nil, err
	}
	var newUser = &model.SysUser{}
	// 生成uuid
	newUser.Id = utils.GenerateUUID()
	newUser.Username = sql.NullString{String: req.Username, Valid: true}
	// 密码加盐
	hashedPassword, salt := utils.EncryptPassword(req.Password)
	newUser.Password = sql.NullString{String: hashedPassword, Valid: true}
	newUser.Salt = sql.NullString{String: salt, Valid: true}
	// 注册用户
	res, err := l.svcCtx.SysUser.Insert(ctx, newUser)
	if err != nil {
		return nil, err
	}
	fmt.Printf("注册用户: %v", res)
	return nil, nil
}
