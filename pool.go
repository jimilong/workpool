package worker

type Pool struct {
	workers []*Worker
	pool    chan *Worker
}

func NewPool(opt *options) *Pool {
	p := &Pool{
		workers: make([]*Worker, opt.maxOpen),
		pool:    make(chan *Worker, opt.maxOpen),
	}

	for i := 0; i < opt.maxOpen; i++ {
		worker := NewWorker(opt)
		p.workers[i] = worker
		p.pool <- worker
	}

	return p
}

func (p *Pool) Stop() {
	for _, worker := range p.workers {
		if !worker.IsClosed() {
			worker.Stop()
		}
	}
	close(p.pool)
}

func (p *Pool) Get() *Worker {
	worker := <-p.pool
	if worker.IsClosed() {
		worker.Start()
	}

	return worker
}

func (p *Pool) Put(w *Worker) {
	p.pool <- w
}
