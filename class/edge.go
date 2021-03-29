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
	Cluster           int
}

func GenerateEdge(lx int, ly int, id int) Edge {
	var computePower = float32(rand.Intn(4)+1) * 1e9
	return Edge{id, lx, ly, []Task{},
		[]Task{}, []Task{}, 0, computePower, 0, -1}
}

func TaskComputeTime(t Task, e Edge) float32 {
	return float32(t.Data * t.ExecuteCircleForOneBit) / float32(e.ComputePower)
}

func InitEdge(o int, edges []Edge){
	for i := 0; i < len(edges); i++ {
		numberOfTasks := rand.Intn(TASK_NUMBER_UPPER_LIMIT[o])
		edges[i].ComputeTaskNumber = 0
		edges[i].TaskQueue = GenerateTaskQueue(numberOfTasks, edges[i].Id * TASK_NUMBER_UPPER_LIMIT[o])
		edges[i].MigQueue = []Task{}
		edges[i].ProcQueue = []Task{}
		edges[i].TimeLine = 0
	}
}

func ComputeLoad(e Edge) float32{
	var cnt float32 = 0.0
	for i:=0; i<len(e.ProcQueue); i++ {
		cnt += float32(e.ProcQueue[i].Data * e.ProcQueue[i].ExecuteCircleForOneBit)
	}
	return 2 * cnt / e.ComputePower
}
