package main

import (
	"fmt"
	"time"

	"k8s.io/client-go/util/workqueue"
)

func fake(name string, result bool, err error) (bool, error) {
	fmt.Printf("execute fake func [%v] result [%v], err [%v] in [%v]\n", name, result, err, time.Now())
	time.Sleep(time.Duration(5) * time.Second)
	return result, err
}

type param struct {
	name   string
	result bool
	err    error
}

func QueueDemo() {
	ps := []*param{
		{
			name:   "t1",
			result: true,
			err:    nil,
		},
		{
			name:   "t2",
			result: true,
			err:    fmt.Errorf("t2 failed"),
		},
		{
			name:   "t3",
			result: true,
			err:    fmt.Errorf("t3 failed"),
		},
	}

	var queue workqueue.RateLimitingInterface

	queue = workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "test")

	go func() {
		for i, _ := range ps {
			queue.Add(ps[i])
		}
	}()

	for {
		p, quit := queue.Get()
		if quit {
			break
		}

		pa := p.(*param)

		_, err := fake(pa.name, pa.result, pa.err)
		if err == nil {
			queue.Forget(p)
		} else {
			queue.AddRateLimited(p)
		}

		queue.Done(p)
	}

}
