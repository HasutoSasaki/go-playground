package main

import (
	"fmt"
	"sort"
)

func main() {

	type Person struct {
		firstName string
		LastName  string
		Age       int
	}

	people := []Person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobbert", 23},
		{"Fred", "Fredson", 18},
	}

	fmt.Println("=== 初期データ ===")
	fmt.Println(people)

	// 姓(LastName) でソート
	sort.Slice(people, func(i, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	fmt.Println("=== 姓（LastName。2番目のフィールド)でソート")
	fmt.Println(people)

	// 年齢(Age)でソート
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("=== 年齢(Age) でソート")
	fmt.Println(people)

	fmt.Println("ソート後のpeople")
	fmt.Println(people)
}
