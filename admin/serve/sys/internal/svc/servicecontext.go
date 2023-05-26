package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"serve/sys/internal/config"
	"serve/sys/model"
)

type ServiceContext struct {
	Config  config.Config
	SysUser model.SysUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	um := model.NewSysUserModel(conn, c.CacheRedis)
	return &ServiceContext{
		Config:  c,
		SysUser: um,
	}
}
