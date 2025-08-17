package composite

import "fmt"

// copy: オリジナルとはメモリを共有しない独立したスライスを生成する
func SliceCopy() {
	fmt.Println("=== Slice Copy Example ===")

	// == basic ==
	x := []int{1, 2, 3, 4}
	y := make([]int, 4)
	num := copy(y, x)   // yにxの内容をコピー
	fmt.Println(y, num) // [1 2 3 4] 4
}
