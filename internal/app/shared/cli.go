package shared

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// CLIBase 提供CLI命令的公共功能
type CLIBase struct{}

// ReadSourceCode 读取源代码（从文件或命令行参数）
func (b *CLIBase) ReadSourceCode(file, source string) (string, error) {
	if source != "" {
		return source, nil
	}
	if file != "" {
		f, err := os.Open(file)
		if err != nil {
			return "", fmt.Errorf("打开文件失败: %w", err)
		}
		defer f.Close()

		var buffer bytes.Buffer
		_, err = io.Copy(&buffer, f)
		if err != nil {
			return "", fmt.Errorf("读取文件失败: %w", err)
		}
		return buffer.String(), nil
	}
	return "", fmt.Errorf("请指定文件(--file)或源代码(--source)")
}

// ReadFile 读取文件内容
func (b *CLIBase) ReadFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("打开文件失败: %w", err)
	}
	defer f.Close()
	return io.ReadAll(f)
}

// NewBaseCommand 创建基础命令模板
func NewBaseCommand(use, short, long string, run func(cmd *cobra.Command, args []string)) *cobra.Command {
	return &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
		Run:   run,
	}
}

// ContextWithTimeout 创建带超时的context
func ContextWithTimeout(ctx context.Context, timeoutSeconds int) (context.Context, context.CancelFunc) {
	timeout := time.Duration(timeoutSeconds) * time.Second
	return context.WithTimeout(ctx, timeout)
}
