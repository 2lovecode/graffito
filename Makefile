.PHONY: build run test clean help vendor fmt lint web-build

# 变量
BINARY_NAME=graffito
MAIN_PATH=main.go
VENDOR_FLAG=-mod=vendor

# 默认目标
.DEFAULT_GOAL := help

## help: 显示帮助信息
help:
	@echo "可用的命令:"
	@echo "  make build      - 编译项目"
	@echo "  make run        - 运行项目 (CLI模式)"
	@echo "  make run-web    - 运行Web服务"
	@echo "  make test       - 运行测试"
	@echo "  make test-cover - 运行测试并生成覆盖率报告"
	@echo "  make clean      - 清理编译文件"
	@echo "  make vendor     - 更新vendor依赖"
	@echo "  make fmt        - 格式化代码"
	@echo "  make lint       - 代码检查"
	@echo "  make web-build  - 构建Web前端"
	@echo "  make install    - 安装到系统"

## build: 编译项目
build:
	@echo "编译项目..."
	@go build $(VENDOR_FLAG) -o $(BINARY_NAME) $(MAIN_PATH)
	@echo "编译完成: $(BINARY_NAME)"

## run: 运行CLI
run:
	@go run $(VENDOR_FLAG) $(MAIN_PATH) cli

## run-web: 运行Web服务
run-web:
	@go run $(VENDOR_FLAG) $(MAIN_PATH) web

## test: 运行测试
test:
	@echo "运行测试..."
	@go test $(VENDOR_FLAG) ./...

## test-cover: 运行测试并生成覆盖率报告
test-cover:
	@echo "运行测试并生成覆盖率报告..."
	@go test $(VENDOR_FLAG) -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "覆盖率报告已生成: coverage.html"

## clean: 清理编译文件
clean:
	@echo "清理编译文件..."
	@rm -f $(BINARY_NAME)
	@rm -f coverage.out coverage.html
	@echo "清理完成"

## vendor: 更新vendor依赖
vendor:
	@echo "更新vendor依赖..."
	@go mod vendor
	@echo "vendor依赖更新完成"

## fmt: 格式化代码
fmt:
	@echo "格式化代码..."
	@go fmt ./...
	@echo "格式化完成"

## lint: 代码检查（需要安装golangci-lint）
lint:
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run; \
	else \
		echo "请先安装 golangci-lint: https://golangci-lint.run/"; \
	fi

## web-build: 构建Web前端
web-build:
	@echo "构建Web前端..."
	@cd web && npm install && npm run build
	@echo "Web前端构建完成"

## install: 安装到系统
install:
	@echo "安装到系统..."
	@go install $(VENDOR_FLAG) $(MAIN_PATH)
	@echo "安装完成"
