package loops

import "fmt"

func Slice() {
	nums := []int{10, 20, 30}
	for i, num := range nums {
		fmt.Println("Slice:", i, num)
	}
	fmt.Println("Slice Loop completed.")
}
