package code_samples

import "fmt"

func LoopMap() {
	var employee = map[string]int{"Mark": 10, "Sandy": 20}
	for name, score := range employee {
		fmt.Println("Name:", name, "=>", "Score:", score)
	}
}
