package main

import "fmt"

type Counter struct {
	Value int
}

func (c Counter) Show() {
	fmt.Printf("現在の値 = %d\n", c.Value)
}

func (c *Counter) Increment() {
	c.Value++
}

func main() {
	fmt.Println("== メソッドを呼び出す ==")

	counter := Counter{Value: 10}
	counter.Show()

	fmt.Println()
	fmt.Println("== ポインタレシーバのメソッド ==")

	counter.Increment()
	counter.Show()

	fmt.Println()
	fmt.Println("== 複数回呼び出す ==")

	counter.Increment()
	counter.Increment()
	counter.Show()
}
