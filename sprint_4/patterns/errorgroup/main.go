package main

import (
	"fmt"

	"golang.org/x/sync/errgroup"
)

type Name string

var names = []Name{"Anna", "Ivan", "Fedor", "Katya", "Gleb"}

// Hello — метод типа Name.
func (n Name) Hello() error {
	fmt.Printf("Hello %v!\n", n)
	return nil
}

func main() {
	g := &errgroup.Group{}

	for _, name := range names {
		// вызываем g.Go с method value в качестве аргумента
		g.Go(name.Hello)
	}

	g.Wait()
}
