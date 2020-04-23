package text

//递归函数
func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

//斐波那契数列
func Fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return Fibonacci(n-2) + Fibonacci(n-1)
}

//第二种
func fibonacci2(n int) (int, int) {
	if n < 2 {
		return 0, n
	}
	a, b := fibonacci2(n - 1)
	return b, a + b
}

func Fibonacci1(n int) int {
	var b, _ int
	_, b = fibonacci2(n)
	return b
}
