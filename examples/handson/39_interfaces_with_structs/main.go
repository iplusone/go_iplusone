package main

import "fmt"

type Data interface {
	Initial(name string, data []int)
	PrintData()
}

type MyData struct {
	Name string
	Data []int
}

type YourData struct {
	Name string
	Data []int
}

func (md *MyData) Initial(name string, data []int) {
	md.Name = name
	md.Data = data
}

func (md *MyData) PrintData() {
	fmt.Printf("MyData -> Name: %s, Data: %v\n", md.Name, md.Data)
}

func (md *MyData) Check() {
	fmt.Printf("Check! [%s]\n", md.Name)
}

func (yd *YourData) Initial(name string, data []int) {
	yd.Name = name
	yd.Data = data
}

func (yd *YourData) PrintData() {
	fmt.Printf("YourData -> Name: %s, Data: %v\n", yd.Name, yd.Data)
}

func showData(d Data, name string, values []int) {
	d.Initial(name, values)
	d.PrintData()
}

func main() {
	fmt.Println("== 構造体がインタフェースを実装する ==")

	var d Data = new(MyData)
	d.Initial("Sachiko", []int{55, 66, 77})
	d.PrintData()

	fmt.Println()
	fmt.Println("== interface 型では定義された機能だけ使える ==")

	my := new(MyData)
	my.Initial("Sachiko", []int{55, 66, 77})
	my.Check()

	fmt.Println("Data 型の変数からは Check は呼べない")

	fmt.Println()
	fmt.Println("== 異なる構造体を同じ interface で扱う ==")

	items := []Data{new(MyData), new(YourData)}
	names := []string{"Alpha", "Beta"}
	values := [][]int{{1, 2, 3}, {10, 20, 30}}

	for i, item := range items {
		showData(item, names[i], values[i])
	}
}
