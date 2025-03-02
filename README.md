# GoWebX

GoWebX 是一个基于 Go 语言开发的 Web 服务框架，采用现代化的技术栈和最佳实践，提供了完整的 Web 服务开发基础设施。

## 技术栈

- **Web 框架**: [Gin](https://github.com/gin-gonic/gin) - 高性能 HTTP Web 框架
- **配置管理**: [Viper](https://github.com/spf13/viper) - 完整的配置解决方案
- **日志系统**: [Zap](https://github.com/uber-go/zap) - 高性能日志库
- **数据库**: 
  - MySQL (使用 [sqlx](https://github.com/jmoiron/sqlx) 增强型数据库操作)
  - Redis (使用 [go-redis](https://github.com/redis/go-redis))
- **日志切割**: [lumberjack](https://github.com/natefinch/lumberjack) - 日志文件切割和管理

## 项目结构

```
GoWebX/
├── config.yaml        # 配置文件
├── main.go           # 主程序入口
├── dao/              # 数据访问层
│   ├── mysql/        # MySQL 相关操作
│   └── redis/        # Redis 相关操作
├── logger/           # 日志模块
└── routes/           # 路由配置
```

## 核心功能

### 1. 配置管理
- 使用 Viper 进行配置管理
- 支持 YAML 格式的配置文件
- 包含应用、日志、MySQL、Redis 等配置项

### 2. 日志系统
- 基于 Uber 的 Zap 日志库
- 支持日志级别配置
- 支持日志文件切割
- 包含访问日志和错误恢复日志中间件

### 3. 数据库支持
- MySQL 连接池管理
- Redis 连接池管理
- 优雅的连接关闭

### 4. Web 服务
- 基于 Gin 的路由管理
- 中间件支持
- 优雅关机支持

## 启动流程

1. **配置初始化**
   - 加载 config.yaml 配置文件
   - 解析配置到对应结构体

2. **日志初始化**
   - 配置日志级别
   - 设置日志输出格式
   - 配置日志切割

3. **数据库初始化**
   - 初始化 MySQL 连接池
   - 初始化 Redis 连接池

4. **路由注册**
   - 注册中间件
   - 配置路由规则

5. **服务启动**
   - 启动 HTTP 服务
   - 监听系统信号
   - 实现优雅关机

## 配置说明

```yaml
app:
  name: "web_app"     # 应用名称
  mode: "dev"         # 运行模式
  port: 8080          # 服务端口

log:
  level: "debug"      # 日志级别
  filename: "web_app.log"  # 日志文件名
  max_size: 30        # 单个日志文件最大尺寸(MB)
  max_age: 30         # 日志保留天数
  max_backups: 7      # 最大备份数量

mysql:
  host: "127.0.0.1"   # MySQL主机地址
  port: 3306          # MySQL端口
  user: "root"        # 用户名
  password: "12345678" # 密码
  dbname: "sql_test"   # 数据库名
  max_open_conns: 200  # 最大打开连接数
  max_idle_conns: 50   # 最大空闲连接数

redis:
  host: "127.0.0.1"    # Redis主机地址
  port: 6379           # Redis端口
  password: ""         # Redis密码
  db: 0                # 数据库索引
  pool_size: 100       # 连接池大小
```

## 优雅关机机制

项目实现了优雅关机机制，确保在服务停止时能够正确处理：
1. 等待现有请求处理完成
2. 关闭数据库连接
3. 关闭 Redis 连接
4. 保存日志

## 开发指南

1. 克隆项目
2. 配置 config.yaml
3. 确保 MySQL 和 Redis 服务可用
4. 运行服务：
   ```bash
   go run main.go
   ```

## 中间件

1. **Logger 中间件**
   - 记录请求详情
   - 包含请求路径、方法、IP、耗时等信息

2. **Recovery 中间件**
   - 捕获 panic
   - 提供错误恢复机制
   - 支持错误堆栈跟踪

## 注意事项

1. 确保配置文件中的数据库连接信息正确
2. 日志文件路径需要有写入权限
3. 建议在生产环境调整日志级别
4. 根据实际需求调整连接池配置

## 性能优化

1. 使用连接池管理数据库连接
2. 采用高性能日志库
3. 实现了优雅关机机制
4. 支持日志异步写入

## 贡献指南

欢迎提交 Issue 和 Pull Request 来帮助改进项目。

## 许可证

[MIT License](LICENSE)
