package data

// this creates an alias for the int type - same type, different name
type integer = int

var x integer

// this creates a new type entirely - should not be confused with type aliases
type json map[string]string

type distance float32
type distanceKm float64

// Methods
// types in go can have methods
// these methods can be called on the types directly
// methods are created by putting brackets before the func and function name
func (miles distance) ToKm() distanceKm {
	return distanceKm(miles * 1.69034)
}

// you can still pass normal arguments into the func, however the concerned type must be declared in between func and function name
func (miles distanceKm) ToMiles() distance {
	return distance(miles / 1.69034)
}



func typeTest() {
	d := distance(34.4)

	km := d.ToKm()

	km.ToMiles()

	print(d)

}

// STRUCTS

// they kind of act like classes  and they are used to group properties together
// they have constructor and methods
type User struct {
	id int
	name string
}

func structTest() {
	var u1 User
	// can be set like key values
	u1 = User {id: 1, name: "hello"}
	// or set by values alone
	u2 := User {2, "world"}

	println(u1, u2)
}

// structs can have methods
// they are declared like normal methods above
func (u User) PrettyPrint() string {
	return string(u.id) + ": " + u.name 
}


// INTERFACES
// A definition of methods, can emulate polymorphism, implicit implementation
