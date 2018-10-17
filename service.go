package worker

import "time"

type Service struct {
	pool *Pool
}

func NewService(opts ...Option) *Service {
	opt := options{
		maxOpen:     10,
		jobQueueLen: 10,
		timeout:     300,
		tickerTime:  10 * time.Second,
	}
	for _, o := range opts {
		o(&opt)
	}

	return &Service{
		pool: NewPool(&opt),
	}
}

func (p *Service) Stop() {
	p.pool.Stop()
}

func (p *Service) SubmitJob(job Job) {
	worker := p.pool.Get()
	worker.AddJob(job)
	p.pool.Put(worker)
}
