package main

import (
	"fmt"
	"sync"
	"time"
)

type Cond struct {
	L sync.Locker
	q chan struct{}
}

func NewCond(l sync.Locker) *Cond {
	return &Cond{L: l, q: make(chan struct{})}
}

func (c *Cond) Wait() {
	c.L.Unlock()
	c.q <- struct{}{}
	c.L.Lock()
}

func (c *Cond) Signal() {
	select {
	case <-c.q:
	default:
	}
}

func (c *Cond) Broadcast() {
	for {
		select {
		case <-c.q:
		default:
			return
		}
	}
}

type State struct {
	ready bool
	cond  *Cond
}

func NewState() *State {
	return &State{cond: NewCond(&sync.Mutex{})}
}

func (s *State) WaitReady() {
	s.cond.L.Lock()
	defer s.cond.L.Unlock()

	for !s.ready {
		s.cond.Wait()
	}
}

func (s *State) SetReady() {
	s.cond.L.Lock()
	defer s.cond.L.Unlock()

	s.ready = true
	s.cond.Broadcast()
}

func main() {
	s := NewState()

	go func() {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("now ready")
		s.SetReady()
	}()

	for i := 0; i < 5; i++ {
		go func() {
			s.WaitReady()
			fmt.Println("ready!")
		}()
	}
	time.Sleep(1000 * time.Millisecond)
}
