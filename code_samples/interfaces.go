package code_samples

import "fmt"

type Animal interface {
	Sound() string
}

type Dog struct{}

func (d Dog) Sound() string {
	return "Woof!"
}

type Cat struct{}

func (c Cat) Sound() string {
	return "Meow!"
}

func main() {
	animals := []Animal{Dog{}, Cat{}}
	for _, animal := range animals {
		fmt.Println(animal.Sound())
	}
}
