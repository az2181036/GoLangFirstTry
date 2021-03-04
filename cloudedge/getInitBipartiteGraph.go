package cloud_edge

import "../class"

func GetInitBipartiteGraph(edges []class.Edge, clustering []int) ([]int, []int) {
	var MigSide, RecSide = make([]int, 0, len(clustering)), make([]int, 0, len(clustering))
	for j := 0; j < len(clustering); j++ {
		if len(edges[clustering[j]].MigQueue) > 0 {
			MigSide = append(MigSide, clustering[j])
		} else {
			RecSide = append(RecSide, clustering[j])
		}
	}
	return MigSide, RecSide
}
