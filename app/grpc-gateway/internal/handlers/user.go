package handlers

import (
	accountv1 "proxima/app/user/api/account/v1"
	healthv1 "proxima/app/user/api/health/v1"

	"proxima/app/grpc-gateway/internal/registry"
)

func init() {
	registry.Add("user", accountv1.RegisterAccountHandler)
	registry.Add("user", healthv1.RegisterHealthHandler)
}