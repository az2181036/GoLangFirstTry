package cloud_edge

import (
	"../class"
	"sort"
)

func GetBipartiteGraphEdge(clustering []class.Edge, MigSide, RecSide []int) [][]int {
	var adjacencyList = make([][]int, 0, len(MigSide))
	for j := 0; j < len(MigSide); j++ {
		var tmpList = make([]int, 0, len(RecSide))
		for k := 0; k < len(RecSide); k++ {
			tmpTaskList := make([]class.Task, len(clustering[RecSide[k]].ProcQueue)+1)
			copy(tmpTaskList, clustering[RecSide[k]].ProcQueue)
			tmpTaskList = append(tmpTaskList, clustering[MigSide[j]].MigQueue[0])
			if JgMigration(tmpTaskList, clustering[RecSide[k]]) {
				tmpList = append(tmpList, k)
			}
		}
		adjacencyList = append(adjacencyList, tmpList)
	}
	return adjacencyList
}

func JgMigration(lst []class.Task, e class.Edge) bool {
	var tmp float32 = 0.0
	sort.Sort(class.NewTasks(lst))
	for i := 0; i < len(lst); i++ {
		tmp += class.TaskComputeTime(lst[i], e)
		if lst[i].DeadLine < tmp {
			return false
		}
	}
	return true
}