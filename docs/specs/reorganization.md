# 项目重组说明

## 重组概述

本次重组按照Go标准项目结构重新组织了代码，提高了项目的可维护性和清晰度。

## 主要变更

### 1. 功能模块统一到 internal/app/
- 原 `cmd/cli/app/*` 下的功能模块已移动到 `internal/app/`
- 新的模块位置：
  - `internal/app/geo/` - 地理坐标转换
  - `internal/app/image2/` - 图片处理
  - `internal/app/media/` - 媒体文件类型识别
  - `internal/app/plantuml/` - PlantUML时序图生成
  - `internal/app/poi/` - POI搜索
  - `internal/app/cli-practice/` - CLI实践代码
  - `internal/app/depends-cli/` - 依赖管理CLI
  - `internal/app/sandbox-cli/` - 沙箱CLI命令
  - `internal/app/search-cli/` - 搜索CLI命令

### 2. 工具程序移动到 cmd/tools/
- 原 `tools/*` 下的工具程序已移动到 `cmd/tools/`
- 新的工具位置：
  - `cmd/tools/color/` - 颜色工具
  - `cmd/tools/excel/` - Excel工具
  - `cmd/tools/helper/` - 帮助工具
  - `cmd/tools/modg/` - 模块图工具
  - `cmd/tools/modgv/` - 模块可视化工具
  - `cmd/tools/onezero/` - 零一工具
  - `cmd/tools/redis/` - Redis工具
  - `cmd/tools/sql2go/` - SQL转Go工具
  - `cmd/tools/string_op/` - 字符串操作工具
  - `cmd/tools/subway/` - 地铁工具
  - `cmd/tools/command.go` - 工具命令入口

### 3. 目录结构说明

```
graffito/
├── cmd/                    # 应用程序入口点
│   ├── cli/               # CLI应用程序
│   ├── demo/              # 演示代码
│   ├── tools/             # 工具程序（原 tools/）
│   └── web/               # Web应用程序
├── internal/              # 内部包（不对外暴露）
│   └── app/               # 应用层代码
│       ├── base/          # 基础接口
│       ├── cli-practice/  # CLI实践（原 cmd/cli/app/practice）
│       ├── depends-cli/   # 依赖CLI（原 cmd/cli/app/depends）
│       ├── experiment/    # 实验性功能
│       ├── geo/           # 地理功能（原 cmd/cli/app/geo）
│       ├── image2/        # 图片处理（原 cmd/cli/app/image2）
│       ├── learn/         # 学习代码
│       ├── leetcode/      # LeetCode题解
│       ├── media/         # 媒体处理（原 cmd/cli/app/media）
│       ├── plantuml/      # PlantUML（原 cmd/cli/app/plantuml）
│       ├── poi/           # POI搜索（原 cmd/cli/app/poi）
│       ├── practice/      # 实践代码
│       ├── sandbox-cli/   # 沙箱CLI（原 cmd/cli/app/sandbox）
│       ├── search/        # 搜索功能
│       ├── search-cli/    # 搜索CLI（原 cmd/cli/app/search）
│       └── ...
├── pkg/                   # 可被外部导入的库代码
│   ├── algorithm/         # 算法实现
│   ├── cache/             # 缓存
│   ├── fs/                # 文件系统
│   ├── pattern/           # 设计模式
│   └── ...
├── tools/                 # 已移动到 cmd/tools/
├── web/                   # Web前端代码
└── docs/                  # 文档

```

## 导入路径变更

### 原路径 → 新路径

- `github.com/2lovecode/graffito/cmd/cli/app/*` → `github.com/2lovecode/graffito/internal/app/*`
- `github.com/2lovecode/graffito/tools/*` → `github.com/2lovecode/graffito/cmd/tools/*`

## 注意事项

1. 所有导入路径已更新
2. 旧的 `cmd/cli/app/` 目录可以安全删除（已备份到 `internal/app/`）
3. 旧的 `tools/` 目录可以安全删除（已移动到 `cmd/tools/`）
4. 如果发现任何导入错误，请检查导入路径是否正确

## 验证

运行以下命令验证重组是否成功：

```bash
go build ./...
```
