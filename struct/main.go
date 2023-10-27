package main

import (
	"fmt"
	"time"

	"frontendmasters.com/go/struct/data"
)

func main()  {

	// you can declare one or more properties in the constructor
	max := data.Instructor { Id: 3, LastName: "World"} // kind of like new Instuctor()
	
	max.FirstName = "Max"

	println(max.PrintName())

	kyle := data.NewInstructor("Kyle", "Montague")

	println(kyle.PrintName())

	goCourse := data.Course{Id: 2, Name: "Go lang", Instructor: max}

	fmt.Printf("%v", goCourse)

	// when setting the values tho, you can set the values directly in the constructor
	// however you can access all the values directly
	swiftWs := data.Workshop{ Course: data.Course{Name: "Swift Course", Instructor: max}, Date: time.Now() }


	fmt.Printf("%v", swiftWs)

	swiftWithFactory := data.NewWorkshop("Go plus Swift", max)

	fmt.Printf("%v", swiftWithFactory)

	// INTERFACES

	// interfaces are a list of methods which can be used as a type

	
	// doing this will not work because even tho swiftWs is of type WOrkshop which has Course in it
	// it still cannot be used with arrays
	// var courses [2]data.Course

	// courses[0] = goCourse
	// courses[1] = swiftWs

	// to get around that problem you can use interfaces
	// this implements polymorphism in a way
	var courses [2]data.Signable

	courses[0] = goCourse
	courses[1] = swiftWs

	for _, course := range courses {
		fmt.Println(course)
	}

}