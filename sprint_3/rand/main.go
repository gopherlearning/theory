package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func main() {
	// определяем слайс нужной длины

	s, err := RandBytes(32) // записываем байты в массив b
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Println(s)

}

func RandBytes(n int) (string, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return ``, err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}
