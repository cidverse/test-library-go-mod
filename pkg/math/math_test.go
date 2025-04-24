package math

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{1, 2, 3},
		{0, 0, 0},
		{-1, -1, -2},
	}

	for _, test := range tests {
		result := Add(test.a, test.b)
		if result != test.expected {
			t.Errorf("Add(%d, %d) = %d; expected %d", test.a, test.b, result, test.expected)
		}
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{3, 2, 1},
		{0, 0, 0},
		{-1, -1, 0},
	}

	for _, test := range tests {
		result := Subtract(test.a, test.b)
		if result != test.expected {
			t.Errorf("Subtract(%d, %d) = %d; expected %d", test.a, test.b, result, test.expected)
		}
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{2, 3, 6},
		{0, 5, 0},
		{-1, -1, 1},
	}

	for _, test := range tests {
		result := Multiply(test.a, test.b)
		if result != test.expected {
			t.Errorf("Multiply(%d, %d) = %d; expected %d", test.a, test.b, result, test.expected)
		}
	}
}
