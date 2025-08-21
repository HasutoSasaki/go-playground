package main

import "fmt"

func main() {
	fmt.Println("=== ex1 ===")
	a := 20
	f := func() {
		fmt.Println(a)
		a = 30
	}
	f()            // 20 　// 変数fに代入された関数を実行
	fmt.Println(a) // 30

	// fmt.Println("=== ex2 ===")
	// 	a := 20
	// f := func() {
	// 	fmt.Println(a)
	// 	a := 30
	// 	fmt.Println(a) // 30
	// }
	// f()
	// fmt.Println(a) // 20
}
