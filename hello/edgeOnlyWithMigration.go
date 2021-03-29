package main

import (
	"../class"
	"../cloudedge"
	"../util"
	"time"
)

func main() {
	var t []float64
	var data [][]int
	var load []float64
	var sigma []float64
	edges := cloud_edge.GenerateEdges()
	for o := 0; o < len(class.TASK_NUMBER_UPPER_LIMIT); o++ {
		class.InitEdge(o, edges)
		for i := 0; i < len(edges); i++ {
			util.EDF(&edges[i])
		}
		startime := time.Now().UnixNano()
		data = append(data, util.EdgeOnlyWithMigration(edges))
		endtime := time.Now().UnixNano()
		execTime := float64((endtime - startime) / 1e6)
		t = append(t, execTime)

		cnt := 0.0
		cnt_sigma := 0.0
		for i := 0; i < len(edges); i++ {
			cnt += float64(class.ComputeLoad(edges[i]))
			if float64(len(edges[i].TaskQueue)) == 0 {
				cnt_sigma += 1
			}else{
				cnt_sigma += float64(edges[i].ComputeTaskNumber) / float64(len(edges[i].TaskQueue))
			}
		}
		load = append(load, cnt/float64(len(edges)))
		sigma = append(sigma, cnt_sigma/float64(len(edges)))
	}
	util.WriteInCSV(data, t, load, sigma, "experiment_EdgeOnlyWithMigration.csv")
}
