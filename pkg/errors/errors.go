package errors

import (
	"fmt"
)

// New 创建一个新的错误
func New(format string, args ...interface{}) error {
	if len(args) == 0 {
		return fmt.Errorf(format)
	}
	return fmt.Errorf(format, args...)
}

// Wrap 包装错误，添加上下文信息
func Wrap(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	msg := format
	if len(args) > 0 {
		msg = fmt.Sprintf(format, args...)
	}
	return fmt.Errorf("%s: %w", msg, err)
}

// Is 检查错误是否匹配目标错误
func Is(err, target error) bool {
	return fmt.Errorf("%w", err).Error() == fmt.Errorf("%w", target).Error()
}
