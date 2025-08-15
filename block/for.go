package main

import "fmt"

func ex1() {
	// 標準のfor文
	fmt.Println("=== for Example 1 ===")
	// 変数の初期化には、必ず := を使います。varは使えません。
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func ex2() {
	fmt.Println("=== for Example 2 ===")
	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}
}

func ex3() {
	fmt.Println("=== for Example 3 ===")
	for i := 0; i < 10; {
		fmt.Println(i)
		if i%2 == 0 {
			i++
		} else {
			i += 2
		}
	}
}

func ex4() {
	fmt.Println("=== for Example 4 ===")
	// 他の言語で言うところのwhile
	i := 1
	for i < 100 {
		fmt.Println(i)
		i = i * 2
	}
}

func ex5() {
	// 無限ループ
	// for {
	// 	fmt.Println("Hello")
	// }
}

func ex6() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "3でも5でも割り切れる")
			continue
		}
		if i%3 == 0 {
			fmt.Println(i, "3で割り切れる")
			continue
		}
		if i%5 == 0 {
			fmt.Println(i, "5で割り切れる")
			continue
		}
		fmt.Println(i)
	}
}

func ex7() {
	// for-range
	fmt.Println("=== for Example 7 ===")

	evenVals := []int{2, 4, 6, 8, 10, 12} //偶数
	for i, v := range evenVals {
		fmt.Println(i, v)
	}
}

func ex8() {
	// for-range
	fmt.Println("=== for Example 8 ===")

	evenVals := []int{2, 4, 6, 8, 10, 12} //偶数
	for _, v := range evenVals {
		fmt.Println(v)
	}
}

func ex9() {
	fmt.Println("=== for Example 9 ===")
	uniqueNames := map[string]bool{"花子": true, "太郎": true, "洋子": true}
	for k := range uniqueNames {
		fmt.Println(k)
	}
}

func ex10() {
	// マップのキーのイテレーションは毎回同じと言うわけではない
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	for i := 0; i < 3; i++ {
		fmt.Println("ループ", i)
		for k, v := range m {
			fmt.Println(k, v)
		}
	}
}

func ex11() {
	samples := []string{"hello", "apple_π!", "これは漢字文字列"}
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}
		fmt.Println()
	}
}

func ex12() {
	// 値を変更してもソースは変更されない
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for _, v := range evenVals {
		v *= 2
	}
	fmt.Println(evenVals) // [2 4 6 8 10 12]
}

func ex13() {
	//ラベルの利用
	fmt.Println("=== for Example 13 ===")
	samples := []string{"hello", "apple", "apple_π!", "これは漢字文字列"}
outer:
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
			if r == 'l' || r == 'は' {
				continue outer
			}
		}
	}
}

func main() {
	ex1()
	ex2()
	ex3()
	ex4()
	ex5()
	ex6()
	ex7()
	ex8()
	ex9()
	ex10()
	ex11()
	ex12()
	ex13()
}
