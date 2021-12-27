package main

import (
	"flag"
	"fmt"
	"strconv"
)

// flagを利用してosと同じ読み込み方法を実装するfizzbuzz

type FbResult struct {
	Number int
	IsFizz bool
	IsBuzz bool
}

func main() {
	flag.Parse()
	if len(flag.Args()) < 1 {
		fmt.Println(flag.Args())
		return
	}

	params := flag.Args()
	size, err := strconv.Atoi(params[0])

	if err != nil {
		fmt.Println(err)
		return
	}

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
