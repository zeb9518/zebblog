package test

import (
	"context"
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"serve/sys/internal/config"
	"serve/sys/internal/logic"
	"serve/sys/internal/svc"
	"serve/sys/internal/types"
	"testing"
)

var configFile = flag.String("f", "../../etc/sys-api.yaml", "the config file")

// 测试注册
func TestRegister(t *testing.T) {
	var c config.Config
	conf.MustLoad(*configFile, &c)
	// 构造SysUserLogic 实例
	logic := logic.NewSysUserLogic(context.Background(), svc.NewServiceContext(c))
	// 构造RegisterReq
	req := types.RegisterReq{
		Username: "test",
		Password: "123456",
	}
	// 创建一个 context.Context 实例
	ctx := context.Background()

	// 调用 Register 方法
	rep, err := logic.Register(ctx, req)
	// 判断返回结果是否符合预期
	if err != nil {
		t.Errorf("Register failed: %v", err)
	}
	if rep == nil {
		t.Error("Register returned nil")
	}
}
