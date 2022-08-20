package main

import (
	"fmt"
	"runtime"
	"sync"
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
	arr  []*Task
	mu   sync.Mutex
	cond *sync.Cond
}

func NewQueue() *Queue {
	q := Queue{}
	q.cond = sync.NewCond(&q.mu) // переменная условия использует мьютекс
	return &q
}

func (q *Queue) Push(t *Task) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.arr = append(q.arr, t)
	q.cond.Signal() // пробуждаем горутину
}

func (q *Queue) PopWait() *Task {
	q.mu.Lock()
	defer q.mu.Unlock()

	for len(q.arr) == 0 {
		// эта функция отпустит мьютекс, усыпит горутину
		// после пробуждения функция возьмёт мьютекс обратно
		q.cond.Wait()
	}

	t := q.arr[0]
	q.arr = q.arr[1:]

	return t
}

type Resizer struct {
}

func NewResizer() *Resizer {
	r := Resizer{}
	return &r
}

func (r *Resizer) Resize(filename string, toSize Size) error {
	// пропустим реализацию
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

	filenames := []string{"gopher.jpg", "test.png"}
	for _, f := range filenames {
		queue.Push(&Task{Filename: f, ToSize: Size{Width: 1024, Height: 1024}})
	}

	time.Sleep(1 * time.Second)

	var mu sync.Mutex
	m := make(map[int]int)

	for i := 0; i < 100; i++ {
		go func() {
			mu.Lock()
			m[i] = i
			mu.Unlock()
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println(len(m))

	var mu1 sync.Mutex
	m1 := make(map[int]int)

	for i := 0; i < 100; i++ {
		go func(v int) {
			mu1.Lock()
			m1[v] = v
			mu1.Unlock()
		}(i)
	}
	fmt.Println(len(m1))
	//////////////////

	var mu2 sync.Mutex
	m2 := make(map[int]int)
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(v int) {
			mu2.Lock()
			m2[v] = v
			mu2.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(len(m2))
}
