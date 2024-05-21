package workerpool

import "fmt"

type Worker struct {
	ID       int
	taskChan chan *Task
	quit     chan bool
}

// NewWorker returns new instance of worker
func NewWorker(channel chan *Task, ID int) *Worker {
	return &Worker{
		ID:       ID,
		taskChan: channel,
		quit:     make(chan bool),
	}
}

// StartBackground starts the worker in background waiting
func (wr *Worker) StartBackground() {
	fmt.Printf("Starting worker %d\n", wr.ID)

	for {
		select {
		case task := <-wr.taskChan:
			task.process(wr.ID)
		case <-wr.quit:
			return
		}
	}
}

// Stop quits the worker
func (wr *Worker) Stop() {
	fmt.Printf("Closing worker %d\n", wr.ID)
	go func() {
		wr.quit <- true
	}()
}
