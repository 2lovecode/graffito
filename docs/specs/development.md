# 开发指南

## 项目概述

Graffito 是一个Go开发工具箱项目，集成了多种开发工具和学习资源。

## 开发环境设置

### 1. 安装Go

确保安装了Go 1.24或更高版本：

```bash
go version
```

### 2. 克隆项目

```bash
git clone https://github.com/2lovecode/graffito.git
cd graffito
```

### 3. 安装依赖

```bash
# 使用vendor模式
make vendor

# 或直接使用
go mod vendor
```

### 4. 构建项目

```bash
make build
```

## 代码组织

### 目录结构说明

- `cmd/` - 应用程序入口点
  - `cli/` - CLI应用主入口
  - `web/` - Web应用入口
  - `tools/` - 工具程序
  
- `internal/` - 内部包，不对外暴露
  - `app/` - 应用层功能模块
  - `controllers/` - Web控制器
  - `services/` - 业务服务层
  - `dto/` - 数据传输对象
  
- `pkg/` - 可被外部项目导入的库
  - `algorithm/` - 算法实现
  - `pattern/` - 设计模式
  - `errors/` - 错误处理
  - `config/` - 配置管理

### 命名规范

- 包名：小写，简短
- 文件名：snake_case 或 kebab-case
- 类型名：PascalCase
- 函数名：PascalCase（导出）或 camelCase（内部）
- 常量：PascalCase

## 添加新功能

### 1. 添加新CLI命令

在 `internal/app/` 下创建新模块：

```go
package myfeature

import "github.com/spf13/cobra"

func NewCommand() *cobra.Command {
    return &cobra.Command{
        Use:   "myfeature",
        Short: "功能描述",
        Long:  "详细描述",
        Run: func(cmd *cobra.Command, args []string) {
            // 实现逻辑
        },
    }
}
```

在 `cmd/cli/main.go` 中注册：

```go
import "github.com/2lovecode/graffito/internal/app/myfeature"

// 在 NewCommand() 中添加
cliCmd.AddCommand(myfeature.NewCommand())
```

### 2. 添加新Web接口

在 `internal/controllers/` 下创建控制器：

```go
package controllers

func (c *MyController) Handler(ctx iris.Context) {
    // 处理逻辑
    response.Success(ctx, data)
}
```

在 `cmd/web/main.go` 中注册路由。

### 3. 添加新的工具包

在 `pkg/` 下创建新的包：

```go
package mypkg

// 导出函数
func PublicFunction() {
    // 实现
}
```

## 测试

### 编写测试

```go
package mypkg_test

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestMyFunction(t *testing.T) {
    result := MyFunction()
    assert.Equal(t, expected, result)
}
```

### 运行测试

```bash
# 运行所有测试
make test

# 运行特定包的测试
go test ./pkg/mypkg/...

# 运行测试并查看覆盖率
make test-cover
```

## 代码规范

### 格式化

```bash
make fmt
```

### 代码检查

建议安装 `golangci-lint`：

```bash
# macOS
brew install golangci-lint

# 运行检查
make lint
```

### 提交规范

建议遵循 [Conventional Commits](https://www.conventionalcommits.org/)：

- `feat:` 新功能
- `fix:` 修复bug
- `docs:` 文档更新
- `style:` 代码格式调整
- `refactor:` 代码重构
- `test:` 测试相关
- `chore:` 构建/工具相关

## 常见问题

### Q: 如何添加新的依赖？

A: 
```bash
go get package/name
go mod vendor
```

### Q: 如何使用环境变量配置？

A: 使用 `pkg/config` 包：

```go
import "github.com/2lovecode/graffito/pkg/config"

cfg := config.Get()
port := cfg.Web.Port
```

### Q: 如何统一处理错误？

A: 使用 `pkg/errors` 包：

```go
import "github.com/2lovecode/graffito/pkg/errors"

if err != nil {
    return errors.Wrap(err, "操作失败")
}
```

## 贡献流程

1. Fork 项目
2. 创建功能分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启 Pull Request
