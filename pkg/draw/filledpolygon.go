package draw

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/pkg/errors"
)

type FillMethod string

const (
	LineSweep   = FillMethod("LineSweep")
	BoundingBox = FillMethod("BoundingBox")
)

func (i *ImgRGBA) DrawFilledTriangle(method FillMethod, vertices [3]mgl32.Vec4, color Color) error {
	switch method {
	case LineSweep:
		if err := i.lineSweep(vertices, color); err != nil {
			return errors.Wrap(err, "draw.DrawFilledTriangle")
		}
		return nil
	case BoundingBox:
		if err := i.boundingBox(vertices, color); err != nil {
			return errors.Wrap(err, "draw.DrawFilledTriangle")
		}
		return nil
	default:
		return errors.Errorf("draw.DrawFilledTriangle invalid FillMethod %v", method)
	}
}

func (i *ImgRGBA) lineSweep(vertices [3]mgl32.Vec4, color Color) error {
	if err := sortTriangleVerticesByYDesc(&vertices); err != nil {
		return errors.Wrap(err, "draw.lineSweep")
	}

	currentLine := vertices[0].Y()
	if vertices[0].Y()-vertices[2].Y() == 0 {
		return errors.Errorf("draw.DrawFilledTriangle.lineSweep: invalid triangle definition vertices %v", vertices)
	}

	basisSlope := (vertices[0].X() - vertices[2].X()) / (vertices[0].Y() - vertices[2].Y())
	basisX := vertices[0].X()

	for vi, vertex := range vertices[1:] {
		bottomLine := vertex.Y()
		slope := float32(0)

		// TODO: maybe fix the readability of this implementation
		if vertex.Y()-vertices[vi].Y() != 0 {
			slope = (vertex.X() - vertices[vi].X()) / (vertex.Y() - vertices[vi].Y())
		}
		x := vertices[vi].X()

		for currentLine >= bottomLine {
			if err := i.DrawLine(int(x+0.5), int(currentLine), int(basisX+0.5), int(currentLine), color); err != nil {
				return errors.Wrap(err, "draw.DrawFilledTriangle.lineSweep: failed to draw line")
			}

			basisX -= basisSlope
			x -= slope
			currentLine -= 1
		}
	}

	return nil
}

func sortTriangleVerticesByYDesc(vertices *[3]mgl32.Vec4) error {
	if vertices == nil {
		return errors.Errorf("sortTriangleVerticesByY: invalid nil vertices")
	}

	if vertices[0].Y() < vertices[1].Y() {
		vertices[0], vertices[1] = vertices[1], vertices[0]
	}

	if vertices[0].Y() < vertices[2].Y() {
		vertices[0], vertices[2] = vertices[2], vertices[0]
	}

	if vertices[1].Y() < vertices[2].Y() {
		vertices[1], vertices[2] = vertices[2], vertices[1]
	}

	return nil
}

func (i *ImgRGBA) boundingBox(vertices [3]mgl32.Vec4, color Color) error {
	return errors.Errorf("draw.DrawFilledTriangle.boundingBox: Not yet implemented")
}
