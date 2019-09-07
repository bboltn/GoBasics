package main

import "errors"

type character struct {
	power int
}

func mostPowerful(c1, c2 *character) (*character, error) {
	if c1.power > c2.power {
		return c1, nil
	} else if c1.power < c2.power {
		return c2, nil
	}
	return nil, errors.New("equal in power")
}

func main() {}
