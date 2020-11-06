package class

import (
	"math/rand"
	"sort"
)

type Task struct {
	Data                   int
	ExecuteCircleForOneBit int
	DeadLine               float32
	StartTime              float32
	FinishTime             float32
}

type NewTasks []Task

func (t NewTasks) Len() int { return len(t) }
func (t NewTasks) Swap(i, j int) { t[i], t[j] = t[j], t[i] }
func (t NewTasks) Less(i, j int) bool { return t[i].DeadLine < t[j].DeadLine }

var numberOfTasks = rand.Intn(TASK_NUMBER_UPPER_LIMIT)

func GenerateTaskQueue(numberOfTasks int) []Task {
	taskQueue := make([]Task, 0, numberOfTasks)
	for i:=0; i<numberOfTasks; i++ {
		//1GHz = 10^9Hz
		data := rand.Intn(2000) + 1                                                                                          // [0,2000] Kb
		executeCircleForOneBit := rand.Intn(CIRCLE_UPPER_LIMIT-CIRCLE_LOWER_LIMIT+ 1) + CIRCLE_LOWER_LIMIT // [CIRCLE_LOWER_LIMIT, CIRCLE_UPPER_LIMIT)
		deadLine := rand.Float32() * 3 + EPS                                                                           // (0, 3]
		tmpTask := Task{data, executeCircleForOneBit, deadLine,0, 0}
		taskQueue = append(taskQueue, tmpTask)
	}
	sort.Sort(NewTasks(taskQueue))
	return taskQueue
}