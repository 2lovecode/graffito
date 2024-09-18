package experiment

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/2lovecode/graffito/internal/app/experiment/ginx"
)

func TestGinx_Run(t *testing.T) {

	engine := ginx.NewEngine()

	if err := http.ListenAndServe(":8888", engine); err != nil {
		fmt.Println(err)
	}
	time.Sleep(30 * time.Second)
}
