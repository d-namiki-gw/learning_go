package main

import "fmt"

// 無名関数を利用したfizzbuzz

func main() {
	arr := [100]int{}
	is_fizz := func(target int) bool { return isEqual(3, target) }
	is_buzz := func(target int) bool { return isEqual(5, target) }
	for i, _ := range arr {
		fmt.Print(i, " ")
		if is_fizz(i) {
			fmt.Print("fizz")
		}
		if is_buzz(i) {
			fmt.Print("buzz")
		}
		fmt.Println("")
	}

}

func isEqual(div int, num int) bool {
	return num > 0 && num%div == 0
}
