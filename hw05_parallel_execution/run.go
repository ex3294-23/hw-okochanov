package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Run(tasks []Task, n int, m int) error {
	workCh := make(chan Task)

	wg := sync.WaitGroup{}
	wg.Add(n)
	var errCounter int32
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for task := range workCh {
				err := task()
				if err != nil {
					atomic.AddInt32(&errCounter, 1)
				}
			}
		}()
	}
	for _, task := range tasks {
		if atomic.LoadInt32(&errCounter) >= int32(m) {
			break
		}
		workCh <- task
	}

	close(workCh)
	wg.Wait()
	if errCounter >= int32(m) {
		return ErrErrorsLimitExceeded
	}
	return nil
}
