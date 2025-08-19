package main

import (
	"fmt"
	"strconv"
)

// 同じシグネチャを持った一連の関数を作成
func _add(i int, j int) int { return i + j }
func _sub(i int, j int) int { return i - j }
func _mul(i int, j int) int { return i * j }
func _div(i int, j int) int { return i / j }

var opMap = map[string]func(int, int) int{ // 「文字列→関数」のマップ
	"+": _add,
	"-": _sub,
	"*": _mul,
	"/": _div,
}

func main() {
	expressions := [][]string{
		[]string{"2", "+", "3"},
		[]string{"2", "-", "3"},
		[]string{"2", "*", "3"},
		[]string{"2", "/", "3"},
		[]string{"2", "%", "3"},
		[]string{"two", "+", "three"},
		[]string{"2", "+", "three"},
		[]string{"5"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 { // 演算子と被演算子の合計数のチェック
			fmt.Print(expression, " --  不正な式です\n")
			continue
		}
		p1, err := strconv.Atoi(expression[0]) // 1番目の被演算子 (oPerand)のチェック
		if err != nil {
			fmt.Print(expression, " -- ", err, "\n")
			continue
		}
		op := expression[1] // 演算子 (OPerator)のチェック
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Print(expression, " -- ", "定義されていない演算子です：", op, "\n")
			continue
		}
		p2, err := strconv.Atoi(expression[2]) // 2番目の演算子チェック
		if err != nil {
			fmt.Print(expression, " -- ", err, "\n")
			continue
		}
		result := opFunc(p1, p2) //実際の計算
		fmt.Print(expression, " -> ", result, "\n")
	}
}
