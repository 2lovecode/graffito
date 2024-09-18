package poi

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func NewPOICommand() *cobra.Command {
	cmd := &cobra.Command{Use: "poi", Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(time.Now().Day())
		// 此处填写您在控制台-应用管理-创建应用后获取的AK
		if len(args) <= 0 {
			fmt.Println("请输入您的ak")
			return
		}
		ak := args[0]

		// 服务地址
		host := "https://api.map.baidu.com"

		// 接口地址
		uri := "/place/v2/search"

		// 设置请求参数
		params := url.Values{
			"query":    []string{"中餐厅"},
			"location": []string{"39.93118123,116.53187525"},
			"radius":   []string{"3000"},
			//"tag":          []string{"高等院校"},
			"output":       []string{"json"},
			"ak":           []string{ak},
			"scope":        []string{"2"},
			"radius_limit": []string{"true"},
			"filter":       []string{"sort_name:distance|sort_rule:1"},
		}

		// 发起请求
		request, err := url.Parse(host + uri + "?" + params.Encode())
		if nil != err {
			fmt.Printf("host error: %v", err)
			return
		}

		resp, err1 := http.Get(request.String())
		fmt.Printf("url: %s\n", request.String())
		defer resp.Body.Close()
		if err1 != nil {
			fmt.Printf("request error: %v", err1)
			return
		}
		body, err2 := ioutil.ReadAll(resp.Body)
		if err2 != nil {
			fmt.Printf("response error: %v", err2)
		}

		fmt.Println(string(body))

		return
	}, Short: "poi搜索"}
	return cmd
}
