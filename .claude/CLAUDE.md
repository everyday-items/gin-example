# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## 项目文档

- **docs/sql/db.sql** - 数据库表结构定义
- **docs/apifox.md** - API 接口文档

## 常用命令

```bash
# 开发环境运行
go run main.go -env dev

# 生产环境运行
go run main.go -env prod

# 构建
go build -o gin-example main.go

# 编译检查
go build -o /dev/null ./...

# Docker 构建
docker build -t gin-example .

# 初始化数据库
mysql -u root -p < docs/sql/db.sql
```

## 架构

基于 Gin + GORM + MySQL 的后端 API 服务。

### 分层架构

```
Router -> API (Controller) -> Service -> DAO -> Model
```

- **app/api/v1/** - 接口层，处理 HTTP 请求/响应，调用 Service
- **service/** - 业务逻辑层，面向对象风格 (struct + 方法)
- **dao/** - 数据访问层，面向对象风格 (struct + 方法)
- **model/** - 数据模型，每个表一个文件

### 依赖注入模式

```go
// DAO 层
type UserProfileDao struct{}
func NewUserProfileDao() *UserProfileDao

// Service 层注入 DAO
type UserProfileService struct {
    dao *dao.UserProfileDao
}
func NewUserProfileService() *UserProfileService

// API 层注入 Service
type UserProfileApi struct {
    svc *service.UserProfileService
}
var userProfileApi = NewUserProfileApi()
```

### Model 定义规范

每个 Model 需要实现 `TableName()` 方法指定表名:

```go
type UserProfile struct {
    ID     uint64 `json:"id" gorm:"primaryKey;autoIncrement"`
    UserID uint64 `json:"userId" gorm:"column:user_id;uniqueIndex"`
    // ...
}

func (UserProfile) TableName() string {
    return "user_profiles"
}
```

### 响应格式

使用 `library/app/` 封装统一响应:

```go
appG := app.Gin{C: c}
appG.Response(http.StatusOK, e.SUCCESS, data)        // 标准响应
appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)  // 错误响应
```

错误码定义在 `library/e/code.go`，消息映射在 `library/e/msg.go`。

### JWT 认证

需要认证的路由使用 `jwt.Auth()` 中间件，从 context 获取用户信息:

```go
userID := c.MustGet("userID").(uint64)
openid := c.MustGet("openid").(string)
```

### 核心库

- **library/db/** - 数据库单例连接 (`db.GetDB()`)
- **library/setting/** - 配置读取 (`setting.Setup(env)`)
- **library/app/** - 响应封装 (`app.Gin{}.Response()`)
- **library/e/** - 错误码定义
- **library/logging/** - 日志 (`logging.Info()`, `logging.Error()`, `logging.Debug()`)
- **library/util/** - 工具函数 (JSON 转换, 类型转换等)
- **middleware/jwt/** - Token 认证中间件

## 配置

配置文件位于 `conf/` 目录，通过 `-env` 参数选择:
- `conf/dev.ini` - 开发环境
- `conf/test.ini` - 测试环境
- `conf/pre.ini` - 预发布环境
- `conf/prod.ini` - 生产环境

主要配置 section: `[server]`, `[mysql]`, `[wechat]`, `[app]`

## 新增接口流程

1. `model/xxx.go` - 定义表模型 (struct + TableName) 和请求/响应 DTO
2. `dao/xxx.go` - 实现 DAO struct 和数据访问方法
3. `service/xxx.go` - 实现 Service struct 和业务逻辑
4. `app/api/v1/xxx.go` - 实现 API 处理函数
5. `routers/router.go` - 注册路由 (认证接口加 `jwt.Auth()` 中间件)
