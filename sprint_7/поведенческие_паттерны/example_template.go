//go:build ignore
// +build ignore

// Шаблонный метод определяет основу алгоритма и даёт объектам возможность переопределить некоторые шаги.
// Паттерн Шаблонный метод используется, когда:
// есть основная часть алгоритма, но детали могут различаться для объектов разных типов;
// нужно расширить алгоритм, не изменяя основных шагов.
package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

type ByAge []Person

// реализуем интерфейс sort.Interface для сортировки по возрасту

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// Sort сортирует слайс ByAge, так он реализует интерфейс sort.Interface.
func (a ByAge) Sort() {
	sort.Sort(a)
}
