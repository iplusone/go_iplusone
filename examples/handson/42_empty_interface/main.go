package main

import "fmt"

type General interface{}

type GData interface {
	Set(name string, value General)
	Print()
}

type GDataImpl struct {
	Name string
	Data General
}

func (gd *GDataImpl) Set(name string, value General) {
	gd.Name = name
	gd.Data = value
}

func (gd *GDataImpl) Print() {
	fmt.Printf("<<%s>> %v (%T)\n", gd.Name, gd.Data, gd.Data)
}

func main() {
	fmt.Println("== 空の interface はさまざまな値を保持できる ==")

	items := []GDataImpl{
		{Name: "Taro", Data: 123},
		{Name: "Hanako", Data: "hello!"},
		{Name: "Sachiko", Data: []int{123, 456, 789}},
		{Name: "Active", Data: true},
	}

	for _, item := range items {
		item.Print()
	}

	fmt.Println()
	fmt.Println("== Set でもいろいろな値を入れられる ==")

	var box GData = new(GDataImpl)
	box.Set("Rate", 0.01)
	box.Print()
}
