package xmutex

import (
	"context"
	"github.com/2lovecode/graffito/internal/app/practice/base"
)

// Golang Mutex V0
// https://github.com/golang/go/blob/bf3dd3f0efe5b45947a991e22660c62d4ce6b671/src/lib/sync/mutex.go

type SimpleMutexImplV0 struct {
}

func NewSimpleMutexImplV0() *SimpleMutexImplV0 {
	return &SimpleMutexImplV0{}
}

func (s *SimpleMutexImplV0) Name() base.Name {
	return base.SimpleMutexImplV0
}

func (s *SimpleMutexImplV0) Description() string {
	return "https://github.com/golang/go/blob/bf3dd3f0efe5b45947a991e22660c62d4ce6b671/src/lib/sync/mutex.go"
}

func (s *SimpleMutexImplV0) Run(ctx context.Context) (string, error) {
	return s.Description(), nil
}
