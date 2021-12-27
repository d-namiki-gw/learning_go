package main

// ベーシックなfizzbuzz

import (
	"fmt"
)

func main() {
	arr := [100]int{}
	for i, _ := range arr {
		fmt.Print(i, " ")
		if isFizz(i) {
			fmt.Print("fizz")
		}
		if isBuzz(i) {
			fmt.Print("buzz")
		}
		fmt.Println("")
	}

}

func isFizz(num int) bool {
	return num > 0 && num%3 == 0
}

func isBuzz(num int) bool {
	return num > 0 && num%5 == 0
}
