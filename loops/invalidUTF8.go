package loops

import "fmt"

func InvalidUTF8() {
	s := "日本\x80\x90\x7F語" // \x80\x90 は不正な単独バイト
	for pos, r := range s {
		fmt.Printf("character %#U starts at byte position %d\n", r, pos)
	}
}
