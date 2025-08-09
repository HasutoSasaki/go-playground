package loops

import "fmt"

func Infinite() {
	count := 0
	for {
		fmt.Println("Infinite:", count)
		count++
		if count >= 5 {
			break
		}
	}
	fmt.Println("Infinite Loop completed.")
}
