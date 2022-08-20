package main

import (
	"fmt"
	"runtime"
	"time"
)

type Size struct {
	Width  uint32
	Height uint32
}

type Task struct {
	Filename string
	ToSize   Size
}

type Queue struct {
	ch chan *Task
}

func NewQueue() *Queue {
	return &Queue{
		ch: make(chan *Task, 1),
	}
}

func (q *Queue) Push(t *Task) {
	q.ch <- t
}

func (q *Queue) PopWait() *Task {
	return <-q.ch
}

type Resizer struct {
}

func NewResizer() *Resizer {
	r := Resizer{}
	return &r
}

func (r *Resizer) Resize(filename string, toSize Size) error {
	// пропустим реализацию
	time.Sleep(50 * time.Millisecond)
	return nil
}

type Worker struct {
	id      int
	queue   *Queue
	resizer *Resizer
}

func NewWorker(id int, queue *Queue, resizer *Resizer) *Worker {
	w := Worker{
		id:      id,
		queue:   queue,
		resizer: resizer,
	}
	return &w
}

func (w *Worker) Loop() {
	for {
		t := w.queue.PopWait()

		err := w.resizer.Resize(t.Filename, t.ToSize)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			continue
		}

		fmt.Printf("worker #%d resized %s\n", w.id, t.Filename)
	}
}

func main() {
	queue := NewQueue()
	workers := make([]*Worker, 0, runtime.NumCPU())
	for i := 0; i < runtime.NumCPU(); i++ {
		workers = append(workers, NewWorker(i, queue, NewResizer()))
	}

	for _, w := range workers {
		go w.Loop()
	}

	filenames := []string{"gopher.jpg", "test.png", "nonce.jpg"}
	for _, f := range filenames {
		queue.Push(&Task{Filename: f, ToSize: Size{Width: 1024, Height: 1024}})
	}

	time.Sleep(1 * time.Second)

	a := make(chan int)
	b := make(chan int)

	go func() {
		a <- 100
	}()
	go func() {
		v := <-a
		b <- v
	}()
	fmt.Println(<-b)
	////////
	ch := make(chan int, 7)
	// специально делаем буфер канала меньше,
	// чем количество чисел Фибоначчи
	go fibonacci(15, ch)

	for i := range ch {
		// считываем значения из канала, пока он не будет закрыт
		fmt.Printf("%d ", i)
	}
}

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x // посылаем значения в канал
		x, y = y, x+y
	}

	close(ch) // закрываем канал
}
