package calculator

func Plus(x, y int) int {
	return x + y
}

func Minus(x, y int) int {
	return x - y
}

func Devide(x, y int) int {
	return x / y
}

func Multiply(x, y int) int {
	return x * y
}

func Pow(x, y int) int {

	ret := 1
	for loop := 0; loop < y; loop++ {
		ret = ret * x
	}
	return ret
}
