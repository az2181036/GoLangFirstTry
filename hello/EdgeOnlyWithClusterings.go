package main

import (
	"../class"
	"../cloudedge"
	"../util"
	"fmt"
	"time"
)

func main() {
	var t,load,sigma []float64
	var data [][]int
	edges := cloud_edge.GenerateEdges()
	clusterings := util.UseClustering(edges)
	for o := 0; o < len(class.TASK_NUMBER_UPPER_LIMIT); o++ {
		class.InitEdge(o, edges)

		for i := 0; i < len(edges); i++ {
			util.EDF(&edges[i])
		}

		for i := 0; i < len(clusterings); i++ {
			for j := 0; j < len(clusterings[i]); j++ {
				clusterings[i][j] = edges[clusterings[i][j].Id]
			}
		}

		startime := time.Now().UnixNano()
		data = append(data, util.EdgeOnlyWithClusterings(edges, clusterings))
		endtime := time.Now().UnixNano()
		execTime := float64((endtime - startime) / 1e6)
		t = append(t, execTime)

		cnt := 0.0
		cnt_sigma := 0.0
		for i := 0; i < len(clusterings); i++ {
			for j := 0; j < len(clusterings[i]); j++ {
				cnt += float64(class.ComputeLoad(clusterings[i][j]))
				fmt.Println(float64(class.ComputeLoad(clusterings[i][j])))
				if float64(len(clusterings[i][j].TaskQueue)) == 0 {
					cnt_sigma += 1
				} else {
					cnt_sigma += float64(clusterings[i][j].ComputeTaskNumber) / float64(len(clusterings[i][j].TaskQueue))
				}
			}
		}
		load = append(load, cnt/float64(len(edges)))
		sigma = append(sigma, cnt_sigma/float64(len(edges)))
	}
	util.WriteInCSV(data, t, load, sigma, "experiment_EdgeOnlyWithClusterings.csv")
}
