package main

import "fmt"

var packageVar = 40

func printData() {
	var message string = "Hello from go "
	price := 34.4

	fmt.Print(message, price)

	fmt.Println("Trying out fmt package")
	fmt.Println(30.00)


	priceAsInt := int(price)

	fmt.Println(priceAsInt)
}