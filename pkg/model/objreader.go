package model

import (
	"bufio"
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
	"os"
	"strconv"
	"strings"
)

// Notes on OBJ Files:
// Geometric vertices:
//		v (x, y, z, [,w]), default w = 1
// Texture coordinates:
//		vt (u, [,v ,w])  0 < u < 1
// Vector Normals (might not be unit vectors):
// 		vn (x, y, z)
// Parameter space vertices:
// 		vp (u, [,v] [,w])
// Polygonal face elements:
// 		f vertex_index/texture_index/normal_index (x3 for tri or more for quad)
// Line element
// 		l vertices

type ObjReader interface {
	ReadObjFile(model *Model) error
}

type objReader struct {
	file string
}

func NewObjReader(file string) ObjReader {
	return &objReader{
		file: file,
	}
}

type objLine struct {
	index int
	line string
}

func (obj *objReader) ReadObjFile(model *Model) error {
	objFile, err := os.Open(obj.file)
	if err != nil {
		return err
	}
	defer objFile.Close()

	scanner := bufio.NewScanner(objFile)

	for scanner.Scan() {
		l := scanner.Text()
		if l == "" {
			continue
		}

		args := strings.Fields(l)
		t := args[0]
		args = args[1:]

		switch t {
		case "v":
			v := mgl32.Vec4{}
			floats := make([]float32, 4)
			if err := obj.parseFloatArguments(&floats, args); err != nil {
				return err
			}
			v[0], v[1], v[2], v[3] = floats[0], floats[1], floats[2], floats[3]
			if floats[3] == 0 {
				v[3] = 1
			}
			model.Vertices = append(model.Vertices, v)
		case "vt":
			v := mgl32.Vec3{}
			floats := make([]float32, 3)
			if err := obj.parseFloatArguments(&floats, args); err != nil {
				return err
			}
			v[0], v[1], v[2] = floats[0], floats[1], floats[2]
			model.Textures = append(model.Textures, v)
		case "vn":
			v := mgl32.Vec3{}
			floats := make([]float32, 3)
			if err := obj.parseFloatArguments(&floats, args); err != nil {
				return err
			}
			v[0], v[1], v[2] = floats[0], floats[1], floats[2]
			model.Normals = append(model.Normals, v)
		case "vp":
			fmt.Println("obj vp not implemented")
		case "f":
			fe := FaceElement{
				Points: [3]FaceElementPoint{},
			}

			for i, point := range args {
				indices := strings.Split(point, "/")

				f := FaceElementPoint{}

				vi, _ := strconv.Atoi(indices[0])
				f.VertexIndex = &vi

				ti, err := strconv.Atoi(indices[1])
				if err != nil {
					f.TextureIndex = nil
				} else {
					f.TextureIndex = &ti
				}

				ni, err := strconv.Atoi(indices[2])
				if err != nil {
					f.NormalIndex = nil
				} else {
					f.NormalIndex = &ni
				}

				fe.Points[i] = f
			}

			model.Faces = append(model.Faces, fe)
		case "l":
			fmt.Println("obj l not implemented")
		default:
		}
	}

	return nil
}

func (obj *objReader) parseFloatArguments(floats *[]float32, args []string) error {
	for i := 1; i < len(args); i++ {
		f, err := strconv.ParseFloat(args[i], 32)
		if err != nil {
			return err
		}

		*floats = append(*floats, float32(f))
	}

	return nil
}