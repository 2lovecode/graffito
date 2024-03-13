package experiment

import (
	"fmt"
	"graffito/experiment/ginx"
	"net/http"
	"testing"
	"time"
)

func TestGinx_Run(t *testing.T) {

	engine := ginx.NewEngine()

	if err := http.ListenAndServe(":8888", engine); err != nil {
		fmt.Println(err)
	}
	time.Sleep(30 * time.Second)
}
