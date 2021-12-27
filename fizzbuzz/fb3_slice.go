package main

import (
	"fmt"
	"strconv"
	"strings"
)

// スライスで結果を保持するfizzbuzz

func main() {
	arr := [100]int{}
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
