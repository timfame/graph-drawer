package graph

type Graph interface {
	Read(path string) error
	GetNumberOfVertices() int
	HasEdge() bool
	NextEdge() (int, int)
}
