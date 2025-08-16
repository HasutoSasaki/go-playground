package main

import (
	"fmt"
	"unicode/utf8"
)

func ex1() {
	fmt.Println("=== for Example 1 ===")
	words := []string{"山", "sun", "微笑み", "人類学者", "モグラの穴", "mountain", "タコの足とイカの足", "antholopologist", "タコの足は8本でイカの足は10本"}
	for _, word := range words {
		switch rc := utf8.RuneCountInString(word); rc { // Rune Count
		case 1, 2, 3, 4:
			fmt.Printf("「%s」の文字列は%dで、短い単語だ。\n", word, rc)
		case 5:
			bc := len(word)
			fmt.Printf("「%s」の文字列は%dで、これはちょうど良い長さだ。ちなみにバイト数は%dだ。\n", word, rc, bc)
		case 6, 7, 8, 9:
		default:
			fmt.Printf("「%s」の文字列数は%dで、とても長い!\n", word, rc)
		}
	}
}

func ex2() {
	fmt.Println("=== for Example 2 ===")
	// ブランク switch
	words := []string{"hi", "salutations", "hello"}
	for _, word := range words {
		switch wordLen := len(word); {
		case wordLen < 5:
			fmt.Println(word, "は短い単語です")
		case wordLen > 10:
			fmt.Println(word, "は長すぎる単語です")
		default:
			fmt.Println(word, "はちょうどよい長さの単語です")
		}
	}
}

func ex3() {
	fmt.Println("=== for Example 3 ===")
	// breakで抜ける方法
loop:
	for i := 0; i < 10; i++ {
		switch {
		case i%2 == 0:
			fmt.Println(i, ":偶数")
		case i%3 == 0:
			fmt.Println(i, ":3で割り切れるが2では割り切れない")
		case i%7 == 0:
			fmt.Println(i, ":ループは終了する！")
			break loop //breakでラベルを指定したら抜けることができる
			//しかし、そもそも、breakが必要というのは複雑すぎる可能性があるので設計を見直す
		default:
			fmt.Println(i, "：退屈な数")
		}
	}
}

func main() {
	ex1()
	ex2()
	ex3()
}
