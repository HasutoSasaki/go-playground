package main

import (
	"errors"
	"fmt"
	"os"
)

func div(num int, denom int) int { //分子(numerator), 分母(denominator)
	if denom == 0 {
		return 0
	}
	return num / denom
}

// ex2
type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

func MyFunc(opts MyFuncOpts) error {
	fmt.Println(opts)
	fmt.Println("＜ここで必要な処理を行う＞")
	return nil
}

// ex3 可変長引数とスライス
func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

// ex4 複数の戻り値
func divAndRemainder(num, denom int) (int, int, error) {
	if denom == 0 {
		return 0, 0, errors.New("0で割ることはできません")
	}
	return num / denom, num % denom, nil
}

// ex5 名前付き戻り値
func divAndRemainder2(num, denom int) (result int, remainder int, err error) {
	if denom == 0 {
		return num, denom, errors.New("0で割ることはできません")
	}
	result, remainder = num/denom, num%denom
	return result, remainder, err
}

// ex6 適当な値を代入
func divAndRemainder3(num, denom int) (result int, remainder int, err error) {
	result, remainder = 20, 30 // 適当な値を代入
	if denom == 0 {
		return num, denom, errors.New("0で割ることはできません")
	}
	return num / denom, num % denom, nil
}

// ex7 ブランク return
// 関数が値を返すなら、絶対にブランクreturn を使ってはなりません。実際に返される値が非常にわかりにくくなる。
func dicAndRemainder4(num int, denom int) (result int, remainder int, err error) {
	if denom == 0 {
		err = errors.New("0で割ることはできません")
		return
	}
	result, remainder = num/denom, num%denom
	return
}

func main() {
	fmt.Println("=== ex1 ===")
	fmt.Println(div(4, 2)) // 2
	fmt.Println(div(4, 0)) // 0

	fmt.Println("=== ex2 ===")
	// 名前月引数やオプション引数は関数への入力が多い時に役立つものなので、使いたいと思う時には、関数が複雑すぎる可能性があります
	MyFunc(MyFuncOpts{
		LastName: "Patel",
		Age:      50,
	})
	MyFunc(MyFuncOpts{
		FirstName: "Joe",
		LastName:  "Smith",
	})

	fmt.Println("=== ex3 ===")
	fmt.Println(addTo(3))
	fmt.Println(addTo(3, 2))
	fmt.Println(addTo(3, 2, 4, 6, 8))
	a := []int{4, 3}
	fmt.Println(addTo(3, a...)) // [7 6]
	fmt.Println(addTo(3, []int{1, 2, 3, 4, 5}...))

	fmt.Println("=== ex4 ===")
	// GOでは関数の返り値はそれぞれ代入しなければいけいない!
	result, remainder, err := divAndRemainder(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, remainder) // 2 1

}
