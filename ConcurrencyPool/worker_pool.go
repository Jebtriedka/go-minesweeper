package concurrencyPool

import (
	"sync"
)

type WorkerPool struct {
	Tasks       []Task
	Concurrency int

	tasksChan chan Task
	wg        sync.WaitGroup
}

func (wp *WorkerPool) worker() {

	for task := range wp.tasksChan {
		task.Process()
		wp.wg.Done()
	}
}

func (wp *WorkerPool) Run() {

	wp.tasksChan = make(chan Task, len(wp.Tasks))

	for i := 0; i < wp.Concurrency; i++ {

		go wp.worker()
	}

	wp.wg.Add(len(wp.Tasks))

	for _, task := range wp.Tasks {
		wp.tasksChan <- task
	}

	close(wp.tasksChan)

	wp.wg.Wait()
}
