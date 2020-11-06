package util

import (
	"github.com/user/class"
)

// Earliest deadline first

func EDF(e *class.Edge){
	for i:=0; i<len(e.TaskQueue); i++{
		if e.TaskQueue[i].DeadLine > e.TimeLine {
			e.ProcQueue = append(e.ProcQueue, e.TaskQueue[i])
			e.TimeLine = e.TimeLine + float32(e.TaskQueue[i].Data * e.TaskQueue[i].ExecuteCircleForOneBit) / e.ComputePower
		} else {
			e.MigQueue = append(e.MigQueue, e.TaskQueue[i])
		}
	}
}
