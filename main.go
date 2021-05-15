package main

import (
	"fmt"
	image2 "github.com/gabriellukechen/tinyrenderergo/pkg/image"
	"github.com/gabriellukechen/tinyrenderergo/pkg/model"
	"github.com/go-gl/mathgl/mgl32"
	"image"
	"image/png"
	"os"
)

func main() {
	//defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()

	f, err := os.Create("output.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	width := 1000
	height := 1000
	img := image2.NewRGBAImage(image.Rect(0, 0, width, height))

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

			x0 := ((v0.X() + 1.) / 2.) * float32(width-1)
			y0 := ((v0.Y() + 1.) / 2.) * float32(height-1)
			x1 := ((v1.X() + 1.) / 2.) * float32(width-1)
			y1 := ((v1.Y() + 1.) / 2.) * float32(height-1)
			img.Line(int(x0), int(y0), int(x1), int(y1), image2.Color{255, 255, 255, 255})
		}
	}

	err = png.Encode(f, img)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
}
