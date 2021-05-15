package model

import "github.com/go-gl/mathgl/mgl32"

type Model struct {
	Vertices []mgl32.Vec4
	Textures []mgl32.Vec3
	Normals  []mgl32.Vec3

	Faces []FaceElement
}

// a single face element some sort of polygon
type FaceElement struct {
	Points [3]FaceElementPoint
}

// a single set of vertex_index/texture_index/normal_index
// 		3 or more of these makes a face element
type FaceElementPoint struct {
	VertexIndex  *int
	TextureIndex *int
	NormalIndex  *int
}
