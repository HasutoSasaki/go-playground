package vars

import "fmt"

// 型変換 / rune / byte / 数値
func Conversion() {
	fmt.Println("Conversion:")
	var i int = 65
	b := byte(i)    // int から byte へ変換
	r := rune(i)    // int から rune へ変換
	f := float64(i) // int から float64 へ変換
	s := string(r)  // rune から string へ変換

	fmt.Printf("i=%d b=%d r=%c f=%f s=%s\n", i, b, r, f, s)
}
