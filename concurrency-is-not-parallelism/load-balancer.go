package main

import (
	"container/heap"
	"math/rand"
	"time"
)

// ref - https://talks.golang.org/2012/waza.slide#46

func main() {
	rand.Seed(time.Now().Unix())

}

type Request struct {
	fn func() int // The operation to perform.
	c  chan int
}

// Worker keps a channel of reqs, some load tracking data
type Worker struct {
	requests chan Request // work to do (buffered channel)
	pending  int          // count of pending tasks
	index    int          // index in the heap
}

type Pool []*Worker

type Balancer struct {
	pool Pool
	done chan *Worker
}

func (b *Balancer) balance(work chan Request) {
	for {
		select {
		case req := <-work:
			b.dispatch(req)
		case w := <-b.done:
			b.complete(w)
		}
	}
}

func (p Pool) Less(i, j int) bool {
	return p[i].pending < p[j].pending
}

func (b *Balancer) dispatch(req Request) {
	w := heap.Pop(&b.pool).(*Worker)
	w.requests <- req
	w.pending++
	heap.Push(&b.pool, w)
}

func (b *Balancer) complete(w *Worker) {
	w.pending--
	heap.Remove(&b.pool, w.index)
	heap.Push(&b.pool, w) // Put it into its place on the heap
}

func (w *Worker) work(done chan *Worker) {
	for {
		req := <-w.requests
		req.c <- req.fn()
		done <- w
	}
}

func Requester(work chan<- Request) {
	c := make(chan int)

	for {
		Sleep()
		work <- Request{workFn, c}
		result := <-c
	}
}

// Sleep kills some time (fake load)
func Sleep() {
	nWorker := 10
	sec := rand.Int63n(int64(nWorker * 2))
	time.Sleep(time.Duration(sec) * time.Second)
}

func workFn() int {
	millis := rand.Intn(1000)
	time.Sleep(time.Duration(millis) * time.Second)
	return 0
}
