package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type pool struct {
	queue chan int
	wg    *sync.WaitGroup
}

func New(size int) *pool {
	if size <= 0 {
		size = 1
	}
	return &pool{
		queue: make(chan int, size),
		wg:    &sync.WaitGroup{},
	}
}

func (p *pool) Add(delta int) {
	for i := 0; i < delta; i++ {
		p.queue <- 1
	}
	for i := 0; i > delta; i-- {
		<-p.queue
	}
	p.wg.Add(delta)
}

func (p *pool) Done() {
	//fmt.Println("done", <-p.queue)
	<-p.queue
	p.wg.Done()
}

func (p *pool) Wait() {
	p.wg.Wait()
}

func main() {
	pool := New(200)
	fmt.Println(runtime.NumGoroutine())
	fmt.Println("add work start")
	for i := 0; i < 1000; i++ {
		pool.Add(1)
		go func() {
			time.Sleep(time.Second)
			fmt.Println(runtime.NumGoroutine())
			pool.Done()
		}()
	}

	fmt.Println("add work over")
	pool.Wait()
	fmt.Println(runtime.NumGoroutine())
}
