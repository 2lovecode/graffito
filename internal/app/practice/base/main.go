package base

import "context"

// Chapter 章节
type Chapter interface {
	GetQA(ctx context.Context, number int) QA
}

// QA QA
type QA interface {
	GetQ(ctx context.Context) Question
	GetA(ctx context.Context) Answer
}

// Question 问题
type Question interface {
	String() string
}

// Answer 答案
type Answer interface {
	String() string
	Run(ctx context.Context) string
}
