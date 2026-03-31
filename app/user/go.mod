module micro-mall.dev/app/user

go 1.25.4

require (
	github.com/gogf/gf/contrib/drivers/mysql/v2 v2.10.0
	github.com/gogf/gf/contrib/rpc/grpcx/v2 v2.10.0
	github.com/gogf/gf/v2 v2.10.0
	github.com/golang-jwt/jwt/v5 v5.3.1
	google.golang.org/grpc v1.79.2
	google.golang.org/protobuf v1.36.11
	micro-mall.dev/pkg v0.0.0
)

replace micro-mall.dev/pkg => ../../pkg
