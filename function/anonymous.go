package main

import "fmt"

func ex1() {
	// 無名関数を変数に代入するパターン
	f := func(j int) {
		fmt.Println("無名関数の中で", j, "を出力")
	}
	for i := 0; i < 5; i++ {
		f(i)
	}
}

func ex2() {
	// 変数に代入しないパターンの無名関数
	for i := 0; i < 5; i++ {
		func(j int) {
			fmt.Println("無名関数の中で", j, "を出力")
		}(i)
	}
}

var (
	add = func(i int, j int) int { return i + j }
	sub = func(i int, j int) int { return i - j }
	mul = func(i int, j int) int { return i * j }
	div = func(i int, j int) int { return i / j }
)

func ex3() {

	x := add(2, 3)
	fmt.Println(x)
	y := sub(6, 3)
	fmt.Println(y)
	z := mul(2, 3)
	fmt.Println(z)
	a := div(10, 5)
	fmt.Println(a)
}

// パッケージレベルで定義した無名関数を書き換えることができてしまう。
func changeAdd() {
	add = func(i, j int) int { return i + j + j }
}

// パッケージレベルの無名関数を使う前に、本当に必要か検討する。フローを理解しやすいものにするには、パッケージレベルの状態は普遍であるべき。
func ex4() {
	x := add(2, 3)
	fmt.Println(x) //5
	changeAdd()
	y := add(2, 3) //8
	fmt.Println(y)
}

func main() {
	ex1()
	ex2()
	ex3()
	ex4()
}
