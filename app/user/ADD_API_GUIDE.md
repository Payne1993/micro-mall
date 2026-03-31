# 在当前 GoFrame 脚手架中新增接口

本文档说明如何在当前仓库结构下，为 `app/user` 模块新增一个接口。

适用范围：
- 给已有 gRPC service 新增一个 rpc
- 新增一个独立的 gRPC service
- 在需要时补充 gateway 对外访问层

## 一、先判断接口应该放哪里

在动手之前，先判断接口归属。

### 1. 放进已有 service

如果接口属于现有业务域，就继续放在已有 proto service 里。

例如：
- 用户注册
- 用户登录
- 用户信息
- 修改密码

这类接口应继续维护在：
- `app/user/manifest/protobuf/account/v1/account.proto`

### 2. 做成独立 service

如果接口是独立能力，不属于 account 业务域，就单独建一个 service。

例如：
- health
- ping
- metrics
- admin

这类接口建议参考：
- `app/user/manifest/protobuf/health/v1/health.proto`

## 二、当前脚手架里的关键目录

以 `app/user` 为例，新增接口时主要会改这些位置：

- `manifest/protobuf/`：接口协议定义
- `api/`：由 `gf gen pb` 自动生成的 pb 文件
- `internal/controller/`：gRPC 控制器
- `internal/logic/`：业务逻辑
- `internal/cmd/cmd.go`：服务启动与注册入口
- `hack/hack.mk`：生成命令入口

其中生成命令定义在：
- `app/user/hack/hack.mk`

当前生成 protobuf 的命令是：

```bash
gf gen pb
```

## 三、给已有 service 新增一个 rpc

这里以给 `Account` 增加一个新接口为例说明。

### 步骤 1：修改 proto

编辑：

- `app/user/manifest/protobuf/account/v1/account.proto`

例如新增 `UserLogout`：

```proto
service Account {
  rpc UserRegister (UserRegisterReq) returns (UserRegisterResp) {}
  rpc UserLogin (UserLoginReq) returns (UserLoginRes){}
  rpc UserInfo (UserInfoReq) returns (UserInfoRes){}
  rpc UserLogout (UserLogoutReq) returns (UserLogoutRes){}
}

message UserLogoutReq {
  string token = 1;
}

message UserLogoutRes {
  bool success = 1;
}
```

命名建议：
- rpc 名和业务动作保持一致
- 请求名使用 `XxxReq`
- 响应名使用 `XxxRes`
- 不要在这里手写 Go 逻辑

### 步骤 2：生成 pb 文件

在 `app/user` 目录执行：

```bash
gf gen pb
```

执行后会自动更新：
- `app/user/api/account/v1/account.pb.go`
- `app/user/api/account/v1/account_grpc.pb.go`

注意：
- 不要手改这两个生成文件
- proto 有变动时，必须重新执行一次 `gf gen pb`

### 步骤 3：实现 controller 方法

编辑：

- `app/user/internal/controller/account/account.go`

补上新方法，例如：

```go
func (*Controller) UserLogout(ctx context.Context, req *v1.UserLogoutReq) (res *v1.UserLogoutRes, err error) {
	success, err := account.Logout(ctx, req.Token)
	if err != nil {
		return nil, err
	}
	return &v1.UserLogoutRes{Success: success}, nil
}
```

controller 职责很简单：
- 接收 gRPC 请求
- 调用 logic
- 组装响应

### 步骤 4：实现 logic

编辑或新增：

- `app/user/internal/logic/account/*.go`

例如：

```go
func Logout(ctx context.Context, token string) (success bool, err error) {
	return true, nil
}
```

logic 负责：
- 业务处理
- 调用 DAO
- 参数转换
- 错误返回

### 步骤 5：编译验证

在 `app/user` 目录执行：

```bash
go build ./...
```

如果通过，说明已有 service 上新增 rpc 已完成。

### 这一类改动的特点

这种场景下，一般不需要改：
- `app/user/internal/cmd/cmd.go`

原因是 `Account` service 已经注册过了，你只是给它新增了一个方法，不是新增一个新的 service。

## 四、新增一个独立 service

如果接口不属于 `Account`，推荐单独建 service。

当前仓库里的 `health` 就是这种方式。

### 步骤 1：新增 proto 文件

例如新建：

- `app/user/manifest/protobuf/health/v1/health.proto`

示例：

```proto
syntax = "proto3";

package health.v1;

option go_package = "proxima/app/user/api/health/v1";

service Health {
  rpc Health (HealthReq) returns (HealthRes) {}
}

message HealthReq {}

message HealthRes {
  string status = 1;
}
```

命名建议：
- service 名和 rpc 名尽量统一
- 包名、目录名、controller 包名、logic 包名保持一致

### 步骤 2：生成 pb 文件

在 `app/user` 目录执行：

```bash
gf gen pb
```

会新增：
- `app/user/api/health/v1/health.pb.go`
- `app/user/api/health/v1/health_grpc.pb.go`

### 步骤 3：新增 logic

新建：

- `app/user/internal/logic/health/health.go`

示例：

```go
package health

import "context"

func Health(ctx context.Context) (status string, err error) {
	return "ok", nil
}
```

### 步骤 4：新增 controller

新建：

- `app/user/internal/controller/health/health.go`

示例：

```go
package health

import (
	"context"
	v1 "proxima/app/user/api/health/v1"
	logic "proxima/app/user/internal/logic/health"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

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
```

### 步骤 5：让 controller 被自动发现

当前 `app/user` 已改为自动注册模式，不需要再去：

- `app/user/internal/cmd/cmd.go`

里手工补 `xxx.Register(s)`。

你需要保证两件事：

1. controller 包里保留标准的 `Register(s *grpcx.GrpcServer)` 方法
2. controller 包在 `init()` 中调用注册中心登记自己

示例：

```go
func init() {
	registry.Add(Register)
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterHealthServer(s.Server, &Controller{})
}
```

然后在 `app/user/internal/controller` 目录执行：

```bash
go generate ./...
```

这会自动更新：

- `app/user/internal/controller/autoload_gen.go`

它的作用是自动导入所有 controller 包，从而触发各自的 `init()` 注册逻辑。

如果忘了执行 `go generate ./...`，常见现象是：

- 编译正常
- 服务启动正常
- 但新 controller 没有被注册

### 步骤 6：编译验证

执行：

```bash
go build ./...
```

## 五、标准推荐顺序

在当前仓库里，推荐你每次新增接口按这个顺序做：

1. 先判断接口属于已有 service 还是独立 service
2. 先改 proto
3. 执行 `gf gen pb`
4. 再写 logic
5. 再写 controller
6. 如果是独立 service，给 controller 增加 `init()` 自注册，并执行 `go generate ./...`
7. 执行 `go build ./...`

这个顺序的好处是：
- 协议层最早确定
- 命名最早收敛
- 生成问题最早暴露
- 不会写完 controller 才发现 proto 定义不合理

## 六、给 gateway 提供 HTTP 接口（grpc-gateway）

当前项目使用 **grpc-gateway** 将 HTTP/JSON 请求自动转发为 gRPC 调用。

### 场景 A：在已有 service 上新增方法（如给 Account 加 UserLogout）

只需改 2 个文件，跑 2 条命令。

#### 1. 添加 HTTP 路由映射

编辑 `app/user/manifest/protobuf/account/v1/account_gateway.yaml`，追加：

```yaml
    - selector: account.v1.Account.UserLogout
      post: /api/v1/user/logout
      body: "*"
```

#### 2. 重新生成 grpc-gateway 代码

在 `app/user/` 目录执行：

```bash
make gw
```

会更新 `api/account/v1/account.pb.gw.go`，自动包含新路由。

#### 无需额外改动

- **gRPC 服务注册**：`Controller` 内嵌了 `UnimplementedAccountServer`，新方法自动生效。
- **Gateway 路由**：`account` handler 已配置在 `app/gateway/manifest/config/config.yaml` 的 `gateway.upstreams` 中，无需修改网关代码。
- **中间件**：新路由自动继承 `/api/v1` 下的 CORS、限流等中间件。

### 场景 B：新增独立 service 的 HTTP 接口

以新增一个 Health 服务的 HTTP 端点为例。

#### 1. 创建 gateway 映射文件

新建 `app/user/manifest/protobuf/health/v1/health_gateway.yaml`：

```yaml
type: google.api.Service
config_version: 3

http:
  rules:
    - selector: health.v1.Health.Health
      get: /api/v1/user/health
```

#### 2. 在 hack.mk 中添加生成命令

编辑 `app/user/hack/hack.mk`，在 `gw` target 中追加：

```makefile
	@PATH="$$PATH:$$HOME/go/bin" protoc \
		-I manifest/protobuf \
		--grpc-gateway_out=paths=source_relative,grpc_api_configuration=manifest/protobuf/health/v1/health_gateway.yaml:api \
		manifest/protobuf/health/v1/health.proto
```

#### 3. 生成 grpc-gateway 代码

```bash
make gw
```

#### 4. 在 gateway 中注册 handler

编辑 `app/gateway/internal/router/gateway.go`，在 `handlerRegistry` 中添加：

```go
var handlerRegistry = map[string]HandlerRegisterFunc{
	"account": accountv1.RegisterAccountHandler,
	"health":  healthv1.RegisterHealthHandler,  // 新增
}
```

#### 5. 在 gateway 配置中声明

编辑 `app/gateway/manifest/config/config.yaml`，在对应 upstream 的 handlers 列表中添加：

```yaml
gateway:
  upstreams:
    - name: user
      handlers:
        - account
        - health  # 新增
```

### 场景 C：接入全新模块的 HTTP 接口（如 word 服务）

除以上步骤外，还需：

1. 在 `app/gateway/manifest/config/config.yaml` 中添加新的 upstream：

```yaml
gateway:
  upstreams:
    - name: word
      handlers:
        - words
```

2. 在 `app/gateway/internal/router/gateway.go` 的 `handlerRegistry` 中注册：

```go
"words": wordsv1.RegisterWordsHandler,
```

## 七、常见错误

### 1. 手改生成文件

不要手改：
- `*.pb.go`
- `*_grpc.pb.go`

这些文件下一次生成会被覆盖。

### 2. 改了 proto 但没执行 `gf gen pb`

会导致：
- 新的消息类型不存在
- 新的方法签名不存在
- controller 编译失败

### 3. 新增独立 service 后忘了生成 controller 自动导入

如果你新增的是独立 service，除了写 `Register` 之外，还要：

- 在 controller 包里 `init()` 调用 `registry.Add(Register)`
- 执行 `go generate ./...`

否则 controller 包不会被自动导入，注册逻辑也不会执行。

### 4. 命名不统一

不建议出现这种情况：
- proto 里叫 `Health`
- logic 里叫 `Check`
- controller 里叫 `Ping`

建议统一：
- service 名
- rpc 名
- controller 方法名
- logic 方法名

### 5. 不相关能力强行塞进已有业务 service

例如把健康检查塞进 `Account`，短期省事，长期会导致 service 边界混乱。

## 八、最小检查清单

每次新增接口后，至少检查以下项目：

- proto 是否已修改
- `gf gen pb` 是否已执行
- controller 是否已实现
- logic 是否已实现
- 独立 service 是否已在 controller 包中完成自注册
- `go generate ./...` 是否已执行
- `go build ./...` 是否通过

## 九、常用命令

在 `app/user` 目录下：

```bash
gf gen pb
go generate ./internal/controller/...
gofmt -w internal/controller/... internal/logic/... internal/cmd/cmd.go
go build ./...
```

如果是新增数据库表相关内容，还会用到：

```bash
gf gen dao
gf gen pbentity
```

## 十、一句话总结

当前 GoFrame 脚手架下，新增接口的标准流程是：

先改 proto，再生成 pb，然后写 logic 和 controller；如果是独立 service，再到 `cmd.go` 注册，最后执行 `go build ./...` 验证。