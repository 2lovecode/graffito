# Graffito - Go开发工具箱

[![Go Version](https://img.shields.io/badge/Go-1.24+-00ADD8?style=flat&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

集成常用开发工具的工具箱，例如Go在线编辑运行沙箱、时序图生成等。同时包含一些代码片段，leetcode题解，设计模式实现，算法实现等。

## ✨ 功能特性

### 🛠️ 开发工具 (tools)
- **沙箱环境** - Go代码在线编辑运行沙箱
- **时序图生成** - PlantUML时序图自动生成工具
- **POI搜索** - 百度地图POI搜索工具
- **图片处理** - 图片合成、水印等处理工具
- **文件类型识别** - 通过文件头识别文件类型
- **地理坐标转换** - 坐标系转换工具
- **辅助工具** - 模块依赖图、Redis工具、SQL转Go、Excel处理等

### 📚 学习资源 (learn)
- **LeetCode题解** - 包含Top100和动态规划等经典题目
- **算法实现** - 数据结构与算法实现（图、树、链表、排序等）
- **设计模式** - 常用设计模式实现（单例、工厂、观察者、策略等）
- **Go实践** - Go语言特性实践代码

### 🔬 实验功能 (labs)
- **实验代码** - 各种实验性功能实现
- **CLI实践** - 命令行工具实践探索

## 🚀 快速开始

### 环境要求

| 开发环境 | 版本 | 说明 |
|:---:|:---:|:---:|
| Go | 1.24+ | 必须 |
| Java | 8+ | PlantUML功能需要 |

### 安装

```bash
# 克隆项目
git clone https://github.com/2lovecode/graffito.git
cd graffito

# 安装依赖（使用vendor模式）
go mod vendor

# 编译
go build -mod=vendor -o graffito main.go

# 或使用Makefile
make build
```

### 使用

```bash
# 查看所有命令
./graffito cli --help

# 使用工具
./graffito cli tools sandbox --source "package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"Hello\")\n}"

# 查看学习资源
./graffito cli learn --help

# 查看实验功能
./graffito cli labs --help
```

## 📖 使用指南

### 工具类命令 (tools)

#### 沙箱 (Sandbox)
在线运行Go代码

```bash
# 从文件运行
./graffito cli tools sandbox --file main.go

# 直接运行源代码
./graffito cli tools sandbox --source "package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"Hello\")\n}"
```

#### PlantUML
生成时序图和UML图

```bash
./graffito cli tools plantuml
# 监听 internal/app/plantuml/data/ 目录，自动生成图片到 images/ 目录
```

#### POI搜索
搜索百度地图POI信息

```bash
./graffito cli tools poi <ak> <经度> <纬度>
```

#### 辅助工具 (utils)

```bash
# 模块依赖图
./graffito cli tools utils modg

# Redis工具
./graffito cli tools utils redis [host] [port]

# SQL转Go
./graffito cli tools utils sql2go

# 其他工具
./graffito cli tools utils --help
```

### 学习资源命令 (learn)

```bash
# LeetCode题解
./graffito cli learn leetcode list

# 算法实现
./graffito cli learn algorithm tree

# 设计模式
./graffito cli learn pattern factory

# 查看所有学习资源
./graffito cli learn --help
```

### 实验功能命令 (labs)

```bash
# 实验代码
./graffito cli labs experiment depends

# CLI实践
./graffito cli labs practice xsync

# 查看所有实验功能
./graffito cli labs --help
```

## 📁 项目结构

```
graffito/
├── cmd/                    # 应用程序入口点
│   ├── cli/               # CLI应用程序主入口
│   ├── demo/              # 演示代码
│   └── tools/             # 工具程序（统一入口在internal/app/tools/utils）
│
├── internal/              # 内部包（不对外暴露）
│   ├── app/               # 应用层代码
│   │   ├── tools/         # 工具类功能（统一入口）
│   │   ├── learning/      # 学习资源类（统一入口）
│   │   ├── labs/          # 实验功能类（统一入口）
│   │   ├── shared/        # 共享工具（CLI基础工具）
│   │   └── [各功能模块]    # 具体功能实现
│   ├── services/          # 业务服务层
│   └── dto/               # 数据传输对象
│
├── pkg/                   # 可被外部导入的库代码
│   ├── algorithm/         # 算法实现
│   ├── pattern/           # 设计模式
│   ├── config/            # 配置管理
│   ├── errors/            # 错误处理
│   ├── logging/           # 日志系统
│   ├── fs/                # 文件系统工具
│   └── ...
│
├── docs/                  # 文档（详细说明）
└── third_party/           # 第三方资源
```

### 架构设计

项目采用清晰的分层架构，将功能按领域分为三大类：

1. **Tools（工具）** - 所有实用的开发工具
2. **Learning（学习）** - 所有学习相关的资源
3. **Labs（实验）** - 所有实验性和探索性功能

每个领域都有统一的入口点，便于管理和扩展。

## 🧪 测试

```bash
# 运行所有测试
make test

# 运行测试并查看覆盖率
make test-cover

# 运行特定包的测试
go test -mod=vendor ./pkg/algorithm/...
```

## 📝 开发指南

### 开发环境设置

```bash
# 1. 确保Go版本 >= 1.24
go version

# 2. 克隆项目
git clone https://github.com/2lovecode/graffito.git
cd graffito

# 3. 安装依赖
make vendor

# 4. 构建项目
make build
```

### 添加新命令

1. **确定命令所属领域**
   - 开发工具 → `internal/app/tools/` 下创建模块
   - 学习资源 → `internal/app/learning/` 或在对应模块添加
   - 实验功能 → `internal/app/labs/` 下创建

2. **实现命令**
   ```go
   // internal/app/tools/myfeature/main.go
   package myfeature
   
   import (
       "github.com/2lovecode/graffito/internal/app/shared"
       "github.com/2lovecode/graffito/pkg/logging"
       "github.com/spf13/cobra"
   )
   
   func NewCommand() *cobra.Command {
       base := &shared.CLIBase{}
       return &cobra.Command{
           Use:   "myfeature",
           Short: "功能描述",
           Long: `详细描述和使用示例`,
           Run: func(cmd *cobra.Command, args []string) {
               // 使用共享工具
               sourceCode, err := base.ReadSourceCode(file, source)
               if err != nil {
                   logging.Errorf("读取失败: %v", err)
                   fmt.Fprintf(cmd.ErrOrStderr(), "错误: %v\n", err)
                   return
               }
               // 实现逻辑...
           },
       }
   }
   ```

3. **注册命令**
   - 在对应的 `command.go` 中注册（`tools/command.go`, `learning/command.go` 或 `labs/command.go`）

### 代码规范

- **命名规范**：包名小写，类型名PascalCase，函数名PascalCase（导出）或camelCase（内部）
- **格式化**：使用 `make fmt` 格式化代码
- **代码检查**：使用 `make lint` 检查代码（需要安装 golangci-lint）
- **测试**：编写单元测试，使用 `make test` 运行

### 最佳实践

**使用共享工具：**
```go
base := &shared.CLIBase{}
sourceCode, err := base.ReadSourceCode(file, source)
```

**统一日志记录：**
```go
logging.Infof("操作成功: %s", result)
logging.Debugf("调试信息: %v", data)
logging.Errorf("错误: %v", err)
```

**统一错误处理：**
```go
if err != nil {
    logging.Errorf("操作失败: %v", err)
    fmt.Fprintf(cmd.ErrOrStderr(), "错误: %v\n", err)
    return
}
```

**配置管理：**
```bash
# 环境变量
export LOG_LEVEL=debug
export SANDBOX_TIMEOUT=60
```

```go
// 代码中使用
import "github.com/2lovecode/graffito/pkg/config"
cfg := config.Get()
timeout := cfg.Sandbox.Timeout
```

### 提交规范

建议遵循 [Conventional Commits](https://www.conventionalcommits.org/)：
- `feat:` 新功能
- `fix:` 修复bug
- `docs:` 文档更新
- `refactor:` 代码重构
- `test:` 测试相关
- `chore:` 构建/工具相关

## 🛠️ 常用命令 (Makefile)

```bash
make build      # 编译项目
make test       # 运行测试
make test-cover # 运行测试并生成覆盖率报告
make clean      # 清理编译文件
make vendor     # 更新vendor依赖
make fmt        # 格式化代码
make lint       # 代码检查（需要安装golangci-lint）
make install    # 安装到系统
```

## 📋 更新日志

### 最新版本

**新增功能：**
- ✨ 完善README文档，添加详细的项目介绍和使用说明
- ✨ 添加统一的错误处理包 `pkg/errors`
- ✨ 添加配置管理包 `pkg/config`，支持环境变量配置
- ✨ 改进日志系统 `pkg/logging`，支持日志级别配置
- ✨ 添加Makefile，简化常用操作
- ✨ 优化CLI命令组织，按三大领域分类（tools/learn/labs）
- ✨ 使用共享工具减少代码重复

**改进：**
- 🔧 重组项目结构，将功能模块统一到 `internal/app/`
- 🔧 将工具程序移动到 `cmd/tools/`
- 🔧 统一错误处理和日志记录
- 🔧 完善命令帮助信息和使用示例

**修复：**
- 🐛 修复未使用的导入问题
- 🐛 修复日志系统初始化问题
- 🐛 修复Context超时函数类型错误

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

1. Fork 本仓库
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启 Pull Request

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情

## 🔗 相关链接

- [Go官方文档](https://golang.org/doc/)
- [Cobra文档](https://github.com/spf13/cobra)

## 👨‍💻 作者

- **2lovecode** - [GitHub](https://github.com/2lovecode)

---

⭐ 如果这个项目对你有帮助，请给个 Star！
