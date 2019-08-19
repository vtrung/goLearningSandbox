package main

import "testing"


func TestAdder(t *testing.T) {
	sum := addingfloat(2.0, 2.0)
	expected := 4.0

	if sum != expected {
		t.Errorf("expected '%f' but got '%f'", expected, sum)
	}

	sum2 := addingfloat(189.4, 21.6)
	expected2 := 211.0

	if sum2 != expected2 {
		t.Errorf("expected '%f' but got '%f'", expected2, sum2)
	}
}