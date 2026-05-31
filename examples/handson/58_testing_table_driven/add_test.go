package tabledriven

import "testing"

func TestAdd(t *testing.T) {
	cases := []struct {
		name string
		a, b int
		want int
	}{
		{"正の数", 1, 2, 3},
		{"負の数", -1, -2, -3},
		{"ゼロ", 0, 0, 0},
		{"正と負", 5, -3, 2},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Add(tc.a, tc.b)
			if got != tc.want {
				t.Errorf("Add(%d, %d) = %d, want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}

func TestAbs(t *testing.T) {
	cases := []struct {
		name string
		n    int
		want int
	}{
		{"正の数", 5, 5},
		{"負の数", -3, 3},
		{"ゼロ", 0, 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Abs(tc.n)
			if got != tc.want {
				t.Errorf("Abs(%d) = %d, want %d", tc.n, got, tc.want)
			}
		})
	}
}
