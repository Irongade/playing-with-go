package data

import "fmt"

type Instructor struct {
	// if you declare variables in struct with lowercase first letter, they becom private
	// and cannot be accessed outside

	// to make them available outside, make the first character capitalized
	Id int
	FirstName string
	LastName string
	Score int
}

func (i Instructor) PrintName() string {
	return fmt.Sprintf("%v, %v", i.FirstName, i.LastName)
}

// FACTORIES
// factories are like static methods of a class
// they can only be called by Structs
// this is a factory function and the name of the function has to start with "New"
func NewInstructor(name string, lastName string) Instructor {
	return Instructor{FirstName: name, LastName: lastName}
}