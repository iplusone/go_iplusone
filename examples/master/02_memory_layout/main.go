package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("=== 値渡しと参照渡し ===")
	demoValueVsPointer()

	fmt.Println("\n=== スライスの内部構造と append ===")
	demoSliceInternals()

	fmt.Println("\n=== 文字列とバイト表現 ===")
	demoStringBytes()

	fmt.Println("\n=== 構造体のサイズとアライメント ===")
	demoStructSize()
}

func demoValueVsPointer() {
	x := 42
	double(x)
	fmt.Printf("値渡し後: x = %d (変わらない)\n", x)

	doublePtr(&x)
	fmt.Printf("ポインタ渡し後: x = %d (変わる)\n", x)
}

func double(n int) {
	n *= 2
}

func doublePtr(n *int) {
	*n *= 2
}

func demoSliceInternals() {
	s := make([]int, 0, 4)
	fmt.Printf("初期状態: len=%d cap=%d\n", len(s), cap(s))

	for i := 0; i < 6; i++ {
		s = append(s, i)
		fmt.Printf("append(%d): len=%d cap=%d\n", i, len(s), cap(s))
	}

	// スライスを別変数に代入しても実体は共有される（容量内の場合）
	a := make([]int, 3)
	a[0], a[1], a[2] = 1, 2, 3
	b := a
	b[0] = 99
	fmt.Printf("a[0]=%d b[0]=%d (同じ実体を共有)\n", a[0], b[0])
}

func demoStringBytes() {
	s := "Hello"
	fmt.Printf("\"Hello\": len=%d (バイト数)\n", len(s))

	j := "日本語"
	fmt.Printf("\"日本語\": len=%d (バイト数、文字数ではない)\n", len(j))
	fmt.Printf("\"日本語\": rune数=%d\n", len([]rune(j)))

	for i, r := range j {
		fmt.Printf("  byte[%d] = %q\n", i, r)
	}

	// string から []byte への変換
	b := []byte(j)
	fmt.Printf("[]byte 長さ: %d\n", len(b))
}

type Inefficient struct {
	A bool    // 1バイト
	B int64   // 8バイト
	C bool    // 1バイト
	D int64   // 8バイト
}

type Efficient struct {
	B int64   // 8バイト
	D int64   // 8バイト
	A bool    // 1バイト
	C bool    // 1バイト
}

func demoStructSize() {
	fmt.Printf("非効率な順序: %d バイト\n", unsafe.Sizeof(Inefficient{}))
	fmt.Printf("効率的な順序: %d バイト\n", unsafe.Sizeof(Efficient{}))
}
