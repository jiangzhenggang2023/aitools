# gotcha
Gotcha - A lightweight Go template to catch your project needs fast.

```shell
my-gin-app/
├── cmd/                  # 应用入口文件 [1,2](@ref)
│   └── server/
│       └── main.go       # 主启动文件（初始化 Gin 和路由）
├── config/               # 配置文件管理 [1,3](@ref)
│   ├── config.go         # 配置加载逻辑
│   └── config.yaml       # YAML/TOML 配置文件
├── internal/             # 私有业务逻辑（禁止外部引用）[2,4](@ref)
│   ├── api/              # 路由层
│   │   ├── v1/           # API 版本管理
│   │   │   ├── user.go   # 用户路由
│   │   │   └── product.go
│   │   └── routes.go     # 路由注册入口
│   ├── controller/       # 控制器（处理请求逻辑）[3,5](@ref)
│   ├── service/          # 业务服务层（核心逻辑）[1,5](@ref)
│   ├── repository/       # 数据持久层（数据库操作）[1,3](@ref)
│   ├── models/           # 数据模型定义 [3,5](@ref)
│   └── middleware/       # 中间件（认证/日志等）[4,7](@ref)
│       ├── auth.go
│       └── logger.go
├── pkg/                  # 公共可复用包 [2,4](@ref)
│   └── utils/            # 工具函数（加密/验证等）
├── web/                  # 前端资源 [2,6](@ref)
│   ├── static/           # CSS/JS/图片
│   └── templates/        # HTML 模板（如有）
├── tests/                # 单元测试/集成测试 [4](@ref)
├── scripts/              # 部署/运维脚本
├── go.mod                # 依赖管理
└── Makefile              # 构建自动化命令 [8](@ref)
```
