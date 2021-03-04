package util

import (
	"../class"
	"sort"
)

// Shortest Task First

func STF(e *class.Edge) {
	tmptasks := e.TaskQueue
	sort.Sort(class.DataTasks(tmptasks))
	for i := 0; i < len(tmptasks); i++ {
		tmp := e.TimeLine + class.TaskComputeTime(e.TaskQueue[i], *e)
		if tmptasks[i].DeadLine >= tmp {
			e.ProcQueue = append(e.ProcQueue, tmptasks[i])
			e.TimeLine = tmp
			e.ComputeTaskNumber++
		} else {
			e.MigQueue = append(e.MigQueue, tmptasks[i])
		}
	}
}
