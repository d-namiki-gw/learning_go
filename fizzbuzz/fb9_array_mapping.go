package main

import (
	"flag"
	"fmt"
	"reflect"
)

// array_mapを実装し利用するfizzbuzz

type FbResult struct {
	Number int
	IsFizz bool
	IsBuzz bool
}

type General interface{}

func main() {
	var size int
	flag.IntVar(&size, "size", 100, "fizz buzz size")
	flag.Parse()

	arr := mapping(
		make([]FbResult, size),
		func(r FbResult, i int) FbResult {
			r.Number = i
			r.IsFizz = isEqual(3, i)
			r.IsBuzz = isEqual(5, i)
			fmt.Println(i)
			return r
		}).([]FbResult)

	for _, v := range arr {
		fmt.Println(createResultStr(v))
	}

}

func mapping(src General, f General) General {

	srcVal := reflect.ValueOf(src)
	mapFunc := reflect.ValueOf(f)

	if srcVal.Kind() == reflect.Slice {
		list := reflect.MakeSlice(
			reflect.TypeOf(src),
			srcVal.Len(),
			srcVal.Cap(),
		)
		len := list.Len()
		// reflect.Valueはrangeで回せなかった
		for i := 0; i < len; i++ {
			value := mapFunc.Call(
				[]reflect.Value{
					list.Index(i),
					reflect.ValueOf(i),
				},
			)[0]
			list.Index(i).Set(value)
		}

		return list.Interface()
	}
	return make([]General, 0)

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
