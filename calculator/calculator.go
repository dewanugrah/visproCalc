package calculator

import (
	"visproCalc/mathops"
)

func Add(a, b float64) float64 {
	return mathops.Addition(a, b)
}

func Subtract(a, b float64) float64 {
	return mathops.Subtraction(a, b)
}

func Multiply(a, b float64) float64 {
	return mathops.Multiplication(a, b)
}

func Divide(a, b float64) (float64, error) {
	return mathops.Division(a, b)
}

func Power(base, exponent float64) (float64, error) {
	return mathops.Power(base, exponent)
}
