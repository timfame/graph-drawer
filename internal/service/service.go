package service

import (
	"github.com/timfame/graph-drawer.git/internal/draw"
	"github.com/timfame/graph-drawer.git/internal/graph"
	"math"
	"strconv"
)

const (
	width        = 1000
	height       = 1000
	radius       = 300.0
	vertexRadius = 20.0
)

type service struct {
	visualiser draw.Drawer
	graph      graph.Graph
}

func New(v draw.Drawer, g graph.Graph) *service {
	return &service{
		visualiser: v,
		graph:      g,
	}
}

func (s *service) Draw(input, output string) error {
	if err := s.graph.Read(input); err != nil {
		return err
	}

	s.visualiser.NewImage(width, height)

	n := s.graph.GetNumberOfVertices()
	for i := 0; i < n; i++ {
		if err := s.drawVertex(i, n); err != nil {
			return err
		}
	}

	for s.graph.HasEdge() {
		from, to := s.graph.NextEdge()
		if err := s.drawEdge(from, to, n); err != nil {
			return err
		}
	}

	return s.visualiser.SaveImage(output)
}

func (s *service) drawVertex(index, total int) error {
	x, y := getVertexCoords(index, total)
	if err := s.visualiser.DrawCircle(x, y, vertexRadius); err != nil {
		return err
	}
	return s.visualiser.DrawText(x, y, strconv.Itoa(index + 1))
}

func (s *service) drawEdge(from, to, total int) error {
	x1, y1 := getVertexCoords(from, total)
	x2, y2 := getVertexCoords(to, total)
	return s.visualiser.DrawLine(x1, y1, x2, y2)
}

func getVertexCoords(index, total int) (int, int) {
	x := width / 2.0 + radius * math.Cos(math.Pi * 2 * float64(index) / float64(total))
	y := height / 2.0 + radius * math.Sin(math.Pi * 2 * float64(index) / float64(total))
	return int(x), int(y)
}
