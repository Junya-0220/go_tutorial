package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println("Hello " + "World")
	fmt.Println(string("Hello World"[0]))
	var s string = "Hello World"

	//文字列の置換
	//型の指定にはstring,文字列置換の関数はstringsを使うので注意。
	s = strings.Replace(s, "H", "X", 1)
	fmt.Println(s)

	//stringsにはContainsという関数があり、指定した文字列が入っていればtrueを返す。
	fmt.Println(strings.Contains(s, "World"))

	//``(バッククォート)で囲ったところは、スペースを認識する。
	fmt.Println(`Test
                       Test
Test`)

	//エスケープする時に\を使用するパターンと``で囲うパターンがある。
	fmt.Println("\"")
	fmt.Println(`"`)
}
