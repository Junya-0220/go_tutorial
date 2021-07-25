package main

import "fmt"

//init関数は一番初めに呼び出される関数main関数より先に呼ばれる
func init() {
	fmt.Println("init")
}

//関数の定義ができる
func bazz() {
	fmt.Println("Bazz")
}

//goはfunc mainが無いとエラーになる
//PrintlnのPは大文字になっている

/*複数行にまたがるコメントアウトの仕方はこのように書きます*/
func main() {
	bazz()
	fmt.Println("Hello world", "test test")
}
