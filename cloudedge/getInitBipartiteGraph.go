package cloud_edge

import "../class"

func GetInitBipartiteGraph(clustering []class.Edge) ([]int, []int) {
	var MigSide, RecSide = make([]int, 0, len(clustering)), make([]int, 0, len(clustering))
	for j := 0; j < len(clustering); j++ {
		if len(clustering[j].MigQueue) > 0 {
			MigSide = append(MigSide, j)
		} else {
			RecSide = append(RecSide, j)
		}
	}
	return MigSide, RecSide
}
