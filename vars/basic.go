package vars

import "fmt"

func Basic() {
	var a int = 10
	var b = 20 // 型推論

	var c, d, s int // ゼロ値
	var str string  // ゼロ値 ""
	var flag bool   // ゼロ値 false

	fmt.Println("Basic:")
	fmt.Println(" a=", a, " b=", b, " c=", c, " d=", d, " s=", s)
	fmt.Printf(" str=%q flag=%v\n", str, flag)
}
