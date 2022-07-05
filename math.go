package main

//Multiply returns the product of two integers
func Multiply(a, b int) int {
	return a * b
}
func Add(x, y int) int {

	return x + y
}
func Subtract(x, y int) int {
	return x - y
}

func Divide(x, y int) float64 {
	if y == 0 {
		return float64(0)
	}
	return float64(x / y)
}
