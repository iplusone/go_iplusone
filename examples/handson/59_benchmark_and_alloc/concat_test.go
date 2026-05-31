package benchmark

import (
	"strings"
	"testing"
)

// concatPlus は += で文字列を連結する（毎回アロケーションが発生する）。
func concatPlus(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += "a"
	}
	return s
}

// concatBuilder は strings.Builder で連結する（アロケーションを最小化）。
func concatBuilder(n int) string {
	var b strings.Builder
	for i := 0; i < n; i++ {
		b.WriteString("a")
	}
	return b.String()
}

func BenchmarkConcatPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatPlus(100)
	}
}

func BenchmarkConcatBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatBuilder(100)
	}
}
