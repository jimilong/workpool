package worker

import "time"

type options struct {
	maxOpen     int
	jobQueueLen int
	timeout     int64
	tickerTime  time.Duration
}

type Option func(*options)

func (o *Option) WithMaxOpen(maxOpen int) Option {
	return func(o *options) {
		o.maxOpen = maxOpen
	}
}

func (o *Option) WithJobQueueLen(len int) Option {
	return func(o *options) {
		o.jobQueueLen = len
	}
}

func (o *Option) WithTimeout(timeout int64) Option {
	return func(o *options) {
		o.timeout = timeout
	}
}

func (o *Option) WithTickerTime(ticker time.Duration) Option {
	return func(o *options) {
		o.tickerTime = ticker
	}
}
