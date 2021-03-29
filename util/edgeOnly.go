package util

import (
	"../class"
	"../cloudedge"
	"sort"
)

func EdgeOnlyNoMigration(edges []class.Edge) []int {
	cnt, cntDDLViolate := 0, 0
	for i := 0; i < len(edges); i++ {
		cnt += len(edges[i].TaskQueue)
		cntDDLViolate += len(edges[i].MigQueue)
	}

	return []int{0, 0, cnt, cntDDLViolate}
}

func EdgeOnlyWithMigration(edges []class.Edge) []int {
	cnt, cntDDLViolate, cntMigrationToEdge, cntMigrationToCloud := 0, 0, 0, 0
	for i := 0; i < len(edges); i++ {
		cnt += len(edges[i].TaskQueue)
	}
	migSide, recSide := cloud_edge.GetInitBipartiteGraph(edges)
	for len(migSide) > 0 {
		vis := make([]int, len(migSide))
		adjacencyList := cloud_edge.GetBipartiteGraphEdge(edges, migSide, recSide)
		p := Hungary(adjacencyList)
		for k := 0; k < len(recSide); k++ {
			if p[k] != -1 {
				vis[p[k]] = 1
				migrateATask(&edges[migSide[p[k]]], &edges[recSide[k]])
				cntMigrationToEdge++
			}
		}
		for k := 0; k < len(vis); k++ {
			if vis[k] == 0 {
				cntDDLViolate++
				edges[migSide[k]].MigQueue = append(edges[migSide[k]].MigQueue[:0],
					edges[migSide[k]].MigQueue[1:]...)
			}
		}
		migSide, recSide = cloud_edge.GetBipartiteGraph(edges)
	}
	return []int{cntMigrationToEdge, cntMigrationToCloud,
		cnt, cntDDLViolate}
}

func EdgeCloudWithMigration(dc class.Datacenter, edges []class.Edge) []int {
	var dcTaskQueue []class.Task
	cnt, cntDDLViolate, cntMigrationToEdge, cntMigrationToCloud := 0, 0, 0, 0
	for i := 0; i < len(edges); i++ {
		cnt += len(edges[i].TaskQueue)
	}
	migSide, recSide := cloud_edge.GetInitBipartiteGraph(edges)
	for len(migSide) > 0 {
		vis := make([]int, len(migSide))
		adjacencyList := cloud_edge.GetBipartiteGraphEdge(edges, migSide, recSide)
		p := Hungary(adjacencyList)
		for k := 0; k < len(recSide); k++ {
			if p[k] != -1 {
				vis[p[k]] = 1
				migrateATask(&edges[migSide[p[k]]], &edges[recSide[k]])
				cntMigrationToEdge++
			}
		}
		for k := 0; k < len(vis); k++ {
			if vis[k] == 0 {
				dcTaskQueue = append(dcTaskQueue, edges[migSide[k]].MigQueue[0])
				edges[migSide[k]].MigQueue = append(edges[migSide[k]].MigQueue[:0],
					edges[migSide[k]].MigQueue[1:]...)
				cntMigrationToCloud++
			}
		}
		migSide, recSide = cloud_edge.GetBipartiteGraph(edges)
	}

	var timeLine float32 = 0
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

	return []int{cntMigrationToEdge, cntMigrationToCloud,
		cnt, cntDDLViolate}
}


func EdgeOnlyWithClusterings(edges []class.Edge, clusterings [][]class.Edge) []int {
	cnt, cntDDLViolate, cntMigrationToEdge, cntMigrationToCloud := 0, 0, 0, 0
	for i := 0; i < len(edges); i++ {
		cnt += len(edges[i].TaskQueue)
	}

	var MigSide, RecSide, vis []int
	for i := 0; i < len(clusterings); i++ {
		MigSide, RecSide = cloud_edge.GetInitBipartiteGraph(clusterings[i])
		for len(MigSide) > 0 {
			vis = make([]int, len(MigSide))
			adjacencyList := cloud_edge.GetBipartiteGraphEdge(clusterings[i], MigSide, RecSide)
			p := Hungary(adjacencyList)
			for k := 0; k < len(RecSide); k++ {
				if p[k] != -1 {
					mig_idx := p[k]
					vis[mig_idx] = 1
					migrateATask(&clusterings[i][MigSide[mig_idx]], &clusterings[i][RecSide[k]])
					cntMigrationToEdge++
				}
			}
			for k := 0; k < len(vis); k++ {
				if vis[k] == 0 {
					clusterings[i][MigSide[k]].MigQueue = append(clusterings[i][MigSide[k]].MigQueue[:0],
						clusterings[i][MigSide[k]].MigQueue[1:]...)
					cntDDLViolate++
				}
			}
			MigSide, RecSide = cloud_edge.GetBipartiteGraph(clusterings[i])
		}
	}
	return []int{cntMigrationToEdge, cntMigrationToCloud, cnt, cntDDLViolate}
}
