package vars

import "fmt"

// 多値返却と破棄
func MultiAssign() {
	fmt.Println("MultiAssign:")
	divRem := func(a, b int) (int, int) {
		return a / b, a % b
	}

	q, r := divRem(10, 3)
	_, onlyR := divRem(10, 3)
	fmt.Printf("q=%d r=%d onlyR=%d\n", q, r, onlyR)
}
