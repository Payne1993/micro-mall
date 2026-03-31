# 在 grpc-gateway 模块中新增 API 接口

本文档说明如何在 `app/grpc-gateway` 模块中新增或接入一个 HTTP API 接口。

`grpc-gateway` 模块的职责是把 HTTP/JSON 请求转发为 gRPC 调用。它本身不包含业务逻辑，所有业务实现在各个 gRPC 服务模块中（如 `app/user`、`app/word`）。

## 一、先判断场景

### 场景 A：给已有 service 新增方法

例如给 `Account` 增加 `UserLogout`。

**grpc-gateway 模块无需任何改动。**

只需要在对应的 gRPC 服务模块中完成：
1. 修改 proto，新增 rpc
2. 新增 gateway yaml 映射
3. 重新生成 pb 和 gw 文件
4. 实现 controller 和 logic

handler 已经注册过（如 `"account": accountv1.RegisterAccountHandler`），新方法会自动包含在生成的 `*.pb.gw.go` 中。

### 场景 B：在已有 gRPC 模块中新增独立 service

例如在 `app/user` 中新增一个 `Profile` service。

需要改 grpc-gateway 的 **2 个文件**。

### 场景 C：接入一个全新的 gRPC 模块

例如新增 `app/chat` 模块并将其 HTTP 接口暴露出来。

需要改 grpc-gateway 的 **2 个文件**，并确保 import 路径正确。

## 二、关键文件

grpc-gateway 模块中，新增接口时只会改动以下文件：

| 文件 | 作用 |
|------|------|
| `internal/gateway/gateway.go` | 注册 handler 到 `runtime.ServeMux` |
| `manifest/config/config.yaml` | 声明 upstream 和 handler 映射 |

## 三、场景 B：新增独立 service 的 HTTP 接口

以在 `app/user` 中新增 `Profile` service 为例。

### 前提

在 gRPC 服务模块中已完成：
1. 编写 `app/user/manifest/protobuf/profile/v1/profile.proto`
2. 创建 `app/user/manifest/protobuf/profile/v1/profile_gateway.yaml`
3. 执行 `gf gen pb` 和 `make gw` 生成代码
4. 实现 controller 和 logic
5. 在 `app/user/internal/cmd/cmd.go` 中注册 service

### 步骤 1：在 handlerRegistry 中注册

编辑 `app/grpc-gateway/internal/gateway/gateway.go`：

```go
import (
    // ...已有 import
    profilev1 "proxima/app/user/api/profile/v1"  // 新增
)

var handlerRegistry = map[string]HandlerRegisterFunc{
    "account": accountv1.RegisterAccountHandler,
    "health":  healthv1.RegisterHealthHandler,
    "words":   wordsv1.RegisterWordsHandler,
    "profile": profilev1.RegisterProfileHandler,  // 新增
}
```

### 步骤 2：在配置文件中声明

编辑 `app/grpc-gateway/manifest/config/config.yaml`，在对应 upstream 的 `handlers` 列表中添加：

```yaml
gateway:
  upstreams:
    - name: user
      handlers:
        - account
        - health
        - profile  # 新增
    - name: word
      handlers:
        - words
```

### 步骤 3：编译验证

```bash
cd app/grpc-gateway
go build ./...
```

## 四、场景 C：接入全新 gRPC 模块

以接入新模块 `app/chat` 的 `Message` service 为例。

### 前提

`app/chat` 模块已完成：
1. proto 定义、pb 生成、gw 生成
2. controller 和 logic 实现
3. gRPC 服务可正常启动并注册到服务发现

### 步骤 1：在 handlerRegistry 中注册

编辑 `app/grpc-gateway/internal/gateway/gateway.go`：

```go
import (
    // ...已有 import
    messagev1 "proxima/app/chat/api/message/v1"  // 新增
)

var handlerRegistry = map[string]HandlerRegisterFunc{
    "account": accountv1.RegisterAccountHandler,
    "health":  healthv1.RegisterHealthHandler,
    "words":   wordsv1.RegisterWordsHandler,
    "message": messagev1.RegisterMessageHandler,  // 新增
}
```

### 步骤 2：在配置文件中添加新 upstream

编辑 `app/grpc-gateway/manifest/config/config.yaml`：

```yaml
gateway:
  upstreams:
    - name: user
      handlers:
        - account
        - health
    - name: word
      handlers:
        - words
    - name: chat       # 新增 upstream
      handlers:
        - message
```

`name` 必须与 gRPC 服务在服务发现中注册的名称一致（即对应服务 `manifest/config/config.yaml` 中的 `grpc.name`）。

### 步骤 3：编译验证

```bash
cd app/grpc-gateway
go build ./...
```

## 五、工作原理简述

```
HTTP 请求
  │
  ▼
net/http Server (:8100)
  │
  ├─ /health → 直接返回 {"status":"ok"}
  │
  └─ /* → middleware 链 (CORS → RateLimiter)
         │
         ▼
       grpc-gateway runtime.ServeMux
         │
         ├─ 根据 URL 路径匹配 *.pb.gw.go 中的路由规则
         │
         ▼
       grpcx.Client.NewGrpcClientConn(serviceName)
         │
         ├─ 通过服务发现（file registry / etcd）找到目标 gRPC 服务
         │
         ▼
       gRPC 后端服务
```

配置中的 `gateway.upstreams` 控制了：
- 连接哪个 gRPC 服务（`name` → 服务发现名称）
- 注册哪些 handler（`handlers` → `handlerRegistry` 中的 key）

每个 handler 的路由规则由 gRPC 服务模块中的 `*_gateway.yaml` 定义，编译在 `*.pb.gw.go` 中。

## 六、最小检查清单

每次新增接口后，检查以下项目：

- [ ] gRPC 服务模块中 proto、pb、gw 文件已就绪
- [ ] gRPC 服务模块中 controller 和 logic 已实现
- [ ] gRPC 服务模块中 service 已在 cmd.go 注册
- [ ] `gateway.go` 中 `handlerRegistry` 已添加新 handler（场景 B/C）
- [ ] `config.yaml` 中 `gateway.upstreams` 已声明（场景 B/C）
- [ ] `go build ./...` 编译通过
- [ ] 启动 grpc-gateway 后 HTTP 接口可正常访问

## 七、常见错误

### 1. handler 名不匹配

`config.yaml` 中 `handlers` 列表里的名字必须与 `handlerRegistry` 的 key 完全一致，否则启动时 panic：

```
unknown gateway handler "xxx" for service "yyy"
```

### 2. upstream name 与服务发现名称不一致

`config.yaml` 中 `name` 必须是 gRPC 服务在服务发现中注册的名称。如果不一致，会连接失败：

```
failed to connect to service "xxx"
```

### 3. 忘了生成 gw 文件

如果 gRPC 服务模块中没有执行 `make gw`，`RegisterXxxHandler` 函数不存在，编译直接报错。

### 4. gRPC 服务未启动

grpc-gateway 启动时并不会校验后端是否在线，但请求时会返回 gRPC 错误（如 `Unavailable`）。

## 八、常用命令

```bash
# 编译
cd app/grpc-gateway && go build ./...

# 开发运行
cd app/grpc-gateway && gf run main.go

# 测试接口
curl http://127.0.0.1:8100/api/v1/user/info
```
