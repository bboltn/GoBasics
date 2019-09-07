package main

import "testing"

func Test_mostPowerful1(t *testing.T) {
	akira := &character{100}
	tetsuo := &character{90}
	result, err := mostPowerful(akira, tetsuo)
	if err != nil {
		t.Errorf("we should not get an error")
	}
	if result != akira {
		t.Errorf("Akira should be more powerful")
	}
}

func Test_mostPowerful2(t *testing.T) {
	tetsuo := &character{90}
	kei := &character{90}
	_, err := mostPowerful(kei, tetsuo)
	if err == nil {
		t.Errorf("we should get an error")
	}
}
