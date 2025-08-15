package composite

import "fmt"

func SliceAppend() {
	x := []string{"a", "b", "c", "d"}
	y := x[:2]
	y = append(y, "z")
	y = append(y, "1")
	y = append(y, "2")

	fmt.Println("x:", x) // x: [a b c d]
	fmt.Println("y:", y) // y: [a b z]
}
