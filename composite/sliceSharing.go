package composite

import "fmt"

func SliceSharing() {

	// == 複数のスライスによる要素の共有 ==
	x := []string{"a", "b", "c", "d"}
	y := x[:2]
	z := x[1:]
	x[1] = "y"
	y[0] = "x"
	z[1] = "z"
	fmt.Println("x:", x) // x: [x y z d]
	fmt.Println("y:", y) // y: [x y]
	fmt.Println("z:", z) // z: [y z d]
}
