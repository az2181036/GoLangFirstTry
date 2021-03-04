package util

import (
	"../class"
	"../cloudedge"
	"sort"
)

type SecNewTasks []class.Task

func (t SecNewTasks) Len() int           { return len(t) }
func (t SecNewTasks) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t SecNewTasks) Less(i, j int) bool { return t[i].StartTime < t[j].StartTime }

func MyAlgorithm(dc class.Datacenter, edges []class.Edge, clusterings [][]int) []int {
	var MigSide, RecSide, vis []int
	var dcTaskQueue []class.Task
	cntMigrationToEdge := 0
	cntMigrationToCloud := 0
	for i := 0; i < len(clusterings); i++ {
		MigSide, RecSide = cloud_edge.GetInitBipartiteGraph(edges, clusterings[i])
		for len(MigSide) > 0 {
			vis = make([]int, len(MigSide))
			adjacencyList := cloud_edge.GetBipartiteGraphEdge(edges, MigSide, RecSide)
			p := Hungary(adjacencyList)
			for k := 0; k < len(p); k++ {
				if p[k] != -1 {
					vis[p[k]] = 1
					edges[RecSide[k]].ProcQueue = append(edges[RecSide[k]].ProcQueue,
						edges[MigSide[p[k]]].MigQueue[0])
					edges[RecSide[k]].TimeLine += class.TaskComputeTime(edges[MigSide[p[k]]].MigQueue[0],
						edges[RecSide[k]])
					cntMigrationToEdge++
				}
			}
			for k := 0; k < len(vis); k++ {
				if vis[k] == 0 {
					dcTaskQueue = append(dcTaskQueue, edges[MigSide[k]].MigQueue[0])
					cntMigrationToCloud++
				}
			}
			MigSide, RecSide = cloud_edge.GetBipartiteGraph(edges, MigSide, RecSide)
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
	for i := 0; i < len(edges); i++ {
		cnt += len(edges[i].TaskQueue)
	}
	return []int{cntMigrationToEdge, cntMigrationToCloud,
		cnt, cntDDLViolate}
}

func GetClusterings(edges []class.Edge) [][]int {
	return Kmeans(class.M, cloud_edge.N, edges)
}
