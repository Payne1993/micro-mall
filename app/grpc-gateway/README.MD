# grpc-gateway

Pure grpc-gateway (grpc-ecosystem/grpc-gateway/v2) based HTTP gateway for Proxima.

Unlike the GoFrame-based gateway (`app/gateway`), this module uses `net/http` directly with
`grpc-gateway`'s `runtime.ServeMux` as the primary HTTP handler.

## Features

- Native grpc-gateway v2 `runtime.ServeMux` with `protojson` marshaling
- Standard `net/http` middleware chain (CORS, rate limiter, auth)
- GoFrame config/logging for configuration management
- GoFrame `grpcx` resolver for service discovery (file registry / etcd)
- Health check endpoint at `/health`

## Run

```bash
cd app/grpc-gateway
gf run main.go
```

## Configuration

See `manifest/config/config.yaml` for server and upstream configuration.
