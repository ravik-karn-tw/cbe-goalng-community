package code_samples

func main() {
	var i int
	var j, k int
	var a, b int = 1, 2
	c := 3
	d, e := 2, "Two"

	var intArray1 [5]int
	intArray2 := [5]int{10, 20, 30, 40, 50}
	var intSlice1 []int
	var intSlice2 = []int{10, 20, 30}
	var intSlice3 = make([]int, 10)

	var emp1 = map[string]int{}
	var emp2 = map[string]int{"Mark":10,"Sandy":20}
	var emp3 = make(map[string]int, 10)
	emp3["Mark"] = 10
	emp3["Sandy"] = 20
}
