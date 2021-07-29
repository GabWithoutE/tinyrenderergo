package main

import (
	"fmt"
	"github.com/gabriellukechen/tinyrenderergo/pkg/draw"
	"github.com/gabriellukechen/tinyrenderergo/pkg/model"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/pkg/errors"
	"image"
	"image/png"
	"math/rand"
	"os"
)

func main() {
	width := 1000
	height := 1000

	m := model.Model{
		Vertices: make([]mgl32.Vec4, 0),
		Normals:  make([]mgl32.Vec3, 0),
		Textures: make([]mgl32.Vec3, 0),
	}

	r := model.NewObjReader("assets/head.obj")

	if err := r.ReadObjFile(&m); err != nil {
		fmt.Printf("failed to read obj file %+v", err)
	}

	f, err := os.Create("wire-frame-output.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// wireframe
	if err = Render(&m, "output-wireframe.png", width, height, WireFrameRender()); err != nil {
		fmt.Printf("Error Main Wireframe: %v\n", err)
	}

	// filled triangles line-sweep
	if err = Render(&m, "output-line-sweep-random-color.png", width, height, FilledPolygonRandomColorRender()); err != nil {
		fmt.Printf("Error Main FilledTriangleLineSweep: %v\n", err)
	}
}

func WireFrameRender() RenderFunction {
	return func(img *draw.ImgRGBA, model *model.Model) error {
		for _, face := range model.Faces {
			for i, point := range face.Points {
				v0 := model.Vertices[*point.VertexIndex]
				v1 := model.Vertices[*face.Points[(i+1)%3].VertexIndex]

				x0 := ((v0.X() + 1.) / 2.) * float32(img.Rect.Max.X-1)
				y0 := ((v0.Y() + 1.) / 2.) * float32(img.Rect.Max.Y-1)
				x1 := ((v1.X() + 1.) / 2.) * float32(img.Rect.Max.X-1)
				y1 := ((v1.Y() + 1.) / 2.) * float32(img.Rect.Max.Y-1)
				if err := img.DrawLine(int(x0), int(y0), int(x1), int(y1), draw.Color{255, 255, 255, 255}); err != nil {
					return errors.Wrapf(err, "WireFrameRender() failed to draw line")
				}
			}
		}

		return nil
	}
}

func FilledPolygonRandomColorRender() RenderFunction {
	return func(img *draw.ImgRGBA, model *model.Model) error {
		for _, face := range model.Faces {
			color := draw.Color{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255}

			triangle := [3]mgl32.Vec4{}
			for i, point := range face.Points {
				v0 := model.Vertices[*point.VertexIndex]
				x0 := ((v0.X() + 1.) / 2.) * float32(img.Rect.Max.X-1)
				y0 := ((v0.Y() + 1.) / 2.) * float32(img.Rect.Max.Y-1)
				triangle[i] = mgl32.Vec4{x0, y0, 0, 0}
			}

			if err := img.DrawFilledTriangle(draw.LineSweep, triangle, color); err != nil {
				fmt.Printf("FilledPolygonRandomColorRender() failed to draw triangle: %v\n", err)
			}
		}

		return nil
	}
}

func Render(model *model.Model, fileName string, width int, height int, renderMethod RenderFunction) error {
	f, err := os.Create(fileName)
	if err != nil {
		return errors.Wrapf(err, "main.Render Failed to open file: %v\n", fileName)
	}
	defer f.Close()

	img := draw.NewRGBAImage(image.Rect(0, 0, width, height))

	if err = renderMethod(img, model); err != nil {
		return errors.Wrapf(err, "main.Render Failed to run render function: %v\n", renderMethod)
	}

	if err = png.Encode(f, img); err != nil {
		return errors.Wrapf(err, "main.Render Failed to Encode: %v\n", fileName)
	}

	return nil
}

type RenderFunction func(img *draw.ImgRGBA, model *model.Model) error
