package graph

type Node struct {
	ID   uint
	Next *Node
}

type Vertex struct {
	id  string
	adj map[*Vertex]int
}

func NewVertex(id string) *Vertex {
	return &Vertex{
		id:  id,
		adj: make(map[*Vertex]int),
	}
}

func (v *Vertex) ID() string {
	return v.id
}

func (v *Vertex) Weight(adj *Vertex) (weight int) {
	if _, ok := v.adj[adj]; ok {
		weight = v.adj[adj]
	}
	return
}

func (v *Vertex) AddAdj(adj *Vertex, weight int) {
	v.adj[adj] = weight
}

type Graph struct {
	vertices    map[string]*Vertex
	verticesLen int
}

func NewGraph() *Graph {
	return &Graph{
		verticesLen: 0,
		vertices:    make(map[string]*Vertex),
	}
}

func (g *Graph) AddVertex(id string) {
	if _, ok := g.vertices[id]; !ok {
		g.verticesLen++
		g.vertices[id] = NewVertex(id)
	}
}

func (g *Graph) GetVertex(id string) *Vertex {
	if _, ok := g.vertices[id]; ok {
		return g.vertices[id]
	}
	return nil
}

func (g *Graph) AddEdge(from string, to string, weight int) {
	if _, ok := g.vertices[from]; !ok {
		g.AddVertex(from)
	}

	if _, ok := g.vertices[to]; !ok {
		g.AddVertex(to)
	}

	g.vertices[from].AddAdj(g.vertices[to], weight)
}
