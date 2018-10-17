package worker

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Job interface {
	Do() error
}

type Worker struct {
	wg   sync.WaitGroup
	job  chan Job
	quit chan bool

	closed     int32
	tickerTime time.Duration
	timeout    int64
	lastTime   int64
}

func NewWorker(opt *options) *Worker {
	return &Worker{
		job:        make(chan Job, opt.jobQueueLen),
		quit:       make(chan bool),
		closed:     1,
		tickerTime: opt.tickerTime,
		timeout:    opt.timeout,
	}
}

func (w *Worker) Start() {
	w.wg.Add(1)
	go func() {
		defer w.wg.Done()
		t := time.NewTicker(w.tickerTime)
		for {
			select {
			case job := <-w.job:
				w.lastTime = time.Now().Unix()
				job.Do()

			case <-w.quit:
				return
			case <-t.C: //定时检查协程是否空闲,空闲则关闭协程
				if len(w.job) == 0 && time.Now().Unix() > w.timeout+w.lastTime {
					t.Stop()

					atomic.StoreInt32(&w.closed, 1)
					fmt.Printf("goroutine idle exit\n")
					return
				}
			}
		}
	}()

	atomic.StoreInt32(&w.closed, 0)
}

func (w *Worker) IsClosed() bool {
	return atomic.LoadInt32(&w.closed) == 1
}

func (w *Worker) Stop() {
	atomic.StoreInt32(&w.closed, 1)
	w.quit <- true

	w.wg.Wait()
}

func (w *Worker) AddJob(job Job) {
	w.job <- job
}
