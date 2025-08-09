package loops

import "fmt"

func ReverseInPlace(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func ReverseDemo() {
	a := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("before:", a)
	ReverseInPlace(a)
	fmt.Println("after:", a)
}

func AltReverse(a []int) {
	i, j := 0, len(a)-1
	for i < j {
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
}
