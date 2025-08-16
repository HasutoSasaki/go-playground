package main

import (
	"fmt"
	"math/rand"
)

func ex1() {
	result := make([]int, 0, 100)
	for len(result) != 100 {
		random := rand.Intn(100)
		result = append(result, random)
	}
	fmt.Println(result)
}

func ex2() {
	fmt.Println("==== ex2 ====")
	result := make([]int, 0, 100)
	for len(result) != 100 {
		random := rand.Intn(100)
		if random%2 == 0 {
			fmt.Println("Two!")
		} else if random%3 == 0 {
			fmt.Println("Three!")
		} else {
			fmt.Println("Never mind")
		}
		result = append(result, random)
	}
	fmt.Println(result)
}

func ex3() {
	fmt.Println("==== ex3 ====")
	// あえて、totalが0になるようになっています。変数のブロックの復習
	var total int = 0
	for i := 0; i < 10; i++ {
		total := total + i
		fmt.Printf("i=%v total=%v\n", i, total)
	}
	fmt.Printf("final total:%v\n", total)
}

func ex3() {
	fmt.Println("==== ex4 ====")
	var total int = 0
	for i := 0; i < 10; i++ {
		total = total + i //:を外して、再代入する形にします。
		fmt.Printf("i=%v total=%v\n", i, total)
	}
	fmt.Printf("final total:%v\n", total)
}

func main() {
	ex1()
	ex2()
	ex3()
}
