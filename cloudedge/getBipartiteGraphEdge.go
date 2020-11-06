package cloud_edge

import "github.com/user/class"

func getBipartiteGraphEdge(MigSide, RecSide []class.Edge) [][]int{
	var adjacencyList = make([][]int, 0, len(MigSide))
	for i:=0; i<len(MigSide); i++ {
		var tmpList = make([]int, 0, len(RecSide))
		for j:=0; j<len(RecSide); j++ {
			if len(MigSide[i].MigQueue) != 0 {
				if MigSide[i].MigQueue[0].DeadLine >= RecSide[j].TimeLine {
					tmpList = append(tmpList, j)
				}
			} else {
				break
			}
		}
		adjacencyList = append(adjacencyList, tmpList)
	}
	return adjacencyList
}