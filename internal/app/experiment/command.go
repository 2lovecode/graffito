package experiment

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os/exec"
	"time"

	"github.com/2lovecode/graffito/internal/app/experiment/concx"
	"github.com/2lovecode/graffito/internal/app/experiment/depends"
	"github.com/2lovecode/graffito/internal/app/experiment/event"
	"github.com/2lovecode/graffito/internal/app/experiment/mode0"
	"github.com/2lovecode/graffito/internal/app/experiment/mode0/card"
	"github.com/2lovecode/graffito/internal/app/experiment/mode0/handler"
	"github.com/2lovecode/graffito/internal/app/experiment/mode0/source"
	"github.com/2lovecode/graffito/pkg/cache"
	"github.com/2lovecode/graffito/pkg/params"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/ianaindex"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"

	"github.com/2lovecode/apiflow"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cacheCmd := &cobra.Command{Use: "cross-cache", Run: func(cmd *cobra.Command, args []string) {
		cache.CrossRun()
	}, Short: "缓存穿透"}

	eventCmd := &cobra.Command{Use: "event", Run: func(cmd *cobra.Command, args []string) {
		mEvent := event.NewMEvent()

		myC := &event.MyCar{}
		myP := &event.MyPhone{}

		mEvent.On("car", myC)
		mEvent.On("car", myP)

		p := params.NewPayload()
		p.Set("name", "vivo")

		mEvent.Trigger("car", p)
	}, Short: "事件"}

	mode0Cmd := &cobra.Command{Use: "mode0", Run: func(cmd *cobra.Command, args []string) {

		//数据聚合层 - 可以替换数据来源
		s0 := source.NewS0()
		aggLayer := mode0.NewDataAggregationLayer(mode0.Source(s0))
		aggData := aggLayer.Output()

		//数据处理层 - 可注册多个处理器
		dealerLayer := mode0.NewDealerLayer(aggData)

		//生成处理器 - 每个处理器可注册不同的处理逻辑
		bigCard := card.NewBigCard()
		bigNameHandler := handler.NewNameBig()
		//...
		bigCard.Register(bigNameHandler)

		smallCard := card.NewSmallCard()
		smallNameHandler := handler.NewNameSmall()
		//...
		smallCard.Register(smallNameHandler)

		dealerLayer.RegisterCard(bigCard)
		dealerLayer.RegisterCard(smallCard)

		ret := dealerLayer.Output()

		fmt.Println(ret[0], "{\"name\":\"type_big_Big_Big_Hooooooo\"}")
		fmt.Println(ret[1], "{\"name\":\"Hoooooo_type_small\"}")
	}, Short: "mode0"}

	dependsCmd := &cobra.Command{Use: "depends", Run: func(cmd *cobra.Command, args []string) {
		ctx := context.WithValue(context.TODO(), "q", "test")

		sA := depends.NewServiceA()
		sB := depends.NewServiceB()
		sC := depends.NewServiceC()

		hd := depends.NewDepends(100 * time.Millisecond)

		hd.Register(sA)
		hd.Register(sB)
		hd.Register(sC)

		hd.AddDepend(sC, []depends.IService{sB})
		hd.AddDepend(sB, []depends.IService{sA})

		hd.Execute(ctx)

		sAData := depends.ServiceAData{}
		sBData := depends.ServiceBData{}
		sCData := depends.ServiceCData{}

		sA.Decode(&sAData)
		sB.Decode(&sBData)
		sC.Decode(&sCData)

		fmt.Println(sAData, sBData, sCData)
	}, Short: "depends"}

	pingCmd := &cobra.Command{
		Use: "ping",
		Run: func(cmd *cobra.Command, args []string) {
			p := exec.Command("powershell", "-Command", "ping", "www.baidu.com")

			var out bytes.Buffer

			p.Stdout = &out

			_ = p.Run()

			// encoding, name, certain := charset.DetermineEncoding(out.Bytes(), "")
			// fmt.Printf("编码：%v\n名称：%s\n确定：%t\n", encoding, name, certain)

			output, _ := GbkToUtf8(out.Bytes())
			fmt.Println("Out: ", string(output))

		},
	}

	expCmd := &cobra.Command{Use: "exp", Short: "试验代码"}

	apiflowCmd := &cobra.Command{Use: "apiflow", Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		flow := apiflow.NewApiFlow(1 * time.Second)

		nodeA := apiflow.NewNode("A", func(ctx context.Context, node *apiflow.Node, inputs map[string]interface{}) (interface{}, error) {
			fmt.Printf("Executing node %s\n", node.ID)
			time.Sleep(100 * time.Millisecond)
			m := make(map[string]string)
			m["node_a_1"] = "goodsA"
			m["node_a_2"] = "morningA"
			return m, nil
		})

		nodeB := apiflow.NewNode("B", func(ctx context.Context, node *apiflow.Node, inputs map[string]interface{}) (interface{}, error) {
			fmt.Printf("Executing node %s\n", node.ID)
			if inputs != nil {
				if aa, ok := inputs["A"]; ok {
					fmt.Println("bbbbbb: ", aa)
				}
			}
			time.Sleep(200 * time.Millisecond)
			m := make(map[string]string)
			m["node_b_1"] = "goodB"
			m["node_b_2"] = "morningB"
			return m, nil
		})

		nodeC := apiflow.NewNode("C", func(ctx context.Context, node *apiflow.Node, inputs map[string]interface{}) (interface{}, error) {
			fmt.Printf("Executing node %s\n", node.ID)
			if inputs != nil {
				if aa, ok := inputs["A"]; ok {
					fmt.Println("cccc: ", aa)
				}
				if bb, ok := inputs["B"]; ok {
					fmt.Println("cccc: ", bb)
				}
			}
			time.Sleep(100 * time.Millisecond)
			return nil, fmt.Errorf("error in node %s", node.ID)
		})

		nodeD := apiflow.NewNode("D", func(ctx context.Context, node *apiflow.Node, inputs map[string]interface{}) (interface{}, error) {
			fmt.Printf("Executing node %s\n", node.ID)
			time.Sleep(300 * time.Millisecond)
			return nil, nil
		})

		// Adding nodes with dependencies
		flow.AddNode(nodeA, []string{})
		flow.AddNode(nodeB, []string{"A"})
		flow.AddNode(nodeC, []string{"B"})
		flow.AddNode(nodeD, []string{"A"})

		// Run the API flow

		flow.Run(ctx)
	}, Short: "apiflow"}
	cmds := []*cobra.Command{
		cacheCmd,
		dependsCmd,
		mode0Cmd,
		eventCmd,
		pingCmd,
		apiflowCmd,
		concx.NewCommand(),
	}

	expCmd.AddCommand(cmds...)
	return expCmd
}

type Charset string

// 中文
const (
	GBK     Charset = "GBK"
	GB18030         = "GB18030"
	GB2312          = "GB2312"
	Big5            = "Big5"
)

// 日文
const (
	EUCJP     Charset = "EUCJP"
	ISO2022JP         = "ISO2022JP"
	ShiftJIS          = "ShiftJIS"
)

// 韩文
const (
	EUCKR Charset = "EUCKR"
)

// Unicode
const (
	UTF_8    Charset = "UTF-8"
	UTF_16           = "UTF-16"
	UTF_16BE         = "UTF-16BE"
	UTF_16LE         = "UTF-16LE"
)

// 其他编码
const (
	Macintosh Charset = "macintosh"
	IBM               = "IBM*"
	Windows           = "Windows*"
	ISO               = "ISO-*"
)

var charsetAlias = map[string]string{
	"HZGB2312": "HZ-GB-2312",
	"hzgb2312": "HZ-GB-2312",
	"GB2312":   "HZ-GB-2312",
	"gb2312":   "HZ-GB-2312",
}

func Convert(dstCharset Charset, srcCharset Charset, src string) (dst string, err error) {
	if dstCharset == srcCharset {
		return src, nil
	}
	dst = src
	// Converting <src> to UTF-8.
	if srcCharset != "UTF-8" {
		if e := getEncoding(srcCharset); e != nil {
			tmp, err := ioutil.ReadAll(
				transform.NewReader(bytes.NewReader([]byte(src)), e.NewDecoder()),
			)
			if err != nil {
				return "", fmt.Errorf("%s to utf8 failed. %v", srcCharset, err)
			}
			src = string(tmp)
		} else {
			return dst, errors.New(fmt.Sprintf("unsupport srcCharset: %s", srcCharset))
		}
	}
	// Do the converting from UTF-8 to <dstCharset>.
	if dstCharset != "UTF-8" {
		if e := getEncoding(dstCharset); e != nil {
			tmp, err := ioutil.ReadAll(
				transform.NewReader(bytes.NewReader([]byte(src)), e.NewEncoder()),
			)
			if err != nil {
				return "", fmt.Errorf("utf to %s failed. %v", dstCharset, err)
			}
			dst = string(tmp)
		} else {
			return dst, errors.New(fmt.Sprintf("unsupport dstCharset: %s", dstCharset))
		}
	} else {
		dst = src
	}
	return dst, nil
}

func ToUTF8(srcCharset Charset, src string) (dst string, err error) {
	return Convert("UTF-8", srcCharset, src)
}

func UTF8To(dstCharset Charset, src string) (dst string, err error) {
	return Convert(dstCharset, "UTF-8", src)
}

func getEncoding(charset Charset) encoding.Encoding {
	if c, ok := charsetAlias[string(charset)]; ok {
		charset = Charset(c)
	}
	if e, err := ianaindex.MIB.Encoding(string(charset)); err == nil && e != nil {
		return e
	}
	return nil
}

func GbkToUtf8(s []byte) ([]byte, error) {
	//第二个参数为“transform.Transformer”接口，simplifiedchinese.GBK.NewDecoder()包含了该接口
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}
