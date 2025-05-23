package main

import (
	"fmt"
	// "go-module/insert" //uncomment the package to use
	// "go-module/object"
)

func Print(print ...interface{}) {
	fmt.Print(print...)
}

func main() {

	//example insert

	/* insertData := []int{1, 2, 3, 4, 6}

	index, elements := insert.Input()
	Print(insert.PushByIndex(*elements, *index, insertData...)) */



	
	//example object
	/* objectOne := object.ObjectMap([]any{"FirstName", "LastName", "Age"}, []any{"James Kyron", "Sostino", "22"})
	objectTwo := object.ObjectMap([]any{"FirstName", "LastName", "Age"}, []any{"Trash", "Bin", "30"})

	data := object.ObjectMapList(objectOne, objectTwo)

	for i := range data {
		for _, elements := range []any{"FirstName", "LastName", "Age"} {
			Print(elements, ": ", data[i][elements], "\n")
		}
	} */

}
