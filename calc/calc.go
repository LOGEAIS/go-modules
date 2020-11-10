package calc

func Multiplication(x, y int) int {
	return x * y
}

func Soustraction(x, y int) int {
	return x - y
}

func Addition(x, y int) int {
	return x + y
}

func AllOperation(x, y int) int {
	return Multiplication(x, y) + Soustraction(x, y) + Addition(x, y)
}
