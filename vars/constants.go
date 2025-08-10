package vars

import "fmt"

func Constants() {
	const Pi = 3.14159           // 型推論 (untyped float constant)
	const TypedPi float64 = 3.14 // 型付き定数

	const (
		Sun = iota // 0
		Mon        // 1
		Tue        // 2
		_          // skip 3
		Thu        // 4
	)

	// ビットシフト列挙 (1,2,4,8,...)
	const (
		_  = iota             // 0
		KB = 1 << (10 * iota) // 1<<10
		MB                    // 1<<20
		GB                    // 1<<30
	)

	fmt.Println("ConstantsDemo:")
	fmt.Println(" Pi=", Pi, " TypedPi=", TypedPi)
	fmt.Println(" Days:", Sun, Mon, Tue, Thu)
	fmt.Println(" Sizes:", KB, MB, GB)
}
