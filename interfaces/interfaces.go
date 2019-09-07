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

// Say is not an interface method
func (tt *ThinkTank) Say() {
	println("Hey Team!")
}

func main() {
	tt := &ThinkTank{"Fuchikoma"}
	tt.Say()

	// take a Vehicle interface and calls its methods
	VehicleReport(tt)
}

// VehicleReport printer
func VehicleReport(v Vehicle) {
	println(v.Designation())
	println(v.Locomotion())
	println("Sentient?", v.Sentient())
	//v.Say() //error
}
