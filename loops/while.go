package loops

import "fmt"

func While() {
	i := 0
	for i < 5 {
		fmt.Println("While:", i)
		i++
	}
	fmt.Println("While Loop completed.")
}
