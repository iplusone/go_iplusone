package calc

import "testing"

func TestSum(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		want int
	}{
		{"空スライス", []int{}, 0},
		{"1要素", []int{5}, 5},
		{"複数要素", []int{1, 2, 3, 4, 5}, 15},
		{"負の値を含む", []int{-1, 2, -3}, -2},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Sum(tc.nums)
			if got != tc.want {
				t.Errorf("Sum(%v) = %d, want %d", tc.nums, got, tc.want)
			}
		})
	}
}

func TestFib(t *testing.T) {
	cases := []struct {
		n    int
		want int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{5, 5},
		{10, 55},
	}

	for _, tc := range cases {
		got := Fib(tc.n)
		if got != tc.want {
			t.Errorf("Fib(%d) = %d, want %d", tc.n, got, tc.want)
		}
		got2 := FibIter(tc.n)
		if got2 != tc.want {
			t.Errorf("FibIter(%d) = %d, want %d", tc.n, got2, tc.want)
		}
	}
}

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(20)
	}
}

func BenchmarkFibIter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibIter(20)
	}
}
