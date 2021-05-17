package draw

import (
	"image"
)

type ImgRGBA struct {
	*image.RGBA
}

type Color [4]uint8

func NewRGBAImage(rectangle image.Rectangle) *ImgRGBA {
	return &ImgRGBA{
		image.NewRGBA(rectangle),
	}
}

func (i *ImgRGBA) Set(x int, y int, color Color) {
	p := i.RGBA.PixOffset(i.RGBA.Rect.Max.X-1-x, i.RGBA.Rect.Max.Y-1-y)
	i.RGBA.Pix[p] = color[0]
	i.RGBA.Pix[p+1] = color[1]
	i.RGBA.Pix[p+2] = color[2]
	i.RGBA.Pix[p+3] = color[3]
}
