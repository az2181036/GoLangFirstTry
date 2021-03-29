package main

import (
	"../class"
	"../cloudedge"
	"../util"
	"time"
)

func main() {
	var t, load, sigma []float64
	var data [][]int
	edges := cloud_edge.GenerateEdges()
	for o := 0; o < len(class.TASK_NUMBER_UPPER_LIMIT); o++ {
		dc := class.Datacenter{1000, make([]class.Task, 0,
			class.N_UPPER*class.TASK_NUMBER_UPPER_LIMIT[3]), 10.0 * 1e9}
		class.InitEdge(o, edges)
		for i := 0; i < len(edges); i++ {
			util.EDF(&edges[i])
		}
		startime := time.Now().UnixNano()
		data = append(data, util.EdgeCloudWithMigration(dc, edges))
		endtime := time.Now().UnixNano()
		execTime := float64((endtime - startime) / 1e6)
		t = append(t, execTime)

		tasksNumberComputeOnCloud := make([]int, len(edges))
		for i := 0; i < len(dc.TaskQueue); i++ {
			edge_id := dc.TaskQueue[i].Id / class.TASK_NUMBER_UPPER_LIMIT[o]
			tasksNumberComputeOnCloud[edge_id]++
		}

		cnt := 0.0
		cnt_sigma := 0.0
		for i := 0; i < len(edges); i++ {
			cnt += float64(class.ComputeLoad(edges[i]))
			if float64(len(edges[i].TaskQueue)) == 0 {
				cnt_sigma += 1
			} else {
				cnt_sigma += float64(edges[i].ComputeTaskNumber+
					tasksNumberComputeOnCloud[edges[i].Id]) / float64(len(edges[i].TaskQueue))
			}
		}
		load = append(load, cnt/float64(len(edges)))
		sigma = append(sigma, cnt_sigma/float64(len(edges)))
	}
	util.WriteInCSV(data, t, load, sigma, "experiment_EdgeCloudWithMigration_800.csv")
}
