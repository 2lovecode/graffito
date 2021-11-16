package color

import (
	"fmt"
	"image/color"
	"regexp"
	"strconv"
)

type ColorHandler interface {
	Decode(c string) color.RGBA
	Encode(c color.RGBA) string
}

type HEXColorHandler struct{}

func (h *HEXColorHandler) Decode(c string) color.RGBA {
	cObj := color.RGBA{R: 0, G: 0, B: 0, A: 0}
	if len(c) == 7 {
		cObj.A = uint8(255)

		r, _ := strconv.ParseUint(c[1:3], 16, 8)
		cObj.R = uint8(r)

		g, _ := strconv.ParseUint(c[3:5], 16, 8)
		cObj.G = uint8(g)

		b, _ := strconv.ParseUint(c[5:7], 16, 8)
		cObj.B = uint8(b)
	}
	return cObj
}
func (h *HEXColorHandler) Encode(c color.RGBA) string {
	r := fmt.Sprintf("%0.2X", c.R)
	g := fmt.Sprintf("%0.2X", c.G)
	b := fmt.Sprintf("%0.2X", c.B)
	return fmt.Sprintf("#%s%s%s", r, g, b)
}

type HEXAColorHandler struct{}

func (h *HEXAColorHandler) Decode(c string) color.RGBA {
	cObj := color.RGBA{R: 0, G: 0, B: 0, A: 0}
	if len(c) == 9 {
		r, _ := strconv.ParseUint(c[1:3], 16, 8)
		cObj.R = uint8(r)

		g, _ := strconv.ParseUint(c[3:5], 16, 8)
		cObj.G = uint8(g)

		b, _ := strconv.ParseUint(c[5:7], 16, 8)
		cObj.B = uint8(b)

		a, _ := strconv.ParseUint(c[7:9], 16, 8)
		cObj.A = uint8(a)
	}
	return cObj
}
func (h *HEXAColorHandler) Encode(c color.RGBA) string {
	a := fmt.Sprintf("%0.2X", c.A)
	r := fmt.Sprintf("%0.2X", c.R)
	g := fmt.Sprintf("%0.2X", c.G)
	b := fmt.Sprintf("%0.2X", c.B)
	return fmt.Sprintf("#%s%s%s%s", r, g, b, a)
}

type AHEXColorHandler struct{}

func (h *AHEXColorHandler) Decode(c string) color.RGBA {
	cObj := color.RGBA{R: 0, G: 0, B: 0, A: 0}

	if len(c) == 9 {
		a, _ := strconv.ParseUint(c[1:3], 16, 8)
		cObj.A = uint8(a)

		r, _ := strconv.ParseUint(c[3:5], 16, 8)
		cObj.R = uint8(r)

		g, _ := strconv.ParseUint(c[5:7], 16, 8)
		cObj.G = uint8(g)

		b, _ := strconv.ParseUint(c[7:9], 16, 8)
		cObj.B = uint8(b)
	}

	return cObj
}
func (h *AHEXColorHandler) Encode(c color.RGBA) string {
	a := fmt.Sprintf("%0.2X", c.A)
	r := fmt.Sprintf("%0.2X", c.R)
	g := fmt.Sprintf("%0.2X", c.G)
	b := fmt.Sprintf("%0.2X", c.B)
	return fmt.Sprintf("#%s%s%s%s", a, r, g, b)
}

type RGBColorHandler struct{}

func (h *RGBColorHandler) Decode(c string) color.RGBA {
	cObj := color.RGBA{R: 0, G: 0, B: 0, A: 0}
	if reg, ok := ColorRegxs[ColorFormatRGB]; ok {
		re, err := regexp.Compile(reg)
		if err == nil {
			cs := re.FindStringSubmatch(c)
			if len(cs) == 5 {
				r, _ := strconv.ParseUint(cs[2], 10, 8)
				g, _ := strconv.ParseUint(cs[3], 10, 8)
				b, _ := strconv.ParseUint(cs[4], 10, 8)
				cObj.R = uint8(r)
				cObj.G = uint8(g)
				cObj.B = uint8(b)
				cObj.A = 255
			}
		}
	}
	return cObj
}
func (h *RGBColorHandler) Encode(c color.RGBA) string {
	return fmt.Sprintf("rgb(%d, %d, %d)", c.R, c.G, c.B)
}

type RGBAColorHandler struct{}

func (h *RGBAColorHandler) Decode(c string) color.RGBA {
	cObj := color.RGBA{R: 0, G: 0, B: 0, A: 0}

	if reg, ok := ColorRegxs[ColorFormatRGBA]; ok {
		re, err := regexp.Compile(reg)
		if err == nil {
			cs := re.FindStringSubmatch(c)
			if len(cs) == 6 {
				r, _ := strconv.ParseUint(cs[2], 10, 8)
				g, _ := strconv.ParseUint(cs[3], 10, 8)
				b, _ := strconv.ParseUint(cs[4], 10, 8)
				a, _ := strconv.ParseFloat(cs[5], 32)
				cObj.R = uint8(r)
				cObj.G = uint8(g)
				cObj.B = uint8(b)
				cObj.A = uint8(a * 255)
			}
		}
	}

	return cObj
}
func (h *RGBAColorHandler) Encode(c color.RGBA) string {
	return fmt.Sprintf("rgba(%d, %d, %d, %s)", c.R, c.G, c.B, strconv.FormatFloat(float64(c.A)/255, 'f', 2, 64))
}

// type HSLColorHandler struct{}

// func (h *HSLColorHandler) Decode(c string) color.RGBA {
// 	cObj := color.RGBA{R: 0, G: 0, B: 0, A: 0}

// 	return cObj
// }
// func (h *HSLColorHandler) Encode(c color.RGBA) string {
// 	return ""
// }

// type HSLAColorHandler struct{}

// func (h *HSLAColorHandler) Decode(c string) color.RGBA {
// 	cObj := color.RGBA{R: 0, G: 0, B: 0, A: 0}

// 	return cObj
// }
// func (h *HSLAColorHandler) Encode(c color.RGBA) string {
// 	return ""
// }
