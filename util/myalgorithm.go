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

func MyAlgorithm(dc class.Datacenter, edges []class.Edge, clusterings [][]class.Edge) []int {
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
			for k := 0; k < len(p); k++ {
				if p[k] != -1 {
					mig_idx := p[k]
					vis[mig_idx] = 1
					migrateATask(&clusterings[i][MigSide[mig_idx]], &clusterings[i][RecSide[k]])
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
	return []int{cntMigrationToEdge, cntMigrationToCloud, cnt, cntDDLViolate}
}

func UseKmeans(edges []class.Edge) [][]class.Edge {
	return Kmeans(class.M, cloud_edge.N, edges)
}

func UseClustering(edges []class.Edge) [][]class.Edge {
	rst, _ := Clustering(class.M, edges)
	return rst
}

func migrateATask(src *class.Edge, des *class.Edge) {
	src.ComputeTaskNumber ++
	des.ProcQueue = append(des.ProcQueue, src.MigQueue[0])
	des.TimeLine += class.TaskComputeTime(src.MigQueue[0], *des)
	src.MigQueue = append(src.MigQueue[:0], src.MigQueue[1:]...)
}
