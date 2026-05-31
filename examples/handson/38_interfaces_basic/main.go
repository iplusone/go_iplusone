package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

func (d Dog) Speak() string {
	return d.Name + " says: wan wan"
}

func (c Cat) Speak() string {
	return c.Name + " says: nyaa"
}

func introduce(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	fmt.Println("== インタフェースの定義 ==")
	fmt.Println("Speaker interface は Speak() string を持つ型を表す")

	fmt.Println()
	fmt.Println("== インタフェースを実装する ==")

	dog := Dog{Name: "Pochi"}
	cat := Cat{Name: "Tama"}

	fmt.Printf("dog = %+v (%T)\n", dog, dog)
	fmt.Printf("cat = %+v (%T)\n", cat, cat)

	fmt.Println()
	fmt.Println("== インタフェースとして利用する ==")

	introduce(dog)
	introduce(cat)

	fmt.Println()
	fmt.Println("== interface 型の変数に入れる ==")

	var speaker Speaker
	speaker = dog
	fmt.Printf("speaker = %T -> %q\n", speaker, speaker.Speak())

	speaker = cat
	fmt.Printf("speaker = %T -> %q\n", speaker, speaker.Speak())
}
