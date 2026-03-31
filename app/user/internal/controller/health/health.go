package health

import (
	"context"
	v1 "micro-mall.dev/app/user/api/health/v1"
	"micro-mall.dev/app/user/internal/controller/registry"
	logic "micro-mall.dev/app/user/internal/logic/health"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

func init() {
	registry.Add(Register)
}

type Controller struct {
	v1.UnimplementedHealthServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterHealthServer(s.Server, &Controller{})
}

func (*Controller) Health(ctx context.Context, req *v1.HealthReq) (res *v1.HealthRes, err error) {
	status, err := logic.Health(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.HealthRes{Status: status}, nil
}
