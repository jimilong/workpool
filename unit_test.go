package worker

import (
	"fmt"
	"testing"
)

func TestPool(t *testing.T) {
	svc := NewService()
	for i := 0; i < 1000000; i++ {
		svc.SubmitJob(newJob(i))
	}

	svc.Stop()
	fmt.Printf("done\n")
}

type testJob struct {
	index int
}

func newJob(i int) *testJob {
	return &testJob{index: i}
}

func (t *testJob) Do() error {
	fmt.Printf("job index: %d\n", t.index)

	return nil
}
