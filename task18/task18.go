package main

import "sync"

type CCounterStatus int

const (
	Running = iota
	Done
)

type CCounter struct {
	status int
	input  chan int
	value  int
	mut    *sync.RWMutex
}

func NewConcurentCounter() CCounter {
	ret := CCounter{
		Running,
		make(chan int, 10),
		0,
		&sync.RWMutex{},
	}

	go func() {
		for n := range ret.input {
			ret.mut.Lock()
			ret.value += n
			ret.mut.Unlock()
		}
	}()

	return ret
}

func (c *CCounter) Status(n int) int {
	ret := -1

	if c.status == Running {
		c.mut.RLock()
		ret = c.status
		c.mut.RUnlock()
	} else {
		ret = c.status
	}

	return ret
}

func (c *CCounter) Add(n int) {
	c.input <- n
}

func (c *CCounter) Peek(n int) int {
	ret := 0

	if c.status == Running {
		c.mut.RLock()
		ret = c.value
		c.mut.RUnlock()
	} else {
		ret = c.value
	}

	return ret
}

func (c *CCounter) Stop(n int) {
	c.status = Done
	close(c.input)

	for n := range c.input {
		c.value += n
	}
}
