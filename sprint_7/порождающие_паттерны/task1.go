//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	fmt.Println("task2")
	defer Close()

	go Dump([]byte("run 1"))
	go Dump([]byte("run 2"))
	go Dump([]byte("run 3"))
	go Dump([]byte("run 4"))
}

var (
	dumpFile  *os.File
	onceOpen  sync.Once
	onceClose sync.Once
)

func Dump(data []byte) error {
	onceOpen.Do(func() {
		fname, err := os.Executable()
		if err == nil {
			dumpFile, err = os.Create(fname + `.dump`)
		}
		if err != nil {
			panic(err)
		}
		fmt.Println("Инициализируем singleton")
	})

	_, err := dumpFile.Write(data)
	return err
}

func Close() {
	onceClose.Do(func() {
		if dumpFile != nil {
			dumpFile.Close()
		}
	})
}
