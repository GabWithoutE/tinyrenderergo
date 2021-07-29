package draw

import (
	"fmt"
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

func (i *ImgRGBA) SetPixel(x int, y int, color Color) error {
	p := i.RGBA.PixOffset(i.RGBA.Rect.Max.X-1-x, i.RGBA.Rect.Max.Y-1-y)

	// check for index out of range error, and return it for easier debugging
	if p >= len(i.RGBA.Pix)-1 {
		return fmt.Errorf("ImgRGBA.SetPixel: Index out of range, Index[%v] Range[%v]", p, len(i.RGBA.Pix))
	}

	i.RGBA.Pix[p] = color[0]
	i.RGBA.Pix[p+1] = color[1]
	i.RGBA.Pix[p+2] = color[2]
	i.RGBA.Pix[p+3] = color[3]

	return nil
}
