package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/afero"
	tips2 "graffito/tools/helper/tips"
	"io"
)

func main() {
	var AppFs = afero.NewOsFs()
	input, err := readFile(AppFs, "tools/helper/data/tips.json")

	if err == nil {
		tips := &tips2.TipGroups{}
		err = json.NewDecoder(bytes.NewReader(input)).Decode(tips)
		if err == nil {
			tips.CliPrintA()
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