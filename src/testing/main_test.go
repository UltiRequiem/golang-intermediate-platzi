package main

import "testing"

/*
func TestSum(t *testing.T) {
	total := Sum(1, 3)

	if total != 4 {
		t.Errorf("Sum was incorrect, got %d expected %d", total, 10)
	}
} */

func TestSum(t *testing.T) {
	tables := []struct {
		a int
		b int
		c int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 26, 51},
	}

	for _, item := range tables {
		total := Sum(item.a, item.b)
		if total != item.c {
			t.Errorf("Sum was incorrect, got %d expected %d", total, item.c)
		}
	}

}
