package graph

import "errors"

// Dijkstra算法
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
