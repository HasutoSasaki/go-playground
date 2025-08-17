package composite

import "fmt"

// composite型の練習問題
func Practice() {
	fmt.Println("=== Practice Problems ===")

	// 1. スライスの操作
	fmt.Println("--- Problem 1: Slice Operations ---")
	greeting := make([]string, 0, 5)
	greeting = append(greeting, "Hello", "Hola", "नमस्कार", "こんにちは", "Привіт")
	fmt.Println(greeting)
	greeting1 := greeting[:2]
	fmt.Println(greeting1)
	greeting2 := greeting[1:4]
	fmt.Println(greeting2)
	greeting3 := greeting[3:]
	fmt.Println(greeting3)

	// 2. Runeの操作
	fmt.Println("--- Problem 2: Rune Operations ---")
	message := string("Hi 👩 and 👨")
	runes := []rune(message)[3] // 4番目のruneを取得
	fmt.Println(string(runes))  // 👩

	// 3. 構造体の作成と比較
	fmt.Println("--- Problem 3: Struct Creation ---")
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	Employee1 := Employee{
		firstName: "John",
		lastName:  "Doe",
		id:        1,
	}
	Employee2 := Employee{
		firstName: "John",
		lastName:  "Doe",
		id:        1,
	}
	var emp Employee
	emp.firstName = "John"
	emp.lastName = "Doe"
	emp.id = 1

	fmt.Println(Employee1, Employee2, emp) // {John Doe 1} {John Doe 1} {John Doe 1}
}
