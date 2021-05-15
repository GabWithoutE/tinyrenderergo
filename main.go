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

	img := image2.NewRGBAImage(image.Rect(0, 0, 100, 100))
	img.Line(13, 20, 80, 40, image2.Color{255, 255, 255, 255})
	img.Line(20, 13, 40, 80, image2.Color{0, 0, 255, 255})
	img.Line(80, 40, 13, 20, image2.Color{255, 0, 0, 255})

	m := model.Model{
		Vertices: make([]mgl32.Vec4, 0),
		Normals: make([]mgl32.Vec3, 0),
		Textures: make([]mgl32.Vec3, 0),
	}

	r := model.NewObjReader("assets/head.obj")
	r.ReadObjFile(&m)

	err = png.Encode(f, img)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
}