package calc

// Sum は整数スライスの合計を返す。
func Sum(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// Fib はフィボナッチ数列の n 番目を返す（再帰版、遅い）。
func Fib(n int) int {
	if n <= 1 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

// FibIter はフィボナッチ数列の n 番目を返す（反復版、速い）。
func FibIter(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}
