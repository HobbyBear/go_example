package pool

import (
	"sync"
	"time"
)

type Pool struct {
	RunningCount   chan struct{} // 工作协程数量
	workers        chan *Worker  // 空闲协程
	lock           sync.RWMutex
	HandleIdleTime time.Duration
	Expiration     time.Duration
}


type f func()

type Setter func(*Pool) error

func WithMaxRunningNum(maxRunningNum int) Setter {
	return func(pool *Pool) error {
		pool.RunningCount = make(chan struct{}, maxRunningNum)
		pool.workers = make(chan *Worker, maxRunningNum)
		return nil
	}
}

func WithHandleIdleTime(handleIdleTime time.Duration) Setter {
	return func(pool *Pool) error {
		pool.HandleIdleTime = handleIdleTime
		return nil
	}
}

func WithExpiration(expiration time.Duration) Setter {
	return func(pool *Pool) error {
		pool.Expiration = expiration
		return nil
	}
}

func NewPool(setters ...Setter) (*Pool, error) {
	p := &Pool{
		RunningCount:   make(chan struct{}, 16),
		workers:        make(chan *Worker, 16),
		HandleIdleTime: 10 * time.Millisecond,
		Expiration:     1 * time.Millisecond,
	}
	for _, s := range setters {
		err := s(p)
		if err != nil {
			return nil, err
		}
	}
	p.cleanIdleWorkers()
	return p, nil
}

func (p *Pool) cleanIdleWorkers() {
	go func() {
		ticker := time.NewTicker(p.Expiration)
		for range ticker.C {
			for w := range p.workers {
				if time.Now().Sub(w.lastExecuteTime) <= p.Expiration {
					p.workers <- w
					break
				}
				w.task <- nil
			}
		}
	}()
}

func (p *Pool) Submit(f f) {
	w := p.getWorker()
	w.task <- f
}


func (p *Pool) getWorker() *Worker {

	select {
	case p.RunningCount <- struct{}{}:
		w := newWorker(p)
		return w
	default:
		w := <-p.workers
		return w
	}
}
