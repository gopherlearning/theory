//go:build ignore
// +build ignore

// В Go есть конструкция for — range, позволяющая обходить в цикле мапы, каналы, слайсы и массивы. Для других видов последовательностей, например списка или дерева, конструкция не работает. Как тогда обходить элементы последовательностей, для которых возможна итерация? Паттерн Итератор позволяет делать такой обход без знания об устройстве последовательности и её элементов.
// Паттерн Итератор используется, когда:
// нужен общий интерфейс для обхода различных структур данных;
// нужно дать возможность перебирать элементы объекта, но при этом скрыть его внутреннюю структуру.
package main

import "fmt"

func newEven() func() int {
	n := 0
	// функциональный литерал замкнёт переменную n
	return func() int {
		n += 2
		return n
	}
}

func main() {
	next := newEven()
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
}
