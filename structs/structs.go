package main

// User struct
type User struct {
	Name    string
	Age     int
	Deleted bool
}

func main() {

	// struct literal
	jackie := User{"Jackie", 21, false}
	println("Struct Literal: ", jackie.Name)

	// struct with named fields
	// allows you to leave some fields blank
	renee := User{
		Name: "Renee",
		Age:  24,
	}
	println("Struct named field: ", renee.Name)

	// struct pointer
	// "&" creates a pointer to a struct
	// - preferred way of declaring structs is to create a reference
	brian := &User{
		Name: "Brian",
	}

	// "*" gets the pointers underlying value
	println("Pointer: ", (*brian).Name)

	// Since this is very common, Go also allows you to drop the asterick
	println("Implicit pointer: ", brian.Name)

	// set fields like so...
	brian.Age = 36
	println("Setting field: ", brian.Age)
}
