package main

import "fmt"

type Balance struct {
	Amount int
}

func (b Balance) Show() {
	fmt.Printf("現在の残高 = %d\n", b.Amount)
}

func (b *Balance) Deposit(value int) {
	b.Amount += value
}

func (b *Balance) Withdraw(value int) {
	b.Amount -= value
}

func main() {
	fmt.Println("== ポインタレシーバで値を更新する ==")

	balance := Balance{Amount: 1000}
	balance.Show()

	balance.Deposit(500)
	fmt.Println("500 入金後")
	balance.Show()

	balance.Withdraw(300)
	fmt.Println("300 出金後")
	balance.Show()

	fmt.Println()
	fmt.Println("== 値レシーバとの違い ==")
	fmt.Println("ポインタレシーバは元の構造体の値を更新できる")
}
