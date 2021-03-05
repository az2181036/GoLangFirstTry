package main

import (
	"../class"
	"../cloudedge"
	"../util"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var clusterings [][]class.Edge
	var data [][]int
	var t []float64
	dc := class.Datacenter{1000, make([]class.Task, 0,
		class.N_UPPER*class.TASK_NUMBER_UPPER_LIMIT[3]), 10.0 * 1e9}
	edges := cloud_edge.GenerateEdges()

	//clusterings := util.UseKmeans(edges)
	clusterings = util.UseClustering(edges)
	for o := 0; o < 4; o++ {
		for i := 0; i < len(edges); i++ {
			numberOfTasks := rand.Intn(class.TASK_NUMBER_UPPER_LIMIT[o])
			edges[i].TaskQueue = class.GenerateTaskQueue(numberOfTasks, edges[i].Id*class.TASK_NUMBER_UPPER_LIMIT[o])
			edges[i].MigQueue = []class.Task{}
			edges[i].ProcQueue = []class.Task{}
			edges[i].TimeLine = 0
		}

		for i := 0; i < len(edges); i++ {
			util.STF(&edges[i])
		}

		for i := 0; i < len(clusterings); i++ {
			for j := 0; j < len(clusterings[i]); j++ {
				clusterings[i][j] = edges[clusterings[i][j].Id]
			}
		}
		var cntTasksNumber = 0
		for i := 0; i < len(edges); i++ {
			cntTasksNumber += len(edges[i].ProcQueue)
		}

		fmt.Println(cntTasksNumber, len(edges))
		startime := time.Now().UnixNano()
		data = append(data, util.MyAlgorithm_v2(dc, edges, clusterings))
		endtime := time.Now().UnixNano()
		execTime := float64((endtime - startime) / 1e6)
		t = append(t, execTime)
	}
	util.WriteInCSV(data, t, "experiment_edf.csv")
}

//4961 121
//12934 121
//21433 121
//30081 121
//
//5459 121
//15072 121
//24703 121
//32603 121
