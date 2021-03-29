package main

import (
	"../class"
	"../cloudedge"
	"../util"
	"sort"
	"time"
)

func main() {
	var clusterings [][]class.Edge
	var data [][]int
	var t []float64
	var load []float64
	var sigma []float64
	edges := cloud_edge.GenerateEdges()

	clusterings = util.UseClustering(edges)
	for o := 0; o < len(class.TASK_NUMBER_UPPER_LIMIT); o++ {
		dc := class.Datacenter{1000, make([]class.Task, 0,
			class.N_UPPER*class.TASK_NUMBER_UPPER_LIMIT[3]), 10.0 * 1e9}
		class.InitEdge(o, edges)

		for i := 0; i < len(edges); i++ {
			util.EDF(&edges[i])
			sort.Sort(class.DataTasks(edges[i].MigQueue))
		}

		for i := 0; i < len(clusterings); i++ {
			for j := 0; j < len(clusterings[i]); j++ {
				clusterings[i][j] = edges[clusterings[i][j].Id]
			}
		}

		startime := time.Now().UnixNano()
		data = append(data, util.MyAlgorithm_v2(&dc, clusterings))
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
	util.WriteInCSV(data, t, load, sigma,"experiment_edf_v2.csv")
}
