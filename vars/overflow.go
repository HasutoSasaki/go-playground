package vars

import (
	"fmt"
	"math"
)

// Overflow は整数型 (byte, int32, uint64) の最大値から +1 した際の
// オーバーフロー挙動を表示するデモ。
func Overflow() {
	var b byte = math.MaxUint8       // 255
	var smallI int32 = math.MaxInt32 // 2147483647
	var bigI uint64 = math.MaxUint64 // 18446744073709551615

	b = b + 1           // 0 に戻る
	smallI = smallI + 1 // -2147483648 に巻き戻る (2の補数)
	bigI = bigI + 1     // 0 に戻る

	fmt.Println(b)
	fmt.Println(smallI)
	fmt.Println(bigI)
}
