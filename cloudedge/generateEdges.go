package cloud_edge

import (
	"../class"
	"math/rand"
)

var N = rand.Intn(class.N_UPPER-class.N_LOWER) + class.N_LOWER

func GenerateEdges() []class.Edge {
	var edges = make([]class.Edge, 0, N)

	// N edges
	lxs := make([]int, 0, N)
	lys := make([]int, 0, N)

	// get the position of each edge from area [100x100] randomly.
	for i := 0; i < N; i++ {
		lxs = append(lxs, rand.Intn(100))
		lys = append(lys, rand.Intn(100))
	}

	for i := 0; i < N; i++ {
		tmpEdge := class.GenerateEdge(lxs[i], lys[i], i)
		edges = append(edges, tmpEdge)
	}
	return edges
}
