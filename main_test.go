package main

import "testing"

func TestAddition(t *testing.T) {
	tables := []struct {
		a int
		b int
		n int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 26, 51},
	}
	for _, item := range tables {
		total := Addition(item.a, item.b)
		if total != item.n {
			t.Errorf("Addition was incorrect, got %d expected %d", total, 10)
		}
	}
}

func TestMax(t *testing.T) {
	tables := []struct {
		a int
		b int
		n int
	}{
		{4, 2, 4},
	}
}
