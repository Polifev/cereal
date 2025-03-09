package math

import "testing"

func TestVector_SquareLength(t *testing.T) {
	vectors := []Vector{
		NewVector(0, 0),
		NewVector(1, 1),
		NewVector(2, 2),
		NewVector(-2, 2),
		NewVector(-1, -2),
	}

	expectedSqrLength := []float64{
		0,
		2,
		8,
		8,
		5,
	}

	for i := 0; i < len(vectors); i++ {
		sqrl := vectors[i].SquareLength()
		if sqrl != expectedSqrLength[i] {
			t.Errorf("SquareLength() of %v returned %f, expected %f", vectors[i], sqrl, expectedSqrLength[i])
		}
	}
}

func TestVector_Length(t *testing.T) {
	vectors := []Vector{
		NewVector(3, 4),
	}

	expectedLength := []float64{
		5,
	}

	for i := 0; i < len(vectors); i++ {
		l := vectors[i].Length()
		if l != expectedLength[i] {
			t.Errorf("Length() of %v returned %f, expected %f", vectors[i], l, expectedLength[i])
		}
	}
}

func TestVector_Add(t *testing.T) {
	a := NewVector(1, 2)
	b := NewVector(3, -4)
	result := a.Add(b)
	if result.X() != 4 || result.Y() != -2 {
		t.Errorf("%v + %v returned %v, expected %v", a, b, result, Vector{4, -2})
	}
}

func TestVector_Sub(t *testing.T) {
	a := NewVector(1, 2)
	b := NewVector(3, -4)
	result := a.Sub(b)
	if result.X() != -2 || result.Y() != 6 {
		t.Errorf("%v + %v returned %v, expected %v", a, b, result, Vector{-2, 6})
	}
}

func TestVector_Dot(t *testing.T) {
	a := NewVector(1, 2)

	b := NewVector(1, 2)
	c := NewVector(0.999999, 2.00000001)
	d := NewVector(0, 4)

	if !a.EqRough(b) {
		t.Errorf("%v == %v returned false, expected true", a, b)
	}

	if !a.EqRough(c) {
		t.Errorf("%v == %v returned false, expected true", a, c)
	}

	if a.EqRough(d) {
		t.Errorf("%v == %v returned true, expected false", a, d)
	}
}
