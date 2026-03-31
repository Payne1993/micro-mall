package account

import (
	"context"
	v1 "micro-mall.dev/app/user/api/account/v1"
	"micro-mall.dev/app/user/api/pbentity"
	"micro-mall.dev/app/user/internal/controller/registry"
	"micro-mall.dev/app/user/internal/logic/account"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func init() {
	registry.Add(Register)
}

type Controller struct {
	v1.UnimplementedAccountServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterAccountServer(s.Server, &Controller{})
}

func (*Controller) UserRegister(ctx context.Context, req *v1.UserRegisterReq) (res *v1.UserRegisterResp, err error) {
	id, err := account.Register(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.UserRegisterResp{
		Id: int32(id),
	}, nil
}

func (*Controller) UserLogin(ctx context.Context, req *v1.UserLoginReq) (res *v1.UserLoginRes, err error) {
	token, err := account.Login(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.UserLoginRes{
		Token: token,
	}, nil
}

func (*Controller) UserInfo(ctx context.Context, req *v1.UserInfoReq) (res *v1.UserInfoRes, err error) {
	data, err := account.Info(ctx, req.Token)
	if err != nil {
		return nil, err
	}
	return &v1.UserInfoRes{
		User: &pbentity.Users{
			Id:        uint32(data.Id),
			Username:  data.Username,
			Password:  data.Password,
			Email:     data.Email,
			CreatedAt: timestamppb.New(data.CreatedAt.Time),
			UpdatedAt: timestamppb.New(data.UpdatedAt.Time),
		},
	}, nil
}
