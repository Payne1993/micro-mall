package handlers

import (
	accountv1 "micro-mall.dev/app/user/api/account/v1"
	healthv1 "micro-mall.dev/app/user/api/health/v1"

	"micro-mall.dev/app/grpc-gateway/internal/registry"
)

func init() {
	registry.Add("user", accountv1.RegisterAccountHandler)
	registry.Add("user", healthv1.RegisterHealthHandler)
}