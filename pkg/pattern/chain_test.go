package pattern

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var chainNum = 0

type ChainHandler1 struct {
	ChainNext
}

func (h *ChainHandler1) Do(ctx context.Context) error {
	chainNum++
	return nil
}

type ChainHandler2 struct {
	ChainNext
}

func (h *ChainHandler2) Do(ctx context.Context) error {
	chainNum += 2
	return nil
}

func Test_Chain(t *testing.T) {
	Convey("Chain", t, func() {
		ctx := context.Background()
		nilH := &NilChainHandler{}
		nilH.SetNext(&ChainHandler1{}).SetNext(&ChainHandler2{})
		nilH.Run(ctx)
		So(chainNum, ShouldEqual, 3)
	})

}
