package main

import "fmt"

func main() {
	var operation string
	var number1, number2 int

	fmt.Println("Calculator GO 1.0!!")
	fmt.Println("====================")

	fmt.Println("Which operation do you want to perform? (add, subtract, divide, multiply)")

	fmt.Scanf("%s", &operation)
	fmt.Println("Enter first number")
	fmt.Scanf("%d", &number1)

	fmt.Println("Enter second number")
	fmt.Scanf("%d", &number2)

	fmt.Println("====================")

	switch operation {
	case "add":
		fmt.Println(number1 + number2)
	case "subtract":
		fmt.Println(number1 - number2)
	case "multiply":
		fmt.Println(number1 * number2)
	case "division":
		fmt.Println(number1 / number2)
	default:
		return
	}
}

