package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/alecthomas/kong"
)

var (
	cli struct {
		Input  string `help:"Входные данные"`
		Output string `help:"Выходные данные"`
	}
	err error
)

func init() {
	kong.Parse(&cli)
	// if err != nil {
	// 	fmt.Println("ошибка чтения аргументов: %v", err)
	// 	return
	// }
}

func main() {
	var input io.ReadCloser = os.Stdin
	if len(cli.Input) != 0 {
		input, err = os.Open(cli.Input)
		if err != nil {
			fmt.Println("ошибка чтения файла: ", err)
			return
		}
		defer input.Close()
	}

	var output io.WriteCloser = os.Stdout
	if len(cli.Input) != 0 {
		output, err = os.Open(cli.Input)
		if err != nil {
			fmt.Println("ошибка записи в файл: ", err)
			return
		}
		defer input.Close()
	}

	reader := bufio.NewReader(input)

	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("ошибка чтения строки: %s\n", text)
		return
	}

	in := strings.Split(text, " ")
	if len(in) != 2 {
		fmt.Println("неверный формат ввода 2")
		return
	}

	nums := make([]int, 2)

	for i, v := range in {
		var num int
		fmt.Println(v)
		num, err = strconv.Atoi(strings.TrimRight(v, "\n"))

		if err != nil {
			fmt.Println(i, "неверный формат ввода 1", v)
			return
		}

		nums[i] = num
	}

	_, err = io.Copy(output, bytes.NewBufferString(fmt.Sprint(nums[0]+nums[1])))
	if err != nil {
		fmt.Println("ошибка вывода")
	}
}
