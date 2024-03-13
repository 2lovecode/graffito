# Go 编译器和编译选项
GO := go
GOFLAGS := -ldflags="-s -w"

# 项目名称
WEB_NAME := gf-web
CLI_NAME := gf-cli



# 源文件目录
SRCDIR := cmd
# 编译输出目录
BINDIR := bin

# 源文件列表
WEB_SOURCES := $(wildcard $(SRCDIR)/web/*.go)
CLI_SOURCES := $(wildcard $(SRCDIR)/cli/*.go)
UML_SOURCES := $(wildcard $(SRCDIR)/plantuml/*.go)

# 文件路径
FRONTEND_DIR := web
FRONTEND_ENTER_FILE := $(FRONTEND_DIR)/prod/index.html

# 目标文件（可执行文件）
WEB_TARGET := $(BINDIR)/$(WEB_NAME)
CLI_TARGET := $(BINDIR)/$(CLI_NAME)

# 默认目标，构建所有可执行文件
all: check_build $(WEB_TARGET) $(CLI_TARGET)

web: $(WEB_TARGET)

cli: $(CLI_TARGET)

# web
$(WEB_TARGET): $(WEB_SOURCES)
	$(GO) build $(GOFLAGS) -o $@ $(WEB_SOURCES)

# cli
$(CLI_TARGET): $(CLI_SOURCES)
	$(GO) build $(GOFLAGS) -o $@ $(CLI_SOURCES)

# 检查文件是否存在，如果不存在则先进入 web 目录运行 npm build
check_build:
	@if [ ! -f $(FRONTEND_ENTER_FILE) ]; then \
		echo "重新构建前端项目..."; \
		cd web; \
		if which cnpm > /dev/null; then \
			echo "使用 cnpm 命令构建..."; \
			cnpm install; \
			cnpm run build; \
		else \
			echo "使用 npm 命令构建..."; \
			npm install; \
			npm run build; \
		fi \
	fi


# dev
run-web: check_build
	$(GO) run -mod=vendor $(WEB_SOURCES)

run-cli:
	$(GO) run -mod=vendor $(CLI_SOURCES)

run-uml:
	$(GO) run -mod=vendor $(UML_SOURCES)



# 清理生成的文件
clean:
	rm -f $(WEB_TARGET) $(CLI_TARGET)

