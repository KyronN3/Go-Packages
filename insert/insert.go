package insert

import "fmt"

func PushByIndex(elements int, index int, num ...int) []int {
	store := []int{}
	store = append(store, num...)
	return append(store[:index], append([]int{elements}, store[index:]...)...)
}

func Input() (*int, *int) {
	var index, elements int

	fmt.Print("Index: ")
	fmt.Scan(&index)

	fmt.Print("Elements: ")
	fmt.Scan(&elements)

	return &index, &elements
}
