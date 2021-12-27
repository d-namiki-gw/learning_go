package main

import (
	"flag"
	"fmt"
)

// flagを利用したオプション取得を利用したfizzbuzz

type FbResult struct {
	Number int
	IsFizz bool
	IsBuzz bool
}

func main() {
	var size int
	flag.IntVar(&size, "size", 0, "fizz buzz size")
	flag.Parse()

	arr := make([]int, size)
	is_fizz := func(target int) bool { return isEqual(3, target) }
	is_buzz := func(target int) bool { return isEqual(5, target) }

	for i, _ := range arr {
		result := FbResult{
			Number: i,
			IsFizz: is_fizz(i),
			IsBuzz: is_buzz(i),
		}
		fmt.Println(createResultStr(result))
	}

}

func isEqual(div int, num int) bool {
	return num > 0 && num%div == 0
}

func createResultStr(r FbResult) string {
	getText := func(v bool, text string) (r string) {
		if v {
			r = text
		}
		return r
	}

	return fmt.Sprintf(
		"%d %s%s",
		r.Number,
		getText(r.IsFizz, "fizz"),
		getText(r.IsBuzz, "buzz"),
	)
}
