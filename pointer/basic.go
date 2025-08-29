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

type person struct {
	FirstName  string
	MiddleName *string
	LastName   string
}

func ex0602() {

	p := person{
		FirstName:  "Pat",
		MiddleName: "Perry", // ←コンパイル時のエラー
		LastName:   "Peterson",
	}
	fmt.Println(p)
}

func ex0603() {

	s := "Perry"
	p := person{
		FirstName:  "Pat",
		MiddleName: &s, // 回避
		LastName:   "Peterson",
	}
	fmt.Println(p)
}

func makePointer[T any](t T) *T { // ヘルパー関数
	return &t
}

func ex0604() {
	p := person{
		FirstName:  "Pat",
		MiddleName: makePointer("Perry"), //これならうまくいく
		LastName:   "Peterson",
	}
	fmt.Println(p)             // {Pat 0xc000010250 Peterson}
	fmt.Println(*p.MiddleName) // Perry
}

func failedUpdate(g *int) {
	x := 10
	g = &x
}

func ex0606() {
	var f *int // fはnil
	failedUpdate(f)
	fmt.Println(f) // <nil>
}

func failedUpdate2(px *int) {
	x2 := 20
	px = &x2
}

func update(px *int) {
	*px = 20
}

func ex0607() {
	x := 10
	failedUpdate2(&x)
	fmt.Println(x) // 10
	update(&x)
	fmt.Println(x) // 20

}

func main() {
	ex0601()
	ex0602()
	ex0603()
	ex0604()
	ex0606()
}
