package main

import (
	"encoding/json"
	"fmt"
)

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

// ポインタは最終の手段
// 関数にポインタを渡して、構造体の中身を埋めるのではなく、関数が構造体インスタンスを生成して返すようにします
// bad example
func MakeFoo(f *Foo) error {
	f.Field1 = "val"
	f.Field2 = 20
	return nil
}

// good example
func MakeFoo() (Foo, error) {
	f := Foo{
		Field1: "val",
		Field2: 20,
	}
	return f, nil
}

// 変数を変更するのにポインタ引数を使わなくてはならないのは関数がインターフェースを受け取るときだけです。
func ex0609b() {
	f := struct {
		Name string // NameのNは大文字！小文字だと他パッケージからは見えない
		Age  int
	}{}

	err := json.Unmarshal([]byte(`{"name": "小野小町", "occupation": "歌人", "age": 20}`), &f)
	// 大文字小文字の違いを無視して、フィールドに対応付けてくれる
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v", f) // {Name:小野小町 Age:20}
}

func main() {
	ex0601()
	ex0602()
	ex0603()
	ex0604()
	ex0606()
}
