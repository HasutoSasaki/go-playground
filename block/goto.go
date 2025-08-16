package main

import (
	"fmt"
	"math/rand"
)

func ex1() {
	//gotoの活用
	a := rand.Intn(10)
	for a < 100 {
		fmt.Println(a)
		if a%5 == 0 {
			goto done
		}
		a = a*2 + 1
	}
	fmt.Println("ループが通常終了した時に行う処理を実行")
done:
	fmt.Println("ループが終わった時に必ず行う処理を実行")
	fmt.Println(a)
}

func main() {
	ex1()
}
