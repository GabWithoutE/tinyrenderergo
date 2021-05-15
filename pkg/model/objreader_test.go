package model

import (
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestObjReader_ReadObjFile(t *testing.T) {
	cases := []struct{
		name string
		testFileContents string
		expected Model
		isErrorExpected bool
	}{
		{	"nothing special",
			"v 1 1 1\n v 2 2 2 2\n v 3 3 3\n vt 3 3 3\n vt 4 4 4\n vn 5 5 5\n vn 6 6 6\n f // // //\n",
			Model{Vertices: []mgl32.Vec4{mgl32.Vec4{1, 1, 1, 1}, mgl32.Vec4{2, 2, 2, 2}, mgl32.Vec4{3, 3, 3, 3}}, Textures: []mgl32.Vec3{mgl32.Vec3{3, 3, 3}, mgl32.Vec3{4, 4, 4}}, Normals: []mgl32.Vec3{mgl32.Vec3{5, 5, 5}, mgl32.Vec3{6, 6, 6}}, Faces: []FaceElement{FaceElement{[3]FaceElementPoint{FaceElementPoint{nil, nil, nil}, FaceElementPoint{nil, nil, nil}, FaceElementPoint{nil, nil, nil}}}}},
			false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			file, err := ioutil.TempFile(".", "test")
			if err != nil {
				t.Fatalf("failed to create test file\n")
			}
			defer os.Remove(file.Name())

			if err = ioutil.WriteFile(file.Name(), []byte(c.testFileContents), 0000); err != nil {
				t.Fatalf("failed to write to test file\n")
			}

			r := NewObjReader(file.Name())
			got := Model{
				Vertices: make([]mgl32.Vec4, 0),
				Normals: make([]mgl32.Vec3, 0),
				Textures: make([]mgl32.Vec3, 0),
			}

			err = r.ReadObjFile(&got)
			if err != nil {
				if !c.isErrorExpected {
					t.Errorf("Name: %v, Expected: no errors, Got: %+v", c.name, err)
				}
				return
			}

			fmt.Printf("%v", got.Vertices)

			// TODO: this is not working as a check...
			if reflect.DeepEqual(got, c.expected) {
				t.Fatalf("Name: %v, Expected: %v, Got: %v", c.name, c.expected, got)
			}

			t.Errorf("")
		})
	}
}
