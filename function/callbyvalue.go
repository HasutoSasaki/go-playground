package main

import "fmt"

type person struct {
	age  int
	name string
}

func modifyFails(i int, s string, p person) {
	i *= 2 // 「 i = i * 2」
	s = "さようなら"
	p.name = "Bob"
}

func modMap(m map[int]string) {
	m[2] = "こんにちは"
	m[3] = "さようなら"
	delete(m, 1)
}

func modSlice(s []int) {
	for k, v := range s {
		s[k] = v * 2
	}
	s = append(s, 10)
}

func main() {
	p := person{}
	i := 2
	s := "こんにちは"
	// 関数は渡された引数の値を変えられないということがわかる。
	fmt.Println(i, s, p) // 2 こんにちは {0 }
	modifyFails(i, s, p)
	fmt.Println(i, s, p) // 2 こんにちは {0 }

	// ==== map と slice の場合は引数の値元も変更される ====
	// map の例　==
	m := map[int]string{
		1: "1番目",
		2: "2番目",
	}
	modMap(m)
	fmt.Println(m) // map[2:こんにちは 3:さようなら]

	// slice の例 ==
	x := []int{1, 2, 3}
	modSlice(x)
	fmt.Println(x) // [2,4,5]

	// なぜ通常の関数と異なるかというと、マップとスライスはポインタを使って実装されているためである。

}
