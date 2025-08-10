package vars

import "fmt"

func ShortDecl() {
	x := 42
	y := "hello"      // 型推論 string
	x, z := 100, 3.14 // x は再代入, z は新規 (少なくとも1つ新規が必要)
	fmt.Println("ShortDecl:", x, y, z)
}
