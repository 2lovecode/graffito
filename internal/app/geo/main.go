package geo

import (
	"fmt"
	"math"

	"github.com/2lovecode/graffito/pkg/logging"
	"github.com/spf13/cobra"
)

func NewTransCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "geo-trans",
		Short: "地理坐标转换工具",
		Long: `地理坐标转换工具，支持不同坐标系之间的转换。

当前示例使用百度坐标（BD-09）转换为GCJ-02坐标系。

示例:
  graffito cli tools geo-trans
`,
		Run: func(cmd *cobra.Command, args []string) {
			xPi := math.Pi * 3000.0 / 180.0

			// 示例坐标（百度坐标）
			lng := 116.393776
			lat := 39.948972

			z := math.Sqrt(lng*lng+lat*lat) + 0.00002*math.Sin(lat*xPi)
			theta := math.Atan2(lat, lng) + 0.000003*math.Cos(lng*xPi)
			toLng := z*math.Cos(theta) + 0.0065
			toLat := z*math.Sin(theta) + 0.006

			logging.Debugf("坐标转换: (%.6f, %.6f) -> (%.6f, %.6f)", lng, lat, toLng, toLat)
			fmt.Fprintf(cmd.OutOrStdout(), "转换结果:\n经度: %.6f\n纬度: %.6f\n", toLng, toLat)
		},
	}
	return cmd
}
