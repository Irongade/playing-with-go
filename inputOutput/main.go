package main

import (
	"fmt"

	"frontendmasters.com/go/io/data"
)


func init()  {
	fmt.Println("A")
}

func init()  {
	fmt.Println("B")
}
func calculateTax(price float32) (float32, float32) {
	return price * 0.09, price * 0.02
}

func calculateTaxWithName(price float32) (stateTax float32, cityTax float32) {
	// creates the return data, and once they are set you do not need to return them explicitly
	stateTax = price * 0.09
	cityTax = price * 0.02

	return
}

// * denotes a pointer to a value itself, when we use it we get the value from the pointer
// & is use to pass in an address or pointer or reference to a variable.
// you can do double pointers e.g. **age -  which is a pointer to a pointer to a value
func birthday(age *int) {
	fmt.Printf("The pointer is %v and the value is %v \n", age, *age)
	*age++
}

// panic is not to be used for error management
// it terminates the current function call, which then bubbles up till it likely closes your app
// of course, there are ways to trap it.
// it is used when app reaches invalid state but you were not expecting an error
func birthdayWithPanic(age *int) {
	if (*age > 140) {
		panic("Too old to be true")
	} 
	fmt.Printf("The pointer is %v and the value is %v \n", age, *age)
	*age++
}

func controlIf()  {
	user :=  true
	// booleans don't need () for ifs
	if user {

	}

	// there are no while and do-while-s

	// the first part is a simple code block, the other one is the main condition
	// the first part creates a variable which exists in the if and else blocks
	// but not outside of the if-else block
	if message := "Hello"; user != false {
		fmt.Println(message)
		// message exists here
	} else {
		// message exists here
		fmt.Println(message)
	}
 	
}

func controlSwitch(day string)  {
	user := "Help"
	switch day {
	case "Monday":
		fmt.Println("Monday")
	case "Saturday":
		// if the case matches this condition, automatically allow the code to fallthrough to the next case
		fallthrough
	case "Sunday":
		fmt.Println("Sunday")
	default:
		fmt.Println("Another working day")

	}

	// switch can exists without any condition
	// the first case that matches executes, so it works like an if-else statement
	switch {
	case user == "nil":
	case user == "false":
	case user == "true":
	default:
	}
}

func controlFor() {
	collection := [3]int {1,2,3}
	colMap :=  map[string]int {"key": 1, "value": 2}
	count := 5

	// classic for
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// similar to for in 
	// the index is the index of the collection
	for index := range collection {
		fmt.Println(index)
	}

	// similar to for each
	for key, value := range colMap {
		fmt.Println(key, value)
	}

	// kind of while loop as well

	for count < 10 {
		count++
	}

	// this is an infinite loop - while loop
	// for {

	// }
 }

func main() {
	print("Hello from Main entrypoint")

	printData()

	// to make functions exportable between packages, use capital case for the Function names

	println(packageVar)

	println(data.MaxSpeed)

	// use to delay execution until the end of the current function call is reached
	// you can have more than one which is stacked (LIFO) - so Bye!! will print before Good
	// if there is a panic - defers will run before the app is closed.
	defer fmt.Println("Good")
	defer fmt.Println("Bye!!!")

	stateTax, _ := calculateTax(100)

	fmt.Println(stateTax)


	age := 22
	birthday(&age)
	fmt.Println(age)

	// controlFor()
}

