package main

import (
	"fmt"
)

func thread(dec int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- (dec*10 + i)
		}
		close(ch)
	}()
	return ch
}

func main() {
	a := thread(1)
	b := thread(2)
	c := make(chan int)
	go func() {
		for a != nil || b != nil {
			select {
			case v, ok := <-a:
				if !ok {
					a = nil
					continue
				}
				c <- v
			case v, ok := <-b:
				if !ok {
					b = nil
					continue
				}
				c <- v
			}
		}
		close(c)

		// допишите код
		// добавьте цикл с оператором select
		// не забудьте в конце закрыть канал 'c'
		// ...
	}()
	for v := range c {
		fmt.Println(v)
	}
}
