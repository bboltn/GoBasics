package main

import (
	"errors"
	"fmt"
)

// Go doesn't have try/catch or throw errors
// Go handles errors by returning an object that implements the error interface

/*
type error interface {
    Error() string
}
*/
func main() {
	// common pattern in go
	// call the method, then check for err != nil, and then handle it
	err := AkiraLosesControl()
	if err != nil {
		println(err.Error())
	}

	// because of multiple return types, you will
	// often have the error as the last type in the return type list
	// you will first check the error and if its not nil, you can then
	// use the values

	chars, err := getFaction("akira")
	if err != nil {
		println(err.Error())
	}
	fmt.Println(chars)

	chars, err = getFaction("other")
	if err != nil {
		println("error:", err.Error())
	}
}

// AkiraLosesControl returns an error
func AkiraLosesControl() error {
	return errors.New("akira causes a massive explosion")
}

func getFaction(faction string) ([]string, error) {
	if faction == "akira" {
		return []string{"tetsuo", "akira"}, nil
	}

	if faction == "rebels" {
		return []string{"kei"}, nil
	}

	return nil, errors.New("not a faction")

}
