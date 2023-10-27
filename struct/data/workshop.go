package data

import "time"

// EMBEDDING

// extends does not exist in go, so instead we embed other struct into structs
// all the values in the other struct is copied into this one
type Workshop struct {
	Course
	Date time.Time
}

func (w Workshop) SignUp() bool {
	return true
}

func NewWorkshop(name string, instructor Instructor) Workshop  {
	w := Workshop {}

	w.Name = name
	w.Instructor = instructor

	// values can be overriden
	// in that case you can access the name via the embedded struct's name
	// e.g w.Course.Name if the workshop also has a Name value

	return w
}