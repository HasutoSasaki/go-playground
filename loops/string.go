package loops

import "fmt"

func String() {
	s := "Hello, Go!"
	for i, char := range s {
		fmt.Printf("Character %d: %c\n", i, char)
	}
	fmt.Println("String Loop completed.")
}
