package poi

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"net/url"
)

func NewPOICommand() *cobra.Command {
	cmd := &cobra.Command{Use: "poi", Run: func(cmd *cobra.Command, args []string) {
		// 此处填写您在控制台-应用管理-创建应用后获取的AK
		if len(args) <= 0 {
			fmt.Println("请输入您的ak")
			return
		}
		if len(args) <= 2 {
			fmt.Println("请输入坐标 经度 纬度")
			return
		}
		ak := args[0]
		lng := args[1]
		lat := args[2]

		// 服务地址
		host := "https://api.map.baidu.com"

		// 接口地址
		uri := "/place/v2/search"

		categories := []string{"地铁站", "公交车站", "购物中心", "百货商场", "超市", "便利店", "公园", "中餐厅", "外国餐厅", "小吃快餐店", "咖啡厅", "运动健身", "电影院", "KTV", "剧院", "综合医院", "诊所", "高等院校", "中学", "小学", "幼儿园"}

		//categories := []string{"地铁站", "公交车站"}
		l := len(categories)
		fmt.Println("{")
		for k, category := range categories {
			// 设置请求参数
			params := url.Values{
				"query":        []string{category},
				"location":     []string{fmt.Sprintf("%s,%s", lat, lng)},
				"radius":       []string{"3000"},
				"output":       []string{"json"},
				"ak":           []string{ak},
				"scope":        []string{"2"},
				"radius_limit": []string{"true"},
				"filter":       []string{"sort_name:distance|sort_rule:1"},
			}

			// 发起请求
			req, err := url.Parse(host + uri + "?" + params.Encode())
			if err != nil {
				fmt.Printf("url parse error: %v", err)
				continue
			}
			res, resErr := http.Get(req.String())
			if resErr != nil {
				fmt.Printf("request error: %v", resErr)
				continue
			}

			if res == nil {
				continue
			}

			if res.Body != nil {
				resByt, readErr := io.ReadAll(res.Body)
				_ = res.Body.Close()
				if readErr != nil {
					fmt.Printf("read error: %v", readErr)
					continue
				}
				sm := ""
				if k < l-1 {
					sm = ","
				}
				fmt.Printf("\"%s\":%s%s\n", category, string(resByt), sm)
			}
		}

		fmt.Println("}")

		return
	}, Short: "poi搜索"}
	return cmd
}
