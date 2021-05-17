package draw

import (
	"github.com/pkg/errors"
	"math"
)

func (i *ImgRGBA) Line(x0 int, y0 int, x1 int, y1 int, color Color) error {
	steep := false
	// flipping x and y when the line is steep allows for the rounding to be
	// 		done at higher resolution.
	if math.Abs(float64(x0-x1)) < math.Abs(float64(y0-y1)) {
		x0, y0 = y0, x0
		x1, y1 = y1, x1
		steep = true
	}

	// compute pixels from low to high
	if x0 > x1 {
		x0, x1 = x1, x0
		y0, y1 = y1, y0
	}

	dx := float64(x1 - x0)
	dy := float64(y1 - y0)
	derror2 := math.Abs(dy) * 2
	error2 := 0.0
	y := y0

	for x := x0; x <= x1; x++ {
		if steep {
			if err := i.SetPixel(y, x, color); err != nil {
				return errors.Wrap(err, "draw.Line: SetPixel error")
			}
		} else {
			if err := i.SetPixel(x, y, color); err != nil {
				return errors.Wrap(err, "draw.Line: SetPixel error")
			}
		}

		error2 += derror2
		if error2 > dx {
			if dy > 0 {
				y = y + 1
			} else {
				y = y - 1
			}
			error2 -= dx * 2
		}
	}

	return nil
}
