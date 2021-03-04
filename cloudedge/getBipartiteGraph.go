package cloud_edge

import (
	"../class"
)

func GetBipartiteGraph(edges []class.Edge, MigSide []int, RecSide []int) ([]int, []int) {
	for j := 0; j < len(MigSide); j++ {
		edges[MigSide[j]].MigQueue = append(edges[MigSide[j]].MigQueue[:0], edges[MigSide[j]].MigQueue[1:]...)
		if len(edges[MigSide[j]].MigQueue) == 0 {
			RecSide = append(RecSide, MigSide[j])
			MigSide = append(MigSide[:j], MigSide[j+1:]...)
		}
	}
	return MigSide, RecSide
}
