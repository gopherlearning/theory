package main

import (
	"crypto/aes"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/sha3"
)

func generateRandom(size int) ([]byte, error) {
	// генерируем случайную последовательность байт
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func main() {
	// подписываемое сообщение
	src := []byte("Видишь гофера? Нет. И я нет. А он есть.")

	// создаём случайный ключ
	key, err := generateRandom(16)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	// подписываем алгоритмом HMAC, используя SHA256
	h := hmac.New(sha256.New, key)
	h.Write(src)
	dst := h.Sum(nil)

	fmt.Printf("%x\n", dst)
	main2()
	main3()
}

var secretkey = []byte("secret keyvcbgdc")

func main2() {
	var (
		data []byte // декодированное сообщение с подписью
		id   uint32 // значение идентификатора
		err  error
		sign []byte // HMAC-подпись от идентификатора
	)
	msg := "048ff4ea240a9fdeac8f1422733e9f3b8b0291c969652225e25c5f0f9f8da654139c9e21"

	// допишите код
	// 1) декодируйте msg в data
	data, err = hex.DecodeString(msg)
	if err != nil {
		panic(err)
	}
	fmt.Println(data[:4])
	// 2) получите идентификатор из первых четырёх байт,
	//    используйте функцию binary.BigEndian.Uint32
	id = binary.BigEndian.Uint32(data[:4])
	// 3) вычислите HMAC-подпись sign для этих четырёх байт
	h := hmac.New(sha256.New, secretkey)
	h.Write(data[:4])
	sign = h.Sum(nil)
	fmt.Println(sign)
	fmt.Println(string(sign))
	// ...

	if hmac.Equal(sign, data[4:]) {
		fmt.Println("Подпись подлинная. ID:", id)
	} else {
		fmt.Println("Подпись неверна. Где-то ошибка")
	}
}

func main3() {
	// var (
	// 	data []byte // декодированное сообщение с подписью
	// 	id   uint32 // значение идентификатора
	// 	err  error
	// 	sign []byte // HMAC-подпись от идентификатора
	// )
	msg := 1

	fmt.Println(signID(int64(msg)))

	msg = 10

	fmt.Println(signID(int64(msg)))

	msg = 55

	fmt.Println(signID(int64(msg)))

	// enc := hex.EncodeToString([]byte(msg))
	// fmt.Println(enc)
	// // допишите код
	// // 1) декодируйте msg в data
	// data, err = hex.DecodeString(enc)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(data))
	// // 2) получите идентификатор из первых четырёх байт,
	// //    используйте функцию binary.BigEndian.Uint32
	// id = binary.BigEndian.Uint32(data[:4])
	// // 3) вычислите HMAC-подпись sign для этих четырёх байт
	// h := hmac.New(sha256.New, secretkey)
	// h.Write(data[:4])
	// sign = h.Sum(nil)
	// fmt.Println(sign)
	// fmt.Println(string(sign))
	// // ...

	// if hmac.Equal(sign, data[4:]) {
	// 	fmt.Println("Подпись подлинная. ID:", id)
	// } else {
	// 	fmt.Println("Подпись неверна. Где-то ошибка")
	// }
}

func signID(id int64) string {

	enc := hex.EncodeToString([]byte(fmt.Sprint(id + 10000000000)))

	aesblock, err := aes.NewCipher(secretkey)
	if err != nil {
		return ""
	}

	dst := make([]byte, aes.BlockSize) // зашифровываем
	aesblock.Encrypt(dst, []byte(fmt.Sprint(id+10000000000)))
	fmt.Printf("encrypted: %x\n", dst)
	h := hmac.New(sha3.New224, secretkey)

	h.Write([]byte(enc))
	sign := h.Sum(nil)
	return hex.EncodeToString(sign)
}
