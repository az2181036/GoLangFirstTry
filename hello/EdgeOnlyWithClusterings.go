package main

import (
	"../class"
	"../cloudedge"
	"../util"
	"math/rand"
	"time"
)

func main() {
	edges := cloud_edge.GenerateEdges()
	clusterings := util.GetClusterings(edges)

	var t []float64
	var data [][]int
	for o := 0; o < 4; o++ {
		for i := 0; i < len(edges); i++ {
			numberOfTasks := rand.Intn(class.TASK_NUMBER_UPPER_LIMIT[o])
			edges[i].TaskQueue = class.GenerateTaskQueue(numberOfTasks, edges[i].Id*class.TASK_NUMBER_UPPER_LIMIT[o])
		}
		startime := time.Now().UnixNano()
		for i := 0; i < len(edges); i++ {
			util.EDF(&edges[i])
		}
		data = append(data, util.EdgeOnlyWithClusterings(edges, clusterings))
		endtime := time.Now().UnixNano()
		execTime := float64((endtime - startime) / 1e9)
		t = append(t, execTime)
	}
	util.WriteInCSV(data, t, "experiment_EdgeOnlyWithClusterings.csv")
}
