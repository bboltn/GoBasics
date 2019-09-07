package main

// FLOW CONTROL
func main() {
	// useful for logging or other cleanup
	defer println("DO THIS LAST!")

	// basic for
	for i := 0; i < 5; i++ {
		println("For: ", i)
	}

	// chop up the for loop all kinds of ways

	// drop declararion
	j := 0
	for ; j < 5; j++ {
		println("Another for ", j)
	}

	// drop declaration and comparison
	k := 0
	for ; ; k++ {
		if k == 5 {
			break
		}
		println("Another nother for ", k)
	}

	// drop all of them to make a "while" loop
	i := 0
	for {
		if i == 5 {
			break
		}
		println("For/while ", i)
		i++
	}

	// ranging over a slice (more on slices later)
	for i, v := range []string{"Major Motoko Kusanagi", "Batou"} {
		println("Character: index: ", i, "value: ", v)
	}

	// if / else if / else
	if i > 5 {
		println("Greater than 5!")
	} else if i < 5 {
		println("Less than 5!")
	} else {
		println("Equal to 5!")
	}

	// switch
	switch i {
	case 1:
		println("Value is 1")
	case 2:
		println("Value is 2")
	case 5:
		println("Value is 5")
	default:
		println("Value is ", i)
	}
}
