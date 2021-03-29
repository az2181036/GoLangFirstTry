package main

import (
	"../class"
	"../cloudedge"
	"../util"
	"fmt"
	"time"
)

func main() {
	var clusterings [][]class.Edge
	var data [][]int
	var t []float64
	var load []float64
	var sigma []float64
	dc := class.Datacenter{1000, make([]class.Task, 0,
		class.N_UPPER*class.TASK_NUMBER_UPPER_LIMIT[3]), 10.0 * 1e9}
	edges := cloud_edge.GenerateEdges()

	clusterings = util.UseClustering(edges)
	for o := 0; o < len(class.TASK_NUMBER_UPPER_LIMIT); o++ {
		class.InitEdge(o, edges)

		for i := 0; i < len(edges); i++ {
			util.EDF(&edges[i])
			// sort.Sort(class.DataTasks(edges[i].MigQueue))
		}

		for i := 0; i < len(clusterings); i++ {
			for j := 0; j < len(clusterings[i]); j++ {
				clusterings[i][j] = edges[clusterings[i][j].Id]
			}
		}

		startime := time.Now().UnixNano()
		data = append(data, util.MyAlgorithm(dc, edges, clusterings))
		endtime := time.Now().UnixNano()
		execTime := float64((endtime - startime) / 1e6)
		t = append(t, execTime)

		tasksNumberComputeOnCloud := make([]int, len(edges))
		for i:=0; i<len(dc.TaskQueue); i++{
			edge_id := dc.TaskQueue[i].Id / class.TASK_NUMBER_UPPER_LIMIT[o]
			tasksNumberComputeOnCloud[edge_id]++
		}

		cnt := 0.0
		cnt_sigma := 0.0
		for i := 0; i < len(clusterings); i++ {
			for j := 0; j < len(clusterings[i]); j++ {
				cnt += float64(class.ComputeLoad(clusterings[i][j]))
				if float64(len(clusterings[i][j].TaskQueue)) == 0 {
					cnt_sigma += 1
				} else {
					cnt_sigma += float64(clusterings[i][j].ComputeTaskNumber +
						tasksNumberComputeOnCloud[clusterings[i][j].Id]) / float64(len(clusterings[i][j].TaskQueue))
				}
			}
		}
		load = append(load, cnt/float64(len(edges)))
		sigma = append(sigma, cnt_sigma/float64(len(edges)))
	}
	fmt.Println(load)
	util.WriteInCSV(data, t, load, sigma, "experiment_edf.csv")
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

// [0.29731497883966024 0.746139357312898 0.9655610773681609 0.9673220499241648]
// [0.2993681731110454 0.758882144758524 0.9654630001418847 0.968332796426844]