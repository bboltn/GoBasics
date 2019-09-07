package main

import "fmt"

func main() {

	// creating a map
	characterLocations := map[string][]string{}

	// and adding values to specific keys
	characterLocations["temple"] = []string{
		"Lady Miyako",
	}
	characterLocations["eighth district"] = []string{
		"kaneda", "yamagata", "tetsuo",
	}

	// map literal
	characterLocations = map[string][]string{
		"temple": {
			"Lady Miyako",
		},
		"eighth district": {
			"kaneda", "yamagata", "tetsuo",
		},
	}

	fmt.Println(characterLocations)

	// deleting keys in a map
	delete(characterLocations, "temple")
	fmt.Println(characterLocations)

	// does a key exist?
	_, ok := characterLocations["temple"]
	if !ok {
		println("Does not exist!")
	}

	value, ok := characterLocations["eighth district"]
	if ok {
		fmt.Println("Value found: ", value)
	}

	// you can also combine this with an if/else
	if value, ok := characterLocations["eighth district"]; ok {
		fmt.Println("Value found: ", value)
	}
}
