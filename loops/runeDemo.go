package loops

import (
	"fmt"
	"unicode/utf8"
)

func RuneDemo() {
	s := "Go言語😀"
	fmt.Printf("String: %s\n", s)
	fmt.Printf("string: %q\n", s)
	fmt.Printf("bytes(len)=%d, runeCount=%d\n", len(s), utf8.RuneCountInString(s))

	fmt.Println("Indexing bytes:")
	for i := 0; i < len(s); i++ {
		fmt.Printf("  i=%d byte=0x%X\n", i, s[i])
	}

	fmt.Println("Range (runes):")
	for pos, r := range s {
		fmt.Printf(" pos=%d rune=%#U\n", pos, r)
	}

	// Manual rune slice
	rs := []rune(s)
	fmt.Println("As []rune:", rs)
	for i, r := range rs {
		fmt.Printf("  runeIndex=%d %#U\n", i, r)
	}
}

func ReverseStringRunes(s string) string {
	rs := []rune(s)
	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}

func ReverseStringBytes(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func ReverseStringDemo() {
	s := "日本語😀ABC"
	fmt.Println("Original:", s)
	fmt.Println("Reverse (Runes):", ReverseStringRunes(s))
	fmt.Println("Reverse (Bytes):", ReverseStringBytes(s))
}
