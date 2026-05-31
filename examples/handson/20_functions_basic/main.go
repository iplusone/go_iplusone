package main

import "fmt"

func main() {
	fmt.Println("== 関数の基本形 ==")

	greet()

	fmt.Println()
	fmt.Println("== 引数を受け取る関数 ==")

	message := sayHello("Go")
	fmt.Printf("sayHello(%q) = %q\n", "Go", message)

	fmt.Println()
	fmt.Println("== 複数の引数を受け取る関数 ==")

	total := add(10, 20)
	fmt.Printf("add(%d, %d) = %d\n", 10, 20, total)

	fmt.Println()
	fmt.Println("== シグネチャを見る ==")
	fmt.Println("greet のシグネチャ: func()")
	fmt.Println("sayHello のシグネチャ: func(string) string")
	fmt.Println("add のシグネチャ: func(int, int) int")
}

func greet() {
	fmt.Println("こんにちは、Go!")
}

func sayHello(name string) string {
	return "こんにちは、" + name + "!"
}

func add(a int, b int) int {
	return a + b
}
