package helper

import (
	"bytes"
	"encoding/json"
	tips2 "github.com/2lovecode/graffito/tools/helper/tips"
	"github.com/spf13/cobra"
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

func NewCommand() *cobra.Command {
	helpCmd := &cobra.Command{Use: "helper", Run: func(cmd *cobra.Command, args []string) {
		str := ""
		if len(args) > 0 {
			str = args[0]
		}
		h := NewHelper(str)
		h.Run()
	}, Short: "命令备忘录，可以添加命令备忘，列出所有命令", Example: "{path/to/exe} tools helper"}

	//helpCmd.AddCommand([]*cobra.Command{
	//	cmd.NewAddCommand(),
	//}...)

	return helpCmd
}
