package main

import "fmt"

// Character akira manga
type Character struct {
	Name    string
	Faction string
}

func main() {
	akira := &Character{"Akira", "akira"}
	tetsuo := &Character{"Tetsuo", "akira"}

	akira.Flight()
	tetsuo.Explode()

	yamagata := Character{"Yamagata", "kaneda"}
	yamagata.Explode()

	// address of yamagata
	fmt.Printf("yamagata's address in main: %p\n", &yamagata)

	// values receivers cannot change their underlying value
	yamagata.ChangeFaction("NEW FACTION")
	println(yamagata.Faction) // still kaneda

	// pointer receivers can change their underlying value
	yamagata.ChangeFactionPointerReceiver("NEW FACTION")
	println(yamagata.Faction)

	// You can call pointer or value receiver methods on either type
	// reason's to use pointer receivers over value receivers
	// - they don't copy the struct (use less memory)
	// - they can mutate the underlying fields in the struct
	// you can see that it's copying the value by looking at its memory address
}

// Flight power
func (c Character) Flight() {
	println(c.Name, "is flying")
}

// Crush power
func (c Character) Crush() {
	println(c.Name, "is crushing")
}

// Explode power
func (c Character) Explode() {
	println(c.Name, "is causing an explosion")
}

// ChangeFaction method
func (c Character) ChangeFaction(newFaction string) {
	fmt.Printf("yamagata's address in change faction: %p\n", &c)
	c.Faction = newFaction
}

// ChangeFactionPointerReceiver method
func (c *Character) ChangeFactionPointerReceiver(newFaction string) {
	fmt.Printf("yamagata's address in pointer version: %p\n", c)
	c.Faction = newFaction
}
