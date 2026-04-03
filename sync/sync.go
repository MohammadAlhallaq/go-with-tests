package sync

import "sync"

type Counter struct {
	mu      sync.Mutex
	counter int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counter++
}

func (c *Counter) Value() int {
	return c.counter
}
