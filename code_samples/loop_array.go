package code_samples

import "fmt"

func LoopArray() {
	intArray := [5]int{10, 20, 30, 40, 50}
	for index, element := range intArray {
		fmt.Println(index, "=>", element)
	}
}
