package main

import "testing"

func TestAddition(t *testing.T) {
	total := Addition(5, 5)
	if total != 10 {
		t.Errorf("Addition was incorrect, got %d expected %d", total, 10)
	}

}
