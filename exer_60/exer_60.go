// package main

// import (
// 	"image"
// 	"image/color"

// 	"golang.org/x/tour/pic"
// )

// type Image struct{}

// func (i Image) ColorModel() color.Model {
// 	return color.RGBAModel
// }

// func (i Image) Bounds() image.Rectangle {
// 	return image.Rect(0, 0, 256, 256)
// }

// func (i Image) At(x, y int) color.Color {
// 	v := uint8(x ^ y)
// 	// v := uint8((x + y) / 2)
// 	// v := uint8(x * y)
// 	return color.RGBA{v, v, 255, 255}
// }

// func main() {
// 	m := Image{}
// 	pic.ShowImage(m)
// }

package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct{}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 256, 256)
}

func (i Image) At(x, y int) color.Color {
	v := uint8(x ^ y)
	// v := uint8((x + y) / 2)
	// v := uint8(x * y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
