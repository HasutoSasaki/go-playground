package composite

import "fmt"

func Array() {
	var x = [...]int{10, 20, 30}
	var y [2][3]int
	y[0] = [3]int{1, 2, 3}
	y[1] = [3]int{4, 5, 6}
	fmt.Println("Composite array:", x)
	fmt.Println("2D array:", y)
}
