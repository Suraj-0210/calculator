package calculator

var logMessage = "[LOG]"

// Version of calculator
var Version = "1.0"

func internalSum(number int) int {
	return number - 1
}

// Sum of two integers
func Sum(number1, number2 int) int {
	return number1 + number2
}
