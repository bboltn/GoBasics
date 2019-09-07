package main

func main() {
	// arrays length is part of its type
	// they cannot be resized
	var characters [4]string
	characters[0] = "Kaneda"
	characters[1] = "Tetsuo"
	characters[2] = "Kei"
	characters[3] = "Akira"
	println(characters[0])

	// array literals
	locations := [2]string{"neo-tokyo", "Harukiya"}
	println(locations[0])

	// slices function like arrays but can be resized
	// they are backed by an underlying array
	moreCharacters := []string{
		"Kaneda", "Tetsuo", "Kei", "Akira",
	}
	println(moreCharacters[0])
	moreCharacters = append(moreCharacters, "Yamagata")

	// they are called slices, because you can create them from a "slice"
	// of an array like so...
	mainCharacters := characters[0:2]

	// [inclusive:exclusive]
	println(mainCharacters[0])
	println(mainCharacters[1])

	//extend length
	println("Length of mainCharacters: ", len(mainCharacters))
	mainCharacters = mainCharacters[:4]
	println("Length of mainCharacters after extending: ", len(mainCharacters))

	// use make to build slices of a particular length and capacity
	// length 2, capacity 4 - capacity is optional
	otherLocations := make([]string, 2, 4)
	otherLocations[0] = "toyko"
	otherLocations[1] = "eighth district"

	// specifying the capacity allocates the memory,
	// but you won't be able to address it until you extend it.
	//otherLocations[2] = "temple"//out of range

	// instead use "append"
	otherLocations = append(otherLocations, "temple")
	println("otherLocations length: ", len(otherLocations))

	// If I don't know how large I want the list to be,
	// I declare slices like so
	newSlice := []string{}

	// and will then use append to add items (typically in a loop)
	// i is the index and v is the value
	for i, v := range otherLocations {
		newSlice = append(newSlice, v)
		println(newSlice[i])
	}

	// I prefer slices over arrays because they give me a
	// dynamic "list" type I'm used to from languages
	// like python

	// nil slice will have length 0
	println(len([]string{}))

	// slice of slices
	factions := make([][]string, 2)

	factions[0] = []string{"kaneda", "yamagata"}
	factions[1] = []string{"akira", "tetsuo"}

	println(factions[0][0])
	println(factions[0][1])
	println(factions[1][0])
	println(factions[1][1])
}
