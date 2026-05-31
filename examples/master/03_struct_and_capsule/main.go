package main

import (
	"fmt"
	"unsafe"

	"github.com/iplusone/go_iplusone/examples/master/03_struct_and_capsule/counter"
)

func main() {
	fmt.Println("=== 構造体のサイズとパディング ===")
	demoStructPadding()

	fmt.Println("\n=== 値レシーバとポインタレシーバ ===")
	demoReceivers()

	fmt.Println("\n=== カプセル化されたカウンタ ===")
	demoEncapsulation()
}

type BadLayout struct {
	A bool   // 1バイト + 7バイトのパディング
	B int64  // 8バイト
	C bool   // 1バイト + 7バイトのパディング
	D int64  // 8バイト
}

type GoodLayout struct {
	B int64  // 8バイト
	D int64  // 8バイト
	A bool   // 1バイト
	C bool   // 1バイト + 6バイトのパディング
}

func demoStructPadding() {
	fmt.Printf("BadLayout:  %d バイト\n", unsafe.Sizeof(BadLayout{}))
	fmt.Printf("GoodLayout: %d バイト\n", unsafe.Sizeof(GoodLayout{}))
}

type Point struct {
	X, Y int
}

// 値レシーバ: コピーに対して操作するので元は変わらない
func (p Point) ScaleValue(factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}

// ポインタレシーバ: 実体を操作する
func (p *Point) ScalePtr(factor int) {
	p.X *= factor
	p.Y *= factor
}

func demoReceivers() {
	p := Point{X: 3, Y: 4}

	result := p.ScaleValue(2)
	fmt.Printf("値レシーバ後: p=%v (元は不変), result=%v\n", p, result)

	p.ScalePtr(2)
	fmt.Printf("ポインタレシーバ後: p=%v (元が変わる)\n", p)
}

func demoEncapsulation() {
	c := &counter.Counter{}
	c.Increment()
	c.Increment()
	c.Increment()
	fmt.Printf("カウンタ値: %d\n", c.Value())
	c.Reset()
	fmt.Printf("リセット後: %d\n", c.Value())
	// c.n = 99 // コンパイルエラー: フィールド n はパッケージ外から見えない
}
