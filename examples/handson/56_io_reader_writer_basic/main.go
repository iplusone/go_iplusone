package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {
	// strings.NewReader → io.Reader として使う
	r := strings.NewReader("Hello, io.Reader!")

	// bytes.Buffer → io.Writer として使う
	var buf bytes.Buffer

	n, err := io.Copy(&buf, r)
	if err != nil {
		fmt.Println("エラー:", err)
		return
	}

	fmt.Printf("転送バイト数: %d\n", n)
	fmt.Printf("内容: %s\n", buf.String())

	// buf は io.Reader にもなる
	out, err := io.ReadAll(&buf)
	if err != nil {
		fmt.Println("エラー:", err)
		return
	}
	fmt.Printf("ReadAll で読み直し: %s\n", string(out))
}
