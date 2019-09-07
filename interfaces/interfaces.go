package main

// interfaces are implicit
// structs don't need to say they are implementing it
// they just implement it

// this is convenient because you don't need to import the interface that
// you implement.

// This also allows you to create a new interface and then any struct
// that already has functions defined by the interface would automatically
// implement it.

// Interfaces are also very convient when unit testing
// they allow you to easily mock out values (more on this later)

// Vehicle interface
type Vehicle interface {
	Designation() string
	Locomotion() string
	Sentient() bool
}

type Character interface {
	Name() string
}

// ThinkTank struct
type ThinkTank struct {
	Model string
}

// Designation method
func (tt *ThinkTank) Designation() string {
	return "ThinkTank Model: " + tt.Model
}

// Locomotion method
func (tt *ThinkTank) Locomotion() string {
	return "ThinkTanks walk using legs"
}

// Sentient method
func (tt *ThinkTank) Sentient() bool {
	return true
}

// Name method
func (tt *ThinkTank) Name() string {
	return tt.Model
}

// Say is not an interface method
func (tt *ThinkTank) Say() {
	println("Hey Team!")
}

func main() {
	tt := &ThinkTank{"Fuchikoma"}
	tt.Say()

	// take a Vehicle interface and calls its methods
	VehicleReport(tt)

	// think tank also implements the character interface
	println(tt.Name())

	// when dealing with interfaces, it's common you'd like to take different
	// actions depending on the interface the object implements
	// we can use a switch

	// type switch assertions
	whatType(tt)
}

// VehicleReport printer
func VehicleReport(v Vehicle) {
	println(v.Designation())
	println(v.Locomotion())
	println("Sentient?", v.Sentient())
	// v.Name() //error
	//v.Say() //error
}

func whatType(tt interface{}) {
	switch tt.(type) {
	case Vehicle:
		println("This is a vehicle")
	case Character:
		println("This is a character")
	default:
		println("I'm not sure what the type is")
	}
}
