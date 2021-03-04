package cloud_edge

import (
	"../class"
)

func GetBipartiteGraph(clustering []class.Edge) ([]int, []int) {
	var MigSide, RecSide []int
	for j := 0; j < len(clustering); j++ {
		if len(clustering[j].MigQueue) > 0 {
			clustering[j].MigQueue = append(clustering[j].MigQueue[:0], clustering[j].MigQueue[1:]...)
			if len(clustering[j].MigQueue) > 0 {
				MigSide = append(MigSide, j)
			}
		} else {
			RecSide = append(RecSide, j)
		}
	}
	return MigSide, RecSide
}
