package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int = 1
	xx := float64(x)
	fmt.Printf("%T %v %f\n", xx, xx, xx) //float64 1 1.0000

	var y float64 = 1.2
	yy := int(y)
	fmt.Printf("%T %v %d\n", yy, yy, yy) //int 1 1

	//文字列の14を数値に変換
	var s string = "14"
	// z = int(s) // stringのsをint型に変換するのは←左の記述だとエラーになる
	//strconv.Atoiで文字列を整数に変換できる。返り値にintとerrorを返すので、_でerrorを受け取る。
	i, _ := strconv.Atoi(s)
	fmt.Printf("%T %v\n", i, i) //int 14

	h := "Hello World"
	fmt.Println(string(h[0]))
}
