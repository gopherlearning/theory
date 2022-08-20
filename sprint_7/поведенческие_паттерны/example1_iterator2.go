//go:build ignore
// +build ignore

// В Go есть конструкция for — range, позволяющая обходить в цикле мапы, каналы, слайсы и массивы. Для других видов последовательностей, например списка или дерева, конструкция не работает. Как тогда обходить элементы последовательностей, для которых возможна итерация? Паттерн Итератор позволяет делать такой обход без знания об устройстве последовательности и её элементов.
// Паттерн Итератор используется, когда:
// нужен общий интерфейс для обхода различных структур данных;
// нужно дать возможность перебирать элементы объекта, но при этом скрыть его внутреннюю структуру.
package main

import "fmt"

// Iterator — интерфейс для получения следующего элемента.
type Iterator interface {
	Next() (string, bool)
}

// Exported — тип, реализующий интерфейс Iterator.
type Exported struct {
	ID        string
	invisible []string
	cursor    int
}

func NewExported(s ...string) *Exported {
	e := new(Exported)
	e.invisible = append(e.invisible, s...)
	return e
}

// Next — метод, реализующий шаблон Итератор.
func (e *Exported) Next() (string, bool) {
	if e.cursor <= len(e.invisible) {
		e.cursor++
	}
	return e.invisible[e.cursor-1], e.cursor < len(e.invisible)
}

func main() {
	// клиентский код
	e := NewExported("foo", "bar", "buzz", "oups")
	for {
		next, hasNext := e.Next()
		fmt.Println(next)
		if !hasNext {
			break
		}
	}
}
