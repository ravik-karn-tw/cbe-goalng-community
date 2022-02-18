package code_samples

import (
	"fmt"
	"time"
)

func SwitchExample() {
	day := time.Now().Day()
	switch day {
	case 5, 10, 15:
		fmt.Println("Clean your house.")
	case 25, 26, 27:
		fmt.Println("Buy some food.")
	case 31:
		fmt.Println("Party tonight.")
	default:
		fmt.Println("No info available for the day.")
	}
}
