# Gin Example

基于 Gin 框架的示例后端 API 服务。

## 技术栈

- Go 1.25+
- Gin v1.11 (Web 框架)
- GORM v1.31 (ORM)
- MySQL 8.0+ (数据库)

## 项目结构

```
gin-example/
├── app/
│   └── api/v1/          # API 接口层 (Controller)
├── conf/                 # 配置文件
│   ├── dev.ini          # 开发环境配置
│   └── prod.ini         # 生产环境配置
├── dao/                  # 数据访问层 (DAO)
├── docs/
│   └── sql/db.sql       # 数据库表结构
├── library/              # 公共库
│   ├── app/             # 响应封装
│   ├── db/              # 数据库连接
│   ├── e/               # 错误码定义
│   ├── logging/         # 日志
│   └── setting/         # 配置读取
├── middleware/           # 中间件
├── model/                # 数据模型
├── routers/              # 路由定义
├── service/              # 业务逻辑层 (Service)
└── main.go              # 入口文件
```

## 功能模块

### 用户模块
- 用户登录认证
- 用户信息管理

## 数据库表

| 表名 | 说明 |
|------|------|
| users | 用户表 |
| user_tokens | 用户令牌表 |

## 快速开始

### 1. 配置数据库

```bash
# 创建数据库并导入表结构
mysql -u root -p < docs/sql/db.sql
```

### 2. 修改配置

编辑 `conf/dev.ini`:

```ini
[mysql]
Host = 127.0.0.1
Port = 3306
Username = root
Password = your_password
Database = gin_example
```

### 3. 运行

```bash
# 开发环境
go run main.go -env dev

# 生产环境
go run main.go -env prod
```

### 4. 构建

```bash
go build -o gin-example main.go
```

## API 接口

### 健康检查
- `GET /api/v1/check`

### 认证接口
- `POST /api/auth/login` - 登录
- `POST /api/auth/check` - 校验登录态
- `POST /api/auth/logout` - 退出登录

### 用户接口（需认证）
- `GET /api/user/info` - 获取用户信息

## 开发说明

### 架构模式

采用标准 MVC 分层架构:

```
Router -> API (Controller) -> Service -> DAO -> Model
```

### 新增接口步骤

1. `model/` 下定义数据模型
2. `dao/` 下实现数据访问
3. `service/` 下实现业务逻辑
4. `app/api/v1/` 下实现接口处理
5. `routers/router.go` 注册路由
