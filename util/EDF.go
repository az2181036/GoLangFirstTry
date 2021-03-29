package util

import (
	"../class"
)

// Earliest deadline first

func EDF(e *class.Edge) {
	for i := 0; i < len(e.TaskQueue); i++ {
		tmp := e.TimeLine + class.TaskComputeTime(e.TaskQueue[i], *e)
		if e.TaskQueue[i].DeadLine >= tmp {
			e.ProcQueue = append(e.ProcQueue, e.TaskQueue[i])
			e.TimeLine = tmp
			e.ComputeTaskNumber++
		} else {
			e.MigQueue = append(e.MigQueue, e.TaskQueue[i])
		}
	}
}
