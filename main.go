package main

import (
	"fmt"
	"go-module/insert"
)

func Print(print ...interface{}) {
	fmt.Print(print...)
}

func main() {

	data := []int{1, 2, 3, 4, 6}

	index, elements := insert.Input()
	Print(insert.PushByIndex(*elements, *index, data...))

}
