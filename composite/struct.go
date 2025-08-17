package composite

import "fmt"

// 構造体の基本的な使用例
func StructBasic() {
	fmt.Println("=== Struct Basic Example ===")

	type person struct {
		name string
		age  int
		pet  string
	}

	julia := person{
		"ジュリア", // name
		40,     // age
		"cat",  // pet
	}
	beth := person{
		age:  30,
		name: "ベス",
	}
	fmt.Println(julia, beth) // {ジュリア 40 cat} {ベス 30 dog}
}

// 無名構造体の比較と代入の例
func StructAnonymous() {
	fmt.Println("=== Anonymous Struct Example ===")

	type firstPerson struct {
		name string
		age  int
	}
	f := firstPerson{
		name: "Bob",
		age:  50,
	}
	var g struct {
		name string
		age  int
	}
	g = f               // 無名構造体は代入可能
	fmt.Println(f == g) // true
}
