package composite

import "fmt"

func SliceAppendBasic() {
	var x []int
	x = append(x, 1)
	x = append(x, 2, 3)
	fmt.Println("x:", x)
}
