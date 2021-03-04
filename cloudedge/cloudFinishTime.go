package cloud_edge

import (
	"../class"
	"math"
)

func CloudFinishTime(timeline float32, task class.Task, dc class.Datacenter) float32 {
	return float32(math.Max(float64(timeline), float64(task.StartTime))) + CloudComputeTime(task, dc)
}

func CloudComputeTime(task class.Task, dc class.Datacenter) float32 {
	return float32(task.Data*task.ExecuteCircleForOneBit) / dc.ComputePower
}
