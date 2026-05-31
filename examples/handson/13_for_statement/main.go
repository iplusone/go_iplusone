package main

import "fmt"

func main() {
	fmt.Println("== 基本の for ==")

	for i := 0; i < 5; i++ {
		fmt.Printf("i = %d\n", i)
	}

	fmt.Println()
	fmt.Println("== 条件だけの for ==")

	count := 0
	for count < 3 {
		fmt.Printf("count = %d\n", count)
		count++
	}

	fmt.Println()
	fmt.Println("== range を使う for ==")

	names := []string{"Go", "Java", "Python"}
	for index, name := range names {
		fmt.Printf("index = %d, name = %q\n", index, name)
	}

	fmt.Println()
	fmt.Println("== index を使わない range ==")

	for _, name := range names {
		fmt.Printf("name = %q\n", name)
	}
}
