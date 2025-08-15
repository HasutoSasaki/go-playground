package main

import "fmt"

func ex1() {
	fmt.Println("=== Shadowing Example 1 ===")
	x := 10
	if x > 5 {
		fmt.Println(x)
		x := 5
		fmt.Println(x)
	}
	fmt.Println(x)
}

func ex2() {
	//複数変数への代入シャドーイング
	fmt.Println("=== Shadowing Example 2 ===")
	x := 10
	if x > 5 {
		a, x := 5, 20
		fmt.Println(a, x) // 5 20
	}
	fmt.Println(x) // 10
}

func ex3() {
	fmt.Println("=== Shadowing Example 3 ===")
	// パッケージ名のシャドーイング ❌
	x := 10
	fmt.Println(x)
	// fmt := "おっと〜"
	// fmt.Println(fmt)
}

func ex4() {
	fmt.Println(true)
	true := 10
	fmt.Println(true)
}

func main() {
	ex1()
	ex2()
	ex3()
	ex4()
}
