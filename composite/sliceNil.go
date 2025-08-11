package composite

import "fmt"

func SliceNil() {
	var x []int                        //xはnilに初期化される
	fmt.Println("x:", x)               // nilスライスの表示
	fmt.Println("x == nil:", x == nil) // nilチェック
}
