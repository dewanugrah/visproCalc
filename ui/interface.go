package ui

import (
	"visproCalc/calculator"
	"fmt"
)

func Start() {
	var num1, num2 float64 // Declare num1 and num2 outside the loop

	for {
		fmt.Println("==================")
		fmt.Println("--- Calculator ---")
		fmt.Println("==================")
		fmt.Println("1. Addition")
		fmt.Println("2. Subtraction")
		fmt.Println("3. Multiplication")
		fmt.Println("4. Division")
		fmt.Println("5. Power")
		fmt.Println("6. Exit")

		fmt.Print("Enter your choice (1-6): ")
		var choice int
		_, err := fmt.Scanln(&choice)

		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			// Clear the buffer
			fmt.Scanln()
			continue
		}

		if choice >= 1 && choice <= 5 {
			// Get input for num1
			fmt.Print("Enter the first number: ")
			_, err := fmt.Scanln(&num1)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number.")
				// Clear the buffer
				fmt.Scanln()
				continue
			}

			// Get input for num2
			fmt.Print("Enter the second number: ")
			_, err = fmt.Scanln(&num2)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number.")
				// Clear the buffer
				fmt.Scanln()
				continue
			}
		}

		switch choice {
		case 1:
			result := calculator.Add(num1, num2)
			fmt.Printf("Result: %f\n", result)
		case 2:
			result := calculator.Subtract(num1, num2)
			fmt.Printf("Result: %f\n", result)
		case 3:
			result := calculator.Multiply(num1, num2)
			fmt.Printf("Result: %f\n", result)
		case 4:
			result, err := calculator.Divide(num1, num2)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Result: %f\n", result)
			}
		case 5:
			result, _ := calculator.Power(num1, num2)
			fmt.Printf("Result: %f\n", result)
		case 6:
			fmt.Println("Exiting the calculator. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 6.")
		}
	}
}
