package main

import "fmt"

/**
FUNCTIONS
**/
func main() {

	// super basic function
	helloGDG("HI GDG!")

	// function with parameters and return value
	result := isEqual(1, 2)
	println("isEqual: ", result)

	// named return type
	value := concat("Let's combine ", "words!")
	println("Concat: ", value)

	// multiple return types
	name, found := getUserName("1")
	println("Name: ", name)
	println("Found: ", found)

	// multiple return types
	name, found = getUserName("2")
	println("Name (empty!): ", name)
	println("Found (nope!): ", found)

	// using functions like variables with closures
	characters := []string{"Major", "Batou", "Togusa"}
	villians := []string{"Laughing Man", "Individual Eleven"}
	gitsCharacters := func(faction string) []string {
		if faction == "section 9" {
			return characters
		}
		return villians
	}
	fmt.Println(gitsCharacters("section 9"))
	takesAFunction(gitsCharacters)
}

func helloGDG(message string) {
	println(message)
}

func isEqual(a, b int) bool {
	return a == b
}

func concat(a, b string) (combinedValue string) {
	combinedValue = a + b
	return combinedValue
}

func getUserName(userID string) (string, bool) {
	if userID == "1" {
		return "Brian", true
	}

	return "", false
}

func takesAFunction(fn func(string) []string) {
	fmt.Println(fn("villians"))
}

/*
HI GDG!
isEqual:  false
Concat:  Let's combine words!
Name:  Brian
Found:  true
Name (nope!):
Found (nope!):  false
*/
