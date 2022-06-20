package main

import "testing"

func TestAddition(t *testing.T) {
	total := Addit(5, 5)
	if total != 10 {
		t.Errorf("")
	}
}
