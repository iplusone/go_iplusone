package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Data interface {
	SetValue(vals map[string]string)
	PrintData()
}

type MyData struct {
	Name string
	Data []int
}

type YourData struct {
	Name string
	Mail string
	Age  int
}

func (md *MyData) SetValue(vals map[string]string) {
	md.Name = vals["name"]

	parts := strings.Split(vals["data"], ",")
	items := make([]int, 0, len(parts))
	for _, part := range parts {
		n, err := strconv.Atoi(strings.TrimSpace(part))
		if err == nil {
			items = append(items, n)
		}
	}

	md.Data = items
}

func (md *MyData) PrintData() {
	fmt.Println("MyData")
	fmt.Println("  Name:", md.Name)
	fmt.Println("  Data:", md.Data)
}

func (yd *YourData) SetValue(vals map[string]string) {
	yd.Name = vals["name"]
	yd.Mail = vals["mail"]

	age, err := strconv.Atoi(vals["age"])
	if err == nil {
		yd.Age = age
	}
}

func (yd *YourData) PrintData() {
	fmt.Printf("YourData\n  Name: %s\n  Mail: %s\n  Age: %d\n", yd.Name, yd.Mail, yd.Age)
}

func main() {
	fmt.Println("== 異なる構造体を interface の配列でまとめる ==")

	items := [2]Data{}

	items[0] = new(MyData)
	items[0].SetValue(map[string]string{
		"name": "Sachiko",
		"data": "55, 66, 77",
	})

	items[1] = new(YourData)
	items[1].SetValue(map[string]string{
		"name": "Mami",
		"mail": "mami@mume.mo",
		"age":  "34",
	})

	for _, item := range items {
		item.PrintData()
		fmt.Println()
	}
}
