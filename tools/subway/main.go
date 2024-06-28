package subway

// 从高德地图获取各个城市的地铁站坐标数据，并转换为百度坐标

import (
	"context"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/cobra"
	"io"
	"math"
	"net/http"
	"strconv"
	"strings"
)

const xPi = math.Pi * 3000.0 / 180.0
const subwayCityListApi = "https://webapi.amap.com/subway/data/citylist.json"

// GCJ02腾讯(高德)地图坐标转BD09百度坐标
func TX2BD(lng float64, lat float64) (toLng float64, toLat float64) {

	z := math.Sqrt(lng*lng+lat*lat) + 0.00002*math.Sin(lat*xPi)
	theta := math.Atan2(lat, lng) + 0.000003*math.Cos(lng*xPi)
	toLng = z*math.Cos(theta) + 0.0065
	toLat = z*math.Sin(theta) + 0.006
	return
}

// 百度地图BD09坐标转腾讯地图(高德)GCJ02坐标
func BD2TX(lng float64, lat float64) (toLng float64, toLat float64) {
	x := lng - 0.0065
	y := lat - 0.006

	z := math.Sqrt(x*x+y*y) - 0.00002*math.Sin(y*xPi)
	theta := math.Atan2(y, x) - 0.000003*math.Cos(x*xPi)
	toLng = z * math.Cos(theta)
	toLat = z * math.Sin(theta)
	return
}

type CityListResponse struct {
	CityList []*CityListEntity `json:"citylist"`
}

type CityListEntity struct {
	Spell    string `json:"spell"`
	ADCode   string `json:"adcode"`
	CityName string `json:"cityname"`
}

type SubwayResponse struct {
	Name   string       `json:"s"`
	ADCode string       `json:"i"`
	List   []SubwayLine `json:"l"`
	O      string       `json:"o"`
}

type SubwayLine struct {
	Name     string          `json:"ln"`
	FullName string          `json:"kn"`
	Su       string          `json:"su"`
	Station  []SubwayStation `json:"st"`
}

type SubwayStation struct {
	Name     string `json:"n"`
	Location string `json:"sl"`

	BDLng float64 `json:"bd_lng"`
	BDLat float64 `json:"bd_lat"`
}

func NewCommand() *cobra.Command {
	return &cobra.Command{Use: "fetch-subway-from-amap", Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			fmt.Println("请输入城市代码【北京:1100, 上海:3100, 广州:4401, 重庆:5000, 南京:3201, 武汉:4201, 成都:5101, 苏州:3205, 徐州:3203, 杭州:3301, 深圳:4403】")
			return
		}

		focusCity := []string{args[0]}
		client := &http.Client{}

		clResponse := &CityListResponse{}
		clReq, err := http.NewRequestWithContext(context.Background(), "GET", subwayCityListApi, nil)
		if err == nil {
			clRes, _ := client.Do(clReq)
			if clRes != nil {
				defer clRes.Body.Close()
				body, _ := io.ReadAll(clRes.Body)
				if len(body) > 0 {
					_ = jsoniter.Unmarshal(body, clResponse)
				}
			}
		}

		if len(clResponse.CityList) > 0 {
			urlList := make([]string, 0)
			for _, eachFocus := range focusCity {
				for _, each := range clResponse.CityList {
					if each.ADCode == eachFocus {
						urlList = append(urlList, fmt.Sprintf("https://webapi.amap.com/subway/data/%s_drw_%s.json", each.ADCode, each.Spell))
						break
					}
				}
			}

			for _, eachUrl := range urlList {
				subwayReq, subwayErr := http.NewRequestWithContext(context.Background(), "GET", eachUrl, nil)
				subwayResponse := &SubwayResponse{}
				if subwayErr == nil {
					subwayRes, _ := client.Do(subwayReq)
					if subwayRes != nil {
						defer subwayRes.Body.Close()

						body, _ := io.ReadAll(subwayRes.Body)
						if len(body) > 0 {
							_ = jsoniter.Unmarshal(body, subwayResponse)
						}
					}
				}

				if len(subwayResponse.List) > 0 {
					for k, v := range subwayResponse.List {
						for kk, vv := range v.Station {
							locs := strings.Split(vv.Location, ",")
							if len(locs) == 2 {
								lng, _ := strconv.ParseFloat(locs[0], 64)
								lat, _ := strconv.ParseFloat(locs[1], 64)
								if lng > 0 && lat > 0 {
									toLng, toLat := TX2BD(lng, lat)
									subwayResponse.List[k].Station[kk].BDLng = toLng
									subwayResponse.List[k].Station[kk].BDLat = toLat
								}
							}
						}
					}
					fmt.Println(subwayResponse.Name, subwayResponse.List)
				}
			}
		}
	}, Short: "从高德地图获取城市地铁信息，并将坐标转为百度坐标", Example: "获取北京地铁信息: {path/to/exe} tools fetch-subway 1100"}
}
