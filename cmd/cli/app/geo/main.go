package geo

import (
	"fmt"
	"github.com/spf13/cobra"
	"math"
)

func NewTransCommand() *cobra.Command {
	cmd := &cobra.Command{Use: "geo-trans", Run: func(cmd *cobra.Command, args []string) {
		xPi := math.Pi * 3000.0 / 180.0

		lng := 116.393776
		lat := 39.948972

		z := math.Sqrt(lng*lng+lat*lat) + 0.00002*math.Sin(lat*xPi)
		theta := math.Atan2(lat, lng) + 0.000003*math.Cos(lng*xPi)
		toLng := z*math.Cos(theta) + 0.0065
		toLat := z*math.Sin(theta) + 0.006

		fmt.Println(toLng, toLat)
	}, Short: "坐标转换"}
	return cmd
}
