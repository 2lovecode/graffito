# 简介

集成常用开发工具的工具箱，例如Go在线编辑运行沙箱、时序图生成等。同时包含一些代码片段，leetcode题解，设计模式实现，算法实现等。

# 运行
### 环境配置
| 开发环境 |   版本    |    说明     |
|:----:|:-------:|:---------:|
|  Go  | 1.22.3  |           |
| Node | 20.15.1 | 其他版本可自行尝试 |


### 开发环境运行
-  `make run-web` 启动web服务
-  `go run -mod=vendor cmd/cli/main.go` 执行cli命令(vendor模式)

### 构建二进制文件
- `make all` 同时构建web和cli
- `make web` 构建web应用
- `make cli` 构建cli命令行工具



# 说明
### 目录结构
- `assets` 静态资源
- `cmd` 应用入口
  - `cli` 命令行工具
  - `web` web应用
- `docs` 文档
- `third_party` 第三方依赖
- `web` web前端代码

### WEB应用简介
##### 1.Go运行沙盒
![沙盒](./assets/images/sandbox_web.png)

### CLI命令行工具简介
##### 1.Go运行沙盒
![沙盒](./assets/images/sandbox_cli.png)