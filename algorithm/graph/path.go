package graph

import "errors"

type ShortestPathMethod string

const (
	Dijkstra ShortestPathMethod = "dijkstra"
)

func ShortestPath(g *Graph, from string, to string, method ShortestPathMethod) (path string, err error) {
	switch method {
	case Dijkstra:
		path, err = BidirectionalDijkstra(g, from, to)

	}
	return
}

func BidirectionalDijkstra(g *Graph, from string, to string) (path string, err error) {
	if g == nil {
		err = errors.New("错误")
		return
	}

	if g.GetVertex(from) == nil {
		err = errors.New("错误")
		return
	}

	if g.GetVertex(to) == nil {
		err = errors.New("错误")
		return
	}
	
	return
}
