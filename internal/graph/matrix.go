package graph

import (
	"bufio"
	"fmt"
	"github.com/timfame/graph-drawer.git/pkg/utils"
	"os"
)

type matrix struct {
	g [][]int
	n int

	iteratorFromVertex int
	iteratorToVertex   int
}

func NewMatrix() *matrix {
	return &matrix{
		iteratorFromVertex: 0,
		iteratorToVertex:   0,
	}
}

func (m *matrix) Read(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	n, err := utils.ReadInt(scanner)
	if err != nil {
		return fmt.Errorf("cannot read number of edges: %w", err)
	}

	m.n = n
	m.g = make([][]int, n)
	for i := 0; i < n; i++ {
		m.g[i] = make([]int, n)
		for j := 0; j < n; j++ {
			v, err := utils.ReadInt(scanner)
			if err != nil {
				return fmt.Errorf("cannot read matrix graph element: %w", err)
			}
			m.g[i][j] = v
		}
	}

	return nil
}

func (m *matrix) GetNumberOfVertices() int {
	return m.n
}

func (m *matrix) HasEdge() bool {
	if m.iteratorFromVertex >= m.n {
		return false
	}
	return true
}

func (m *matrix) NextEdge() (int, int) {
	defer m.incIterator()
	for m.HasEdge() {
		if m.g[m.iteratorFromVertex][m.iteratorToVertex] == 1 {
			break
		}
		m.incIterator()
	}
	return m.iteratorFromVertex, m.iteratorToVertex
}

func (m *matrix) incIterator() {
	m.iteratorToVertex++
	if m.iteratorToVertex >= m.n {
		m.iteratorFromVertex++
		m.iteratorToVertex = 0
	}
	if m.iteratorFromVertex > m.n {
		m.iteratorFromVertex = m.n
	}
}
