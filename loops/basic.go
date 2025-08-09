package loops

import "fmt"

// BasicLoop demonstrates a simple for loop in Go.
func Basic() {
	for i := 0; i < 5; i++ {
		fmt.Println("Iteration:", i)
	}
	fmt.Println("Basic Loop completed.")
}
