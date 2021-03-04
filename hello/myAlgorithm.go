package main

import (
	"../class"
	"../cloudedge"
	"../util"
	"math/rand"
	"time"
)

func main() {
	var data [][]int
	var t []float64
	dc := class.Datacenter{1000, make([]class.Task, 0,
		class.N_UPPER*class.TASK_NUMBER_UPPER_LIMIT[3]), 10.0 * 1e9}
	edges := cloud_edge.GenerateEdges()
	for o := 0; o < 4; o++ {
		for i := 0; i < len(edges); i++ {
			numberOfTasks := rand.Intn(class.TASK_NUMBER_UPPER_LIMIT[o])
			edges[i].TaskQueue = class.GenerateTaskQueue(numberOfTasks, edges[i].Id*class.TASK_NUMBER_UPPER_LIMIT[o])
		}
		clusterings := util.GetClusterings(edges)
		startime := time.Now().UnixNano()
		for i := 0; i < len(edges); i++ {
			util.EDF(&edges[i])
		}
		data = append(data, util.MyAlgorithm(dc, edges, clusterings))
		endtime := time.Now().UnixNano()
		execTime := float64((endtime - startime) / 1e9)
		t = append(t, execTime)
	}
	util.WriteInCSV(data, t, "experiment_my_150_max_edges_number_12_clusterings.csv")
}
