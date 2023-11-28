package mathops

import (
	"testing"
)

func TestAddition(t *testing.T) {
	result := Addition(3.0, 5.0)
	expected := 8.0
	if result != expected {
		t.Errorf("Expected %f, but got %f", expected, result)
	}
}

func TestSubtraction(t *testing.T) {
	result := Subtraction(10.0, 4.0)
	expected := 6.0
	if result != expected {
		t.Errorf("Expected %f, but got %f", expected, result)
	}
}

func TestMultiplication(t *testing.T) {
	result := Multiplication(2.0, 3.0)
	expected := 6.0
	if result != expected {
		t.Errorf("Expected %f, but got %f", expected, result)
	}
}

func TestDivision(t *testing.T) {
	result, err := Division(8.0, 4.0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := 2.0
	if result != expected {
		t.Errorf("Expected %f, but got %f", expected, result)
	}

	// Test division by zero
	_, err = Division(10.0, 0.0)
	if err == nil {
		t.Error("Expected error for division by zero, but got none")
	}
}

func TestPower(t *testing.T) {
	result, err := Power(2.0, 3.0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := 8.0
	if result != expected {
		t.Errorf("Expected %f, but got %f", expected, result)
	}

	// Test undefined result: zero to the power of a negative number
	_, err = Power(0.0, -2.0)
	if err == nil {
		t.Error("Expected error for zero to the power of a negative number, but got none")
	}

	// Test undefined result: negative base to a non-integer power
	_, err = Power(-2.0, 0.5)
	if err == nil {
		t.Error("Expected error for negative base to a non-integer power, but got none")
	}
}
