package vars

import "fmt"

// 変数シャドーと注意点
func Shadow() {
	x := 10

	if x > 5 {
		x := 20
		fmt.Println("Inner x:", x)
	}

	for x := 0; x < 3; x++ {
		fmt.Println("Loop x:", x)
	}
}
