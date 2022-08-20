package main

import (
	_ "net/http/pprof" // подключаем пакет pprof

	"github.com/labstack/echo-contrib/pprof"
	"github.com/labstack/echo/v4"
)

const (
	addr    = ":8080"  // адрес сервера
	maxSize = 10000000 // будем растить слайс до 10 миллионов элементов
)

func foo() {
	// полезная нагрузка
	// for {
	// 	var s []int
	// 	for i := 0; i < maxSize; i++ {
	// 		s = append(s, i)
	// 	}
	// }
	for {
		s := make([]int, 0, maxSize)
		for i := 0; i < maxSize; i++ {
			s = append(s, i)
		}
	}
}

func main() {
	go foo() // запускаем полезную нагрузку в фоне
	e := echo.New()
	pprof.Register(e)
	e.Start(addr) // запускаем сервер
}
