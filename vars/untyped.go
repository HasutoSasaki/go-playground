package vars

import "fmt"

func Untyped() {
	const U = 3             // untyped int
	const F = 1.5           // untyped float
	var f64 float64 = U + F // (untyped int + untyped float) -> typed float64
	var i8 int8 = U         // untyped int -> typed int8
	fmt.Println("Untyped:", f64, i8)

	// bad pattern ❌
	// var i32 int32 = U
	// var f32 float32 = F
	// var f64Bad float64 = i32 + f32 // typeが指定されている場合は、明示的な型変換が必要
}
