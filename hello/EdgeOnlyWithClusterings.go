package main

import (
	"../cloudedge"
	"../util"
)

func main() {
	var data [][]int
	for o := 0; o < 4; o++ {
		edges := cloud_edge.GenerateEdges(o)
		for i := 0; i < len(edges); i++ {
			util.EDF(&edges[i])
		}
		data = append(data, util.EdgeOnlyWithClusterings(edges, o))
	}
	util.WriteInCSV(data, "experiment_EdgeOnlyWithClusterings.csv")
}
