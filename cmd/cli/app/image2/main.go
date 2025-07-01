package image2

import (
	"fmt"
	"image"
	"image/color"
	"net/http"

	"github.com/2lovecode/graffito/pkg/fs"
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{Use: "image", Run: func(cmd *cobra.Command, args []string) {

		appPath := fs.GetCurrentAbsPath()
		dataDirName := "data"

		dataDirPath := fmt.Sprintf("%s/%s", appPath, dataDirName)
		fontFile := fmt.Sprintf("%s/%s/%s", appPath, "fonts", "sanjiheisongti.ttf")
		locationIcon := fmt.Sprintf("%s/%s/%s", appPath, "images", "location.png")
		dstFile := fmt.Sprintf("%s/%s", dataDirPath, "dst.jpg")
		outFile := fmt.Sprintf("%s/%s", dataDirPath, "out.png")
		if resp, err := http.Get("https://imgpro.ziroom.com/70f3d32d-f94f-4724-b0b7-f0b20fb91451_110.jpg_C_800_600_Q90.jpg"); err == nil {
			defer resp.Body.Close()

			if img, err2 := imaging.Decode(resp.Body); err2 == nil {
				dst := imaging.Resize(img, 100, 0, imaging.Lanczos)
				imaging.Save(dst, dstFile)
			} else {
				fmt.Println("err2", err2)
			}
		} else {
			fmt.Println("err", err)
		}

		// 加载 location 图标
		// 加载原始图标
		iconRaw, err := imaging.Open(locationIcon)
		if err != nil {
			panic("加载location图标失败: " + err.Error())
		}

		// 设置缩放后的尺寸
		iconW := 40
		iconH := 40
		iconImg := imaging.Resize(iconRaw, iconW, iconH, imaging.Lanczos)

		// icon := file
		if respg, errg := http.Get("https://imgpro.ziroom.com/70f3d32d-f94f-4724-b0b7-f0b20fb91451_110.jpg_C_800_600_Q90.jpg"); errg == nil {
			defer respg.Body.Close()

			if img, _, errg2 := image.Decode(respg.Body); errg2 == nil {

				// 获取原图尺寸
				imgW := img.Bounds().Dx()
				imgH := img.Bounds().Dy()

				// 设置文本内容和区域高度
				infoHeight := 150

				// 创建新画布（高度 = 原图高度 + 文字区域）
				dc := gg.NewContext(imgW, imgH+int(infoHeight))

				// 绘制背景为白色
				dc.SetRGB(1, 1, 1)
				dc.Clear()

				// 绘制原图在顶部
				dc.DrawImage(img, 0, 0)

				// 1. 渐变设置（左→右）
				grad := gg.NewLinearGradient(0, 0, 200, 0)
				grad.AddColorStop(0, color.RGBA{205, 192, 168, 255}) // 金色亮
				grad.AddColorStop(1, color.RGBA{230, 225, 223, 255}) // 金色暗

				// 2. 绘制圆角矩形背景（左上角）
				rectW := 250.0
				rectH := 50.0
				radius := 12.0
				rectWM := 10.0
				rectHM := 10.0

				dc.SetFillStyle(grad)
				dc.DrawRoundedRectangle(rectWM, rectHM, rectW, rectH, radius)
				dc.Fill()

				// 3. 设置字体
				if err := dc.LoadFontFace(fontFile, 24); err != nil {
					panic(err)
				}

				// 4. 设置文字颜色（深金色）并绘制居中文字
				dc.SetRGB255(126, 99, 86)
				dc.DrawStringAnchored("焕新好房 No.00012", rectWM+rectW/2, rectHM+rectH/2, 0.5, 0.5)

				// 设置字体（根据你的系统修改路径）
				if err := dc.LoadFontFace(fontFile, 28); err != nil {
					panic(err)
				}

				dc.SetRGBA(1, 1, 1, 0.6) // 白色 + 60% 透明

				// 水印文字
				text := "朝阳区"

				// 获取文字宽高
				_, th := dc.MeasureString(text)

				// 设置水印位置（右下角，预留10px边距）
				margin := 15.0
				xText := margin + float64(iconW) + 5 // 文字 x 坐标
				y := float64(imgH) - th - margin

				// 先绘制图标
				dc.DrawImageAnchored(iconImg, int(margin+float64(iconW)/2), int(y+th/2), 0.5, 0.5)

				// 绘制文字
				dc.DrawStringAnchored(text, xText, y+th/2, 0, 0.5)
				// 设置文案内容
				titles := []string{"120万", "3室1厅1卫", "143.56㎡"}
				labels := []string{"售价", "户型", "面积"}

				// 均分排版
				cols := len(titles)
				colWidth := float64(imgW) / float64(cols)
				// 第一行：加粗加大
				if err := dc.LoadFontFace(fontFile, 36); err != nil {
					panic(err)
				}
				dc.SetRGB(0, 0, 0)
				for i, text := range titles {
					x := colWidth*float64(i) + colWidth/2
					y := float64(imgH) + 45
					dc.DrawStringAnchored(text, x, y, 0.5, 0.5)
				}

				// 第二行：小字说明
				if err := dc.LoadFontFace(fontFile, 20); err != nil {
					panic(err)
				}
				dc.SetRGB(0.4, 0.4, 0.4) // 灰色
				for i, text := range labels {
					x := colWidth*float64(i) + colWidth/2
					y := float64(imgH) + 90
					dc.DrawStringAnchored(text, x, y, 0.5, 0.5)
				}

				// 保存合成图
				dc.SavePNG(outFile)
			} else {
				fmt.Println("errg2", errg2)
			}
		} else {
			fmt.Println("errg", errg)
		}

	}, Short: "图片合成"}
	return cmd
}
