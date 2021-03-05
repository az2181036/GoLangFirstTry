package main

import (
	"../cloudedge"
	"../util"
	"time"
)

func main() {
	var t []float64
	var data [][]int
	edges := cloud_edge.GenerateEdges()
	for o := 0; o < 4; o++ {
		for i := 0; i < len(edges); i++ {
			util.EDF(&edges[i])
		}
		startime := time.Now().UnixNano()
		data = append(data, util.EdgeOnlyWithClusterings(edges, o))
		endtime := time.Now().UnixNano()
		execTime := float64((endtime - startime) / 1e6)
		t = append(t, execTime)
	}
	util.WriteInCSV(data, t, "experiment_EdgeOnlyWithClusterings.csv")
}
