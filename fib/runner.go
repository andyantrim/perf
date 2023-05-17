package fib

func FibRecursive(n int) int {
        if n < 2 {
                return n
        }
        return FibRecursive(n-1) + FibRecursive(n-2)
}

func FibIterative(n int) int {
	if n < 2 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}
