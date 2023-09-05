package cmd

import (
	"bytes"
	"context"
	"fmt"
	"github.com/speps/go-hashids/v2"
	"github.com/spf13/cobra"
	"graffito/app/sandbox"
	"io"
	"math/rand"
	"os"
)

func NewCliCommand() *cobra.Command {
	file := ""
	source := ""

	cli := &cobra.Command{Use: "cli", Run: func(cmd *cobra.Command, args []string) {
		sourceCode := ""
		if source != "" {
			sourceCode = source
		} else if file != "" {
			f, err := os.Open(file)
			if err != nil {
				fmt.Println(err)
				return
			}
			var buffer bytes.Buffer
			_, err = io.Copy(&buffer, f)
			if err != nil {
				fmt.Println(err)
			}
			sourceCode = buffer.String()
		}

		sandboxApp := sandbox.NewApplication()
		out, err := sandboxApp.Exec(context.Background(), &sandbox.Input{
			SourceCode: sourceCode,
		})

		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		so := sandbox.Output{}
		err = out.To(&so)

		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		fmt.Println("Output: ", so.Data)
	}}
	cli.Flags().StringVarP(&file, "file", "f", "", "指定文件")
	cli.Flags().StringVarP(&source, "source", "s", "", "指定源代码")
	return cli
}

func DefaultHash() *hashids.HashID {
	hash := hashids.NewData()
	hash.Salt = "APpAUptXsssdfffww"
	hash.MinLength = 6
	hash.Alphabet = "abcdefghijklmnopqrstuvwxyz1234567890+-*=@#$%"
	h, _ := hashids.NewWithData(hash)
	return h
}

func EncodeHash(num int64) string {
	h := DefaultHash()

	e, _ := h.EncodeInt64([]int64{num})
	return e
}

func DecodeHash(str string) int64 {
	h := DefaultHash()
	d, _ := h.DecodeInt64WithError(str)
	return d[0]
}

func consistentShuffleInPlace(alphabet []rune, salt []rune) {
	if len(salt) == 0 {
		return
	}

	for i, v, p := len(alphabet)-1, 0, 0; i > 0; i-- {
		p += int(salt[v])
		j := (int(salt[v]) + v + p) % i
		alphabet[i], alphabet[j] = alphabet[j], alphabet[i]
		v = (v + 1) % len(salt)
	}
	return
}

func fisherYatesShuffleWithConsistency(arr []int, seed int64) {
	rng := rand.New(rand.NewSource(seed))

	for i := len(arr) - 1; i > 0; i-- {
		j := rng.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
}
