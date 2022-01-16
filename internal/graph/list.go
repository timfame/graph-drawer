package graph

import (
	"bufio"
	"fmt"
	"github.com/timfame/graph-drawer.git/pkg/utils"
	"os"
)

type list struct {
	g []edge
	n, m int

	edgeIterator int
}

type edge struct {
	from int
	to   int
}

func NewList() *list {
	return &list{edgeIterator: 0}
}

func (l *list) Read(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	n, err := utils.ReadInt(scanner)
	if err != nil {
		return fmt.Errorf("cannot read number of vertices: %w", err)
	}
	m, err := utils.ReadInt(scanner)
	if err != nil {
		return fmt.Errorf("cannot read number of edges: %w", err)
	}

	l.n = n
	l.m = m
	l.g = make([]edge, 0)
	for i := 0; i < m; i++ {
		a, err := utils.ReadInt(scanner)
		if err != nil {
			return fmt.Errorf("cannot read first vertice of edge: %w", err)
		}
		b, err := utils.ReadInt(scanner)
		if err != nil {
			return fmt.Errorf("cannot read second vertice of edge: %w", err)
		}
		a--; b--

		l.g = append(l.g, edge{
			from: a,
			to:   b,
		})
	}

	return nil
}

func (l *list) GetNumberOfVertices() int {
	return l.n
}

func (l *list) HasEdge() bool {
	if l.edgeIterator >= l.m {
		return false
	}
	return true
}

func (l *list) NextEdge() (int, int) {
	defer l.incIterator()
	if l.HasEdge() {
		return l.g[l.edgeIterator].from, l.g[l.edgeIterator].to
	}
	return 0, 0
}

func (l *list) incIterator() {
	l.edgeIterator++
	if l.edgeIterator > l.m {
		l.edgeIterator = l.m
	}
}
