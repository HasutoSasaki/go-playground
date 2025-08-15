package main

import (
	"fmt"
	"math/rand"
)

func ex1() {
	fmt.Println("=== If Example 1 ===")
	n := rand.Intn(10)
	if n == 0 {
		fmt.Println("少し小さすぎます：", n)
	} else if n > 5 {
		fmt.Println("大きすぎます：", n)
	} else {
		fmt.Println("いい感じの数字です：", n)
	}
}

func ex2() {
	// if文の変数のスコープ
	if n := rand.Intn(10); n == 0 { // nはelseのブロックまで有効になる
		fmt.Println("少し小さすぎます：", n)
	} else if n > 5 {
		fmt.Println("大きすぎます：", n)
	} else {
		fmt.Println("いい感じの数字です：", n)
	}
	// fmt.Println(n) //エラーになる❌
}

func main() {
	ex1()
	ex2()
}
