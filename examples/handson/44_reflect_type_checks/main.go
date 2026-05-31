package main

import (
	"fmt"
	"reflect"
)

type General interface{}

type GData interface {
	Set(name string, value General) GData
	Print()
}

type NData struct {
	Name string
	Data int
}

type SData struct {
	Name string
	Data string
}

func (nd *NData) Set(name string, value General) GData {
	nd.Name = name
	if reflect.TypeOf(value).Kind() == reflect.Int {
		nd.Data = value.(int)
	}
	return nd
}

func (nd *NData) Print() {
	fmt.Printf("<<%s>> value: %d\n", nd.Name, nd.Data)
}

func (sd *SData) Set(name string, value General) GData {
	sd.Name = name
	if reflect.TypeOf(value).Kind() == reflect.String {
		sd.Data = value.(string)
	}
	return sd
}

func (sd *SData) Print() {
	fmt.Printf("* %s [%s] *\n", sd.Name, sd.Data)
}

func main() {
	fmt.Println("== reflect.TypeOf と Kind で型をチェックする ==")

	items := []GData{}
	items = append(items, new(NData).Set("Taro", 123))
	items = append(items, new(SData).Set("Jiro", "hello!"))
	items = append(items, new(NData).Set("Hanako", "98700"))
	items = append(items, new(SData).Set("Sachiko", []string{"happy?"}))

	for _, item := range items {
		item.Print()
	}

	fmt.Println()
	fmt.Println("型が合わない値は代入されず、ゼロ値のままになる")
}
