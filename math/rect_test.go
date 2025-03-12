package math

import "testing"

func TestRect_Contains(t *testing.T) {
	rect := NewRectFromFloat(-5.0, 4.0, 10.0, 6.0)

	vectors := []Vector{
		NewVector(1.0, 1.0),
		NewVector(4.0, 7.0),
	}

	expectedResults := []bool{
		false,
		true,
	}

	for i := 0; i < len(vectors); i++ {
		if rect.Contains(vectors[i]) != expectedResults[i] {
			t.Errorf(
				"Rect %v Contains(%v) returned %v, expected %v",
				vectors[i],
				vectors[i],
				!expectedResults[i],
				expectedResults[i])
		}
	}
}
