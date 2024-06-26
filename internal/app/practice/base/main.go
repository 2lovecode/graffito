package base

import "context"

// Question 问题
type Question interface {
	Name() Name          // 唯一标识
	Description() string // 问题描述
	Run(ctx context.Context) (string, error)
}
