package main

import (
	"fmt"
	"reflect"
	"strings"
)

type General interface{}

type GData interface {
	Set(name string, value General) GData
	Print()
}

type NData struct {
	Name string
	Data []int
}

type SData struct {
	Name string
	Data []string
}

func (nd *NData) Set(name string, value General) GData {
	nd.Name = name
	if reflect.TypeOf(value) == reflect.SliceOf(reflect.TypeOf(0)) {
		nd.Data = value.([]int)
	}
	return nd
}

func (nd *NData) Print() {
	fmt.Printf("<<%s>> value: %v\n", nd.Name, nd.Data)
}

func (sd *SData) Set(name string, value General) GData {
	sd.Name = name
	if reflect.TypeOf(value) == reflect.SliceOf(reflect.TypeOf("")) {
		sd.Data = value.([]string)
	}
	return sd
}

func (sd *SData) Print() {
	fmt.Printf("* %s [%s] *\n", sd.Name, strings.Join(sd.Data, " "))
}

func main() {
	fmt.Println("== スライスの型を reflect でチェックする ==")

	items := []GData{}
	items = append(items, new(NData).Set("Taro", []int{1, 2, 3}))
	items = append(items, new(SData).Set("Jiro", []string{"hello", "bye"}))
	items = append(items, new(NData).Set("Hanako", 98700))
	items = append(items, new(SData).Set("Sachiko", "happy?"))

	for _, item := range items {
		item.Print()
	}

	fmt.Println()
	fmt.Println("int スライスと string スライスだけを受け取るようにしている")
}
