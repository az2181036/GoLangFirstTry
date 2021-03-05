package util

import (
	"../class"
	"../cloudedge"
)

func EdgeOnlyNoMigration(edges []class.Edge, o int) []int {
	cnt, cntDDLViolate := 0, 0
	for i := 0; i < len(edges); i++ {
		cnt += len(edges[i].TaskQueue)
		cntDDLViolate += len(edges[i].MigQueue)
	}

	return []int{class.TASK_NUMBER_UPPER_LIMIT[o], 0, 0,
		cnt, cntDDLViolate}
}

func EdgeOnlyWithMigration(edges []class.Edge, o int) []int {
	cnt, cntDDLViolate, cntMigrationToEdge, cntMigrationToCloud := 0, 0, 0, 0
	for i := 0; i < len(edges); i++ {
		cnt += len(edges[i].TaskQueue)
	}
	migSide, recSide := cloud_edge.GetInitBipartiteGraph(edges)
	for len(migSide) > 0 {
		vis := make([]int, len(migSide))
		adjacencyList := cloud_edge.GetBipartiteGraphEdge(edges, migSide, recSide)
		p := Hungary(adjacencyList)
		for k := 0; k < len(p); k++ {
			if p[k] != -1 {
				vis[p[k]] = 1
				edges[recSide[k]].ProcQueue = append(edges[recSide[k]].ProcQueue,
					edges[migSide[p[k]]].MigQueue[0])
				edges[recSide[k]].TimeLine += class.TaskComputeTime(edges[migSide[p[k]]].MigQueue[0],
					edges[recSide[k]])
				cntMigrationToEdge++
			}
		}
		for k := 0; k < len(vis); k++ {
			if vis[k] == 0 {
				cntDDLViolate++
			}
		}
		migSide, recSide = cloud_edge.GetBipartiteGraph(edges)
	}
	return []int{class.TASK_NUMBER_UPPER_LIMIT[o], cntMigrationToEdge, cntMigrationToCloud,
		cnt, cntDDLViolate}
}

func EdgeOnlyWithClusterings(edges []class.Edge, o int) []int {
	cnt, cntDDLViolate, cntMigrationToEdge, cntMigrationToCloud := 0, 0, 0, 0
	for i := 0; i < len(edges); i++ {
		cnt += len(edges[i].TaskQueue)
	}

	var MigSide, RecSide, vis []int
	clusterings := UseKmeans(edges)
	for i := 0; i < len(clusterings); i++ {
		MigSide, RecSide = cloud_edge.GetInitBipartiteGraph(clusterings[i])
		for len(MigSide) > 0 {
			vis = make([]int, len(MigSide))
			adjacencyList := cloud_edge.GetBipartiteGraphEdge(clusterings[i], MigSide, RecSide)
			p := Hungary(adjacencyList)
			for k := 0; k < len(p); k++ {
				if p[k] != -1 {
					vis[p[k]] = 1
					clusterings[i][RecSide[k]].ProcQueue = append(clusterings[i][RecSide[k]].ProcQueue,
						clusterings[i][MigSide[p[k]]].MigQueue[0])
					clusterings[i][RecSide[k]].TimeLine += class.TaskComputeTime(clusterings[i][MigSide[p[k]]].MigQueue[0],
						clusterings[i][RecSide[k]])
					cntMigrationToEdge++
				}
			}
			for k := 0; k < len(vis); k++ {
				if vis[k] == 0 {
					cntDDLViolate++
				}
			}
			MigSide, RecSide = cloud_edge.GetBipartiteGraph(clusterings[i])
		}
	}
	return []int{class.TASK_NUMBER_UPPER_LIMIT[o], cntMigrationToEdge, cntMigrationToCloud,
		cnt, cntDDLViolate}
}
