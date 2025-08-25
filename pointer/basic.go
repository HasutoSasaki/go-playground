package main

import "fmt"

func ex0601() {
	x := 10
	pointerToX := &x
	fmt.Println(pointerToX)  // アドレスが表示される
	fmt.Println(*pointerToX) // 10 // デリファレンスする
	z := 5 + *pointerToX
	fmt.Println(z)

	// nilポインタをデリファレンスすると「パニック」になる
	var y *int
	fmt.Println(y == nil) // true
	// fmt.Println(*y) // パニックになる！

	var a = new(int)      // aの参照先にはintのゼロ値(0)が記憶される
	fmt.Println(a == nil) // false
	fmt.Println(a)        // 0x1400000e130 (例: xのアドレスが表示される)
	fmt.Println(*a == 0)  // true
	fmt.Println(*a)       // 0

	// 関数newはほとんど使われず、構造体については、構造体リテラルの前いに「&」をつけてポインタのインスタンスを作成。
	b := &Foo{} //「Foo{}」は構造体リテラル
	var c string
	d := &c // stringへのポインタ
}

func main() {
	ex0601()

}
