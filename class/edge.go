package class

import "math/rand"

type Edge struct {
	Lx           int
	Ly           int
	TaskQueue    []Task
	ProcQueue    []Task
	MigQueue     []Task
	TimeLine    float32
	ComputePower float32
}

func GenerateEdge(lx int, ly int) Edge{
	computePower := float32(rand.Intn(4)+1)
	return Edge{lx,ly,GenerateTaskQueue(numberOfTasks),
		[]Task{}, []Task{}, 0,computePower}
}