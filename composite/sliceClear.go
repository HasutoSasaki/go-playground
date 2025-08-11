package composite

import "fmt"

func SliceClear() {
	s := []string{"first", "second", "third"}
	fmt.Println(s, len(s), cap(s)) // [first second third] 3 3
	clear(s)
	fmt.Println(s, len(s), cap(s)) // [   ] 3 3 (注: 実際にはlen,capは変わらない)
}
