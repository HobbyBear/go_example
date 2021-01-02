package pool

import (
	"fmt"
	"time"
)

type Worker struct {
	task            chan f
	lastExecuteTime time.Time
	pool            *Pool
}

func newWorker(p *Pool) *Worker {
	w := &Worker{
		task: make(chan f, 1),
		pool: p,
	}
	w.loop()
	return w
}

func (w *Worker) loop() {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				<-w.pool.RunningCount
				fmt.Println(err)
			}
		}()
		for f := range w.task {
			if f == nil {
				<-w.pool.RunningCount
				return
			}
			f()
			w.lastExecuteTime = time.Now()
			w.pool.workers <- w
		}
	}()
}
