//go:build ignore
// +build ignore

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
)

func Foo(m sync.Mutex) {
	m.Lock()
	defer m.Unlock()
	// какие-то действия
}
func main() {
	var wg sync.WaitGroup
	for _, v := range []int{0, 1, 2, 3} {
		wg.Add(1)
		go func() {
			fmt.Print(v)
			wg.Done()
		}()
	}
	wg.Wait()
	ctx, _ := context.WithCancel(context.Background())
	CtxFunc(ctx)
	url := fmt.Sprintf(`%s://%s:%d`, `https`, 3235, `localhost`)
	fmt.Print(url)
	j, err := json.Marshal(Foo1{Name: "JohnDoe"})
	fmt.Println(string(j), err)

}

func CtxFunc(ctx context.Context) {}

type Foo1 struct {
	Name string `json: "nickname"`
}

func Never() {
	switch 10 {
	case 10:
		// что-то делаем
		return
	case 20:
		// что-то делаем
		return
	default:
		// что-то делаем
		return
	}
	fmt.Println("Никогда не выполнится")
}
