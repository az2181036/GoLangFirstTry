package class

import "math/rand"

type Edge struct {
	Id                int
	Lx                int
	Ly                int
	TaskQueue         []Task
	ProcQueue         []Task
	MigQueue          []Task
	TimeLine          float32
	ComputePower      float32
	ComputeTaskNumber int
}

func GenerateEdge(lx int, ly int, id int) Edge {
	var computePower = float32(rand.Intn(4)+1) * 1e9
	return Edge{id, lx, ly, []Task{},
		[]Task{}, []Task{}, 0, computePower, 0}
}

func TaskComputeTime(t Task, e Edge) float32 {
	return float32(t.Data*t.ExecuteCircleForOneBit) / e.ComputePower
}
