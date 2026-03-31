package boot

import (
	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/registry/file/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gsel"
	"github.com/gogf/gf/v2/os/gctx"
)

func Init() {
	ctx := gctx.GetInitCtx()

	// 1. 注册中心 - 根据配置选择 file 或 etcd
	registryType := g.Cfg().MustGet(ctx, "registry.type", "file").String()
	switch registryType {
	case "etcd":
		address := g.Cfg().MustGet(ctx, "registry.address", "127.0.0.1:2379").String()
		grpcx.Resolver.Register(etcd.New(address))
		g.Log().Infof(ctx, "registry initialized: etcd@%s", address)
	default:
		path := g.Cfg().MustGet(ctx, "registry.path", "/tmp/proxima/services").String()
		grpcx.Resolver.Register(file.New(path))
		g.Log().Infof(ctx, "registry initialized: file@%s", path)
	}

	// 2. 负载均衡策略 - 轮训（GoFrame 内置)
	gsel.SetBuilder(gsel.NewBuilderRoundRobin())
}
