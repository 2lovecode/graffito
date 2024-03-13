package graph

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
