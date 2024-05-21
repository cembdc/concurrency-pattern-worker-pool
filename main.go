package main

import (
	"concurrency_pattern_worker_pool/workerpool"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var allTask []*workerpool.Task
	pool := workerpool.NewPool(allTask, 5)

	go func() {
		for {
			taskID := rand.Intn(100) + 20

			if taskID%7 == 0 {
				fmt.Printf("Task %d detected to close\n", taskID)
				pool.Stop()
			}

			time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
			task := workerpool.NewTask(func(data interface{}) error {
				taskID := data.(int)
				time.Sleep(100 * time.Millisecond)
				fmt.Printf("Task %d processed\n", taskID)
				return nil
			}, taskID)
			pool.AddTask(task)
		}
	}()
	pool.RunBackground()
}
