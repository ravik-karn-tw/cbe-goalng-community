package code_samples

type Person struct {
	Name string
	Age  int
}

func (person Person) CanVote() bool {
	return person.Age >= 18
}

func (person *Person) IncrementAge() {
	person.Age = person.Age + 1
}

func main() {
	var p1 Person
	p1.Name = "ABC"
	p1.Age = 15
	p2 := Person{Name: "Mr. Fred", Age: 45}
	p3 := Person{"Mr. Fred", 45}
	p4 := &Person{Name: "Mr. Fred", Age: 45}
	p5 := &Person{"Mr. Fred", 45}
	p1.CanVote()
	p1.IncrementAge()
}
