package cloud_edge

import (
	"fmt"
	"github.com/user/class"
	"github.com/user/util"
	"math/rand"
)

func GetClusterings() [][]class.Edge {
	var N = rand.Intn(class.N_UPPER-class.N_LOWER) + class.N_LOWER
	var edges = make([]class.Edge, 0, N)
	var clusterings [][]class.Edge
	var clusterCenter []class.Edge

	// N edges
	lxs := make([]int, 0, N)
	lys := make([]int, 0, N)

	// get the position of each edge from area [100x100] randomly.
	for i := 0; i < N; i++ {
		lxs = append(lxs, rand.Intn(100))
		lys = append(lys, rand.Intn(100))
	}

	for i := 0; i < N; i++ {
		tmpEdge := class.GenerateEdge(lxs[i], lys[i])
		edges = append(edges, tmpEdge)
	}
	clusterings = util.Kmeans(class.M,N,clusterCenter,edges, clusterings)
	for i := 0; i < class.M; i++ {
		fmt.Println(clusterings[i])
	}
	return clusterings
}