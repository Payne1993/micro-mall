package account

import (
	"context"
	"proxima/app/user/internal/model/entity"

	"github.com/gogf/gf/v2/os/gtime"
)

func Register(ctx context.Context) (id int, err error) {
	return 1, nil
}

func Login(ctx context.Context) (token string, err error) {
	return "I am token", nil
}

func Info(ctx context.Context, token string) (user *entity.Users, err error) {
	return &entity.Users{
		Id:        1,
		Username:  "oldme",
		Password:  "123456",
		Email:     "cainiao@113.com",
		CreatedAt: gtime.New("2026-03-17 13:00:00"),
		UpdatedAt: gtime.New("2026-03-17 13:00:00"),
	}, nil
}
