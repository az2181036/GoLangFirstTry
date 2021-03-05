package main

import (
	"../cloudedge"
	"../util"
	"encoding/csv"
	"os"
	"strconv"
)

func main() {
	f1, err := os.Create("clusterings.csv")
	if err != nil {
		panic(err)
	}
	f2, err := os.Create("centers.csv")
	if err != nil {
		panic(err)
	}

	defer f1.Close()
	f1.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
	w := csv.NewWriter(f1)
	w.Write([]string{"X", "Y", "Clusterings"})
	edges := cloud_edge.GenerateEdges()
	c, center := util.Clustering(12, edges)
	for i := 0; i < len(c); i++ {
		for j := 0; j < len(c[i]); j++ {
			w.Write([]string{strconv.Itoa(c[i][j].Lx), strconv.Itoa(c[i][j].Ly), strconv.Itoa(i)})
		}
	}
	w.Flush()

	defer f2.Close()
	f2.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
	w = csv.NewWriter(f2)
	w.Write([]string{"X", "Y", "center"})
	for i := 0; i < len(center); i++ {
		w.Write([]string{strconv.FormatFloat(center[i].X, 'f', 2, 64),
			strconv.FormatFloat(center[i].Y, 'f', 2, 64)})
	}
	w.Flush()
}
