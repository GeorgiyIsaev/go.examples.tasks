package s01_math

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Add(2, 3) = %d; expected %d", result, expected)
	}
}

func TestDivision(t *testing.T) {
	if result, err := Divide(10, 0); err == nil {
		t.Errorf("Expected error for division by zero, got %d", result)
	}
}
