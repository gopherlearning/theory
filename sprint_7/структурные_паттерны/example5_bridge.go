//go:build ignore
// +build ignore

// Мост — это структурный паттерн проектирования, который позволяет разделить монолитный объект на две иерархии: абстракцию и реализацию.
// Благодаря этому можно изменять их код независимо друг от друга. Объекты абстракции будут ссылаться на объекты реализации и делегировать им основную работу.
// Так как у всех реализаций один интерфейс, их можно будет взаимозаменять внутри абстракции.
// Паттерн Мост используется, когда:
// - нужно разделить монолитный объект, содержащий несколько реализаций функциональности;
// - нужно расширить объект в двух независимых направлениях — можно выделить одно направление в отдельную иерархию;
// - нужно менять реализацию во время выполнения программы.

package main

import "fmt"

// Computer — абстракция компьютера.
type Computer interface {
	Print()
	SetPrinter(Printer)
}

// Mac — компьютер Mac.
type Mac struct {
	printer Printer
}

func (m *Mac) Print() {
	fmt.Println("Печать для Mac.")
	m.printer.PrintFile()
}

func (m *Mac) SetPrinter(p Printer) {
	m.printer = p
}

// Windows — компьютер Windows.
type Windows struct {
	printer Printer
}

func (w *Windows) Print() {
	fmt.Println("Печать для Windows.")
	w.printer.PrintFile()
}

func (w *Windows) SetPrinter(p Printer) {
	w.printer = p
}

// Printer — интерфейс для принтера.
type Printer interface {
	PrintFile()
}

type Epson struct {
}

func (p *Epson) PrintFile() {
	fmt.Println("Печать на принтере Epson.")
}

type HP struct {
}

func (p *HP) PrintFile() {
	fmt.Println("Печать на принтере HP.")
}

func main() {
	// создаём два принтера
	hp := &HP{}
	epson := &Epson{}

	// печать на Mac
	mac := &Mac{}
	mac.SetPrinter(hp)
	mac.Print()
	mac.SetPrinter(epson)
	mac.Print()

	// печать на Windows
	win := &Windows{}
	win.SetPrinter(hp)
	win.Print()
	win.SetPrinter(epson)
	win.Print()
}
