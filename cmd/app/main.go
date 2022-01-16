package main

import (
	"flag"
	"github.com/timfame/graph-drawer.git/internal/draw"
	"github.com/timfame/graph-drawer.git/internal/graph"
	"github.com/timfame/graph-drawer.git/internal/service"
	"log"
	"os"
)

func main() {
	visualiserFlag := flag.String("v", "lib", "Visualisation method: [lib|std]")
	graphFlag := flag.String("g", "matrix", "Graph representation: [matrix|list]")
	input := flag.String("i", "input.txt", "Path to file with graph input")
	output := flag.String("o", "result.png", "Path to file with result of drawing in png format")
	flag.Parse()

	var (
		v draw.Drawer
		g graph.Graph
	)

	if *visualiserFlag == "lib" {
		v = draw.NewGraphicsLib()
	} else if *visualiserFlag == "std" {
		v = draw.NewSTD()
	} else {
		flag.Usage()
		os.Exit(2)
	}

	if *graphFlag == "matrix" {
		g = graph.NewMatrix()
	} else if *graphFlag == "list" {
		g = graph.NewList()
	} else {
		flag.Usage()
		os.Exit(2)
	}

	var s service.Service = service.New(v, g)

	if err := s.Draw(*input, *output); err != nil {
		log.Fatalln("Draw graph error:", err)
	}

	log.Println("Graph was drawn")
}
