package practice

import (
	"graffito/practice/bingfa"
	"testing"
)

func Benchmark_BingfaRun(t *testing.B) {
	t.RunParallel(func(pb *testing.PB) {
		bingfa.Run()
	})
}
