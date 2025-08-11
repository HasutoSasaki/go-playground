package vars

import "fmt"

func Composite() {
	var arr [3]int = [3]int{1, 2, 3} // array
	s := []int{4, 5, 6}              // slice
	a := make([]int, 0, 5)
	m := map[string]int{"a": 1}

	type Point struct {
		X, Y int
	}
	p := Point{X: 10, Y: 20} // struct

	anon := struct{ Flag bool }{true}
	fmt.Println("Composite:", arr, s, a, m, p, anon)

}
