package cloud_edge

import (
	"github.com/user/class"
	"github.com/user/util"
)

func GetInitBipartiteGraph(clustering []class.Edge) ([]class.Edge, []class.Edge){
	var MigSide, RecSide = make([]class.Edge,0,len(clustering)),make([]class.Edge,0,len(clustering))
	for i:=0; i<len(clustering); i++ {
		util.EDF(&clustering[i])
		if len(clustering[i].MigQueue) > 0 {
			MigSide = append(MigSide, clustering[i])
		} else {
			RecSide = append(RecSide, clustering[i])
		}
	}
	return MigSide, RecSide
}