package util

import (
	"../class"
	"../cloudedge"
	"sort"
)

func MyAlgorithm_v2(dc class.Datacenter, edges []class.Edge, clusterings [][]class.Edge) []int {
	var MigSide, RecSide, vis []int
	var dcTaskQueue []class.Task
	cntMigrationToEdge := 0
	cntMigrationToCloud := 0
	for i := 0; i < len(clusterings); i++ {
		MigSide, RecSide = cloud_edge.GetInitBipartiteGraph(clusterings[i])
		for len(MigSide) > 0 {
			vis = make([]int, len(MigSide))
			adjacencyList := cloud_edge.GetBipartiteGraphEdge(clusterings[i], MigSide, RecSide)
			p := Hungary(adjacencyList)
			for k := 0; k < len(RecSide); k++ {
				if p[k] != -1 {
					vis[p[k]] = 1
					migrateATask(&clusterings[i][MigSide[p[k]]], &clusterings[i][RecSide[k]])
					for ; len(clusterings[i][MigSide[p[k]]].MigQueue) > 0; cntMigrationToEdge++ {
						tmpTaskQueue := clusterings[i][RecSide[k]].ProcQueue
						tmpTaskQueue = append(tmpTaskQueue, clusterings[i][MigSide[p[k]]].MigQueue[0])
						if !cloud_edge.JgMigration(tmpTaskQueue, clusterings[i][RecSide[k]]) {
							break
						}
						migrateATask(&clusterings[i][MigSide[p[k]]], &clusterings[i][RecSide[k]])
					}
					cntMigrationToEdge++
				}
			}
			for k := 0; k < len(vis); k++ {
				if vis[k] == 0 {
					dcTaskQueue = append(dcTaskQueue, clusterings[i][MigSide[k]].MigQueue[0])
					clusterings[i][MigSide[k]].MigQueue = append(clusterings[i][MigSide[k]].MigQueue[:0],
						clusterings[i][MigSide[k]].MigQueue[1:]...)
					cntMigrationToCloud++
				}
			}
			MigSide, RecSide = cloud_edge.GetBipartiteGraph(clusterings[i])
		}
	}

	for i := 0; i < len(dcTaskQueue); i++ {
		trans_time := float32(float32(dcTaskQueue[i].Data) / float32(dc.Bandwidth*1e6*8))
		dcTaskQueue[i].StartTime = trans_time
	}

	var timeLine float32 = 0
	var cntDDLViolate int = 0
	sort.Sort(SecNewTasks(dcTaskQueue))
	for i := 0; i < len(dcTaskQueue); i++ {
		tmp := cloud_edge.CloudFinishTime(timeLine, dcTaskQueue[0], dc)
		if tmp <= dcTaskQueue[i].DeadLine {
			dc.TaskQueue = append(dc.TaskQueue, dcTaskQueue[i])
			timeLine = tmp
		} else {
			cntDDLViolate++
		}
	}

	cnt := 0
	for i := 0; i < len(clusterings); i++ {
		for j := 0; j < len(clusterings[i]); j++ {
			cnt += len(clusterings[i][j].TaskQueue)
		}
	}
	return []int{cntMigrationToEdge, cntMigrationToCloud,
		cnt, cntDDLViolate}
}
