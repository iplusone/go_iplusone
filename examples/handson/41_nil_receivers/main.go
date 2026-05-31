package main

import (
	"fmt"
	"strconv"
	"strings"
)

type MyData struct {
	Name string
	Data []int
}

func (md *MyData) SetValue(vals map[string]string) {
	if md == nil {
		return
	}

	md.Name = vals["name"]

	parts := strings.Fields(vals["data"])
	items := make([]int, 0, len(parts))
	for _, part := range parts {
		n, err := strconv.Atoi(part)
		if err == nil {
			items = append(items, n)
		}
	}

	md.Data = items
}

func (md *MyData) PrintData() {
	if md == nil {
		fmt.Println("**This is nil value.**")
		return
	}

	fmt.Println("Name:", md.Name)
	fmt.Println("Data:", md.Data)
}

func main() {
	fmt.Println("== nil レシーバを確認する ==")

	var item *MyData
	item.PrintData()

	fmt.Println()
	fmt.Println("== 値を入れたあとに同じメソッドを使う ==")

	item = &MyData{}
	item.SetValue(map[string]string{
		"name": "Jiro",
		"data": "123 456 789",
	})
	item.PrintData()

	fmt.Println()
	fmt.Println("nil でもメソッド自体は呼べるが、nil チェックが必要")
}
