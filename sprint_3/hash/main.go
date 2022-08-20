package main

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func main() {
	src := []byte("Здесь могло быть написано, чем Go лучше Rust. " +
		"Но после хеширования уже не прочитаешь.")

	h := sha256.New()
	h.Write(src)
	dst := h.Sum(nil)

	fmt.Printf("%x\n", dst)

	main2()
}
func main2() {

	var (
		data  []byte         // слайс случайных байт
		hash1 []byte         // хеш с использованием интерфейса hash.Hash
		hash2 [md5.Size]byte // хеш, возвращаемый функцией md5.Sum
	)

	// допишите код
	// 1) генерация data длиной 512 байт
	data = make([]byte, 512)
	_, err := rand.Read(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 2) вычисление hash1 с использованием md5.New
	h := md5.New()
	h.Write(data)
	hash1 = h.Sum(nil)
	// 3) вычисление hash2 функцией md5.Sum
	hash2 = md5.Sum(data)
	// ...

	// hash2[:] приводит массив байт к слайсу
	if bytes.Equal(hash1, hash2[:]) {
		fmt.Println("Всё правильно! Хеши равны")
	} else {
		fmt.Println("Что-то пошло не так")
	}
}
