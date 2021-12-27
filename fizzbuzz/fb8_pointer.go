package main

import (
	"flag"
	"fmt"
)

// ポインタを利用して書き換えを行うfizzbuzz

type FbResult struct {
	Number int
	IsFizz bool
	IsBuzz bool
}

func main() {
	var size int
	flag.IntVar(&size, "size", 100, "fizz buzz size")
	flag.Parse()

	arr := make([]FbResult, size)
	for i, v := range arr {
		createResult(&v, i)
		fmt.Println(createResultStr(v))
	}
}

func createResult(r *FbResult, i int) {

	r.Number = i
	r.IsFizz = isEqual(3, i)
	r.IsBuzz = isEqual(5, i)

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
