package main

import (
	"fmt"
	"github.com/gabriellukechen/tinyrenderergo/pkg/draw"
	"github.com/gabriellukechen/tinyrenderergo/pkg/model"
	"github.com/go-gl/mathgl/mgl32"
	"image"
	"image/png"
	"os"
)

func main() {
	f, err := os.Create("output.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	width := 1000
	height := 1000

	widestpix := width - 1
	highestpix := height - 1

	img := draw.NewRGBAImage(image.Rect(0, 0, width, height))

	m := model.Model{
		Vertices: make([]mgl32.Vec4, 0),
		Normals:  make([]mgl32.Vec3, 0),
		Textures: make([]mgl32.Vec3, 0),
	}

	r := model.NewObjReader("assets/head.obj")

	if err = r.ReadObjFile(&m); err != nil {
		fmt.Printf("failed to read obj file %+v", err)
	}

	for _, face := range m.Faces {
		for i, point := range face.Points {
			v0 := m.Vertices[*point.VertexIndex]
			v1 := m.Vertices[*face.Points[(i+1)%3].VertexIndex]

			x0 := ((v0.X() + 1.) / 2.) * float32(widestpix)
			y0 := ((v0.Y() + 1.) / 2.) * float32(highestpix)
			x1 := ((v1.X() + 1.) / 2.) * float32(widestpix)
			y1 := ((v1.Y() + 1.) / 2.) * float32(highestpix)
			if err := img.DrawLine(int(x0), int(y0), int(x1), int(y1), draw.Color{255, 255, 255, 255}); err != nil {
				fmt.Printf("%+v\n", err)
			}
		}
	}

	//if err := img.DrawFilledTriangle(
	//	draw.LineSweep,
	//	[3]mgl32.Vec4{mgl32.Vec4{0, 0, 0, 0}, mgl32.Vec4{999, 0, 0, 0}, mgl32.Vec4{500, 999, 0, 0}},
	//	draw.Color{255, 255, 255, 255},
	//); err != nil {
	//	fmt.Printf("failed to draw triangle, %+v\n", err)
	//}

	if err = png.Encode(f, img); err != nil {
		fmt.Printf("%+v\n", err)
	}
}
