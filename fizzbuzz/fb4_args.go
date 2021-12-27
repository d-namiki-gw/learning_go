package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// コマンドラインから実行する数を指定できるfizzbuzz

func main() {
	if len(os.Args) < 2 {
		return
	}

	params := os.Args[1:]
	size, err := strconv.Atoi(params[0])

	if err != nil {
		fmt.Println(err)
		return
	}

	arr := make([]int, size)
	is_fizz := func(target int) bool { return isEqual(3, target) }
	is_buzz := func(target int) bool { return isEqual(5, target) }

	for i, _ := range arr {
		result := []string{strconv.Itoa(i)}

		if is_fizz(i) {
			result = append(result, "fizz")
		}
		if is_buzz(i) {
			result = append(result, "buzz")
		}
		fmt.Println(strings.Join(result, " "))
	}

}

func isEqual(div int, num int) bool {
	return num > 0 && num%div == 0
}
