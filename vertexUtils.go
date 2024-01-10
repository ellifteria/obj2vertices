package main

import "math"

const MAX_COORDINATE_VALUE float64 = 0.999

type Vertex struct {
	X, Y, Z float64
}

func GetMaxCoordinateOfVertex(vertex Vertex) float64 {
	return math.Max(math.Max(math.Abs(vertex.X), math.Abs(vertex.Y)), math.Abs(vertex.Z))
}

func GetMaxCoordinateOfSlice(vertices []Vertex) float64 {
	max := 0.0

	for _, vertex := range vertices {
		vertexMax := GetMaxCoordinateOfVertex(vertex)
		if vertexMax > max {
			max = vertexMax
		}
	}

	return max
}

func NormalizeSlice(vertices []Vertex) []Vertex {
	normalizationConstant := MAX_COORDINATE_VALUE / GetMaxCoordinateOfSlice(vertices)

	var newVertices = make([]Vertex, len(vertices))

	for index, vertex := range vertices {
		newVertices[index] = Vertex{
			vertex.X * normalizationConstant,
			vertex.Y * normalizationConstant,
			vertex.Z * normalizationConstant,
		}
	}

	return newVertices
}
