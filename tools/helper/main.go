package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	tips2 "graffito/tools/helper/tips"
	"io"

	"github.com/spf13/afero"
)

type Helper struct {
	dataPath string
}

func NewHelper(dataPath string) *Helper {
	if dataPath == "" {
		dataPath = "tools/helper/tips.json"
	}
	return &Helper{dataPath: dataPath}
}

func (h *Helper) Run() {
	var AppFs = afero.NewOsFs()
	input, err := readFile(AppFs, h.dataPath)

	if err == nil {
		tips := &tips2.TipGroups{}
		err = json.NewDecoder(bytes.NewReader(input)).Decode(tips)
		if err == nil {
			tips.CliPrintB()
		}
	}
	fmt.Println(err)
}

func readFile(fs afero.Fs, filename string) ([]byte, error) {
	f, err := fs.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var n int64

	if fi, err := f.Stat(); err == nil {
		// Don't preallocate a huge buffer, just in case.
		if size := fi.Size(); size < 1e9 {
			n = size
		}
	}

	return readAll(f, n+bytes.MinRead)
}

func readAll(r io.Reader, capacity int64) (b []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0, capacity))
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
			err = panicErr
		} else {
			panic(e)
		}
	}()
	_, err = buf.ReadFrom(r)
	return buf.Bytes(), err
}