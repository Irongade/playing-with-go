package data

import "fmt"

var Countries [10]string
var Slice = []int {1, 2, 3}
var Codes map[int]bool

func init() {
	Countries[0] = "Brazil"
	Countries[1] = "Argentina"
	Countries[2] = "Nigeria"

	names := []string {"Mary", "John"}

	qty := len(Countries)

	names = append(names, "Mathew")

	wellKnownPorts := map[string]int {"http": 80, "https": 443}

	fmt.Println("Countries", qty, len(names), wellKnownPorts)
}

func save(text string) {

}

func add(x int, y int) int {
	return x + y
}

func addAndSubtract(x int, y int) (int, int) {
	return x + y, x-y
}