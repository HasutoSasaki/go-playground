package composite

import "fmt"

func SliceMake() {
	x := make([]int, 5)
	fmt.Println(x, len(x), cap(x)) // [0 0 0 0 0] 5 5
	// makeを使って生成したスライスの要素に、appendするのはNG❌
	x = append(x, 10)              // [0 0 0 0 0 10]
	fmt.Println(x, len(x), cap(x)) // [0 0 0 0 0 10] 6 10
	// appendは常にスライスの長さを増やす。
}
