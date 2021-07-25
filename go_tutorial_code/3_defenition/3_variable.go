package main

import "fmt"

//varの変数宣言はどこでもできる。
//var()でまとめて変数宣言できる。
var (
	i    int     = 1
	f64  float64 = 1.2
	s    string  = "test"
	t, f bool    = true, false
)

// ○○ := 1のような変数宣言は関数の中だけで使用ができる。
// 明確な型宣言をする場合はvar,簡略して宣言する場合は:を使う。
// 再代入するときは:は使わない。
func foo() {
	xi := 1
	xi = 2
	// xf64 := 1.2
	var x32 float32 = 1.2
	xfs := "test"
	xt, xf := true, false
	fmt.Println(xi, xfs, xt, xf)
	//xf64の型をチェックできる
	fmt.Printf("%T\n", x32)
	fmt.Printf("%T\n", xi)
}

func main() {
	fmt.Println(i, f64, s, t, f)
	foo()
}
