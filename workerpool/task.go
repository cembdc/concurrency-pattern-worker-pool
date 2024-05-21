package workerpool

import (
	"fmt"
)

type Task struct {
	Err  error
	Data interface{}
	f    func(interface{}) error
}

func NewTask(f func(interface{}) error, data interface{}) *Task {
	return &Task{f: f, Data: data}
}

func (t *Task) process(workerID int) {
	fmt.Printf("Worker %d processes task %v\n", workerID, t.Data)
	t.Err = t.f(t.Data)
}
