package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func appendLine(path string, line string) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.WriteString(line + "\n"); err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println("== ファイルアクセスの基本 ==")

	tempDir, err := os.MkdirTemp("", "go-file-access-*")
	if err != nil {
		fmt.Println("temp dir error:", err)
		return
	}
	defer os.RemoveAll(tempDir)

	path := filepath.Join(tempDir, "data.txt")
	fmt.Printf("作業ファイル: %s\n", path)

	initialText := "Hello from Go\nThis file was created with os.WriteFile.\n"
	if err := os.WriteFile(path, []byte(initialText), 0o644); err != nil {
		fmt.Println("write error:", err)
		return
	}
	fmt.Println("1. os.WriteFile で初期内容を書き込みました")

	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("read error:", err)
		return
	}
	fmt.Println("2. os.ReadFile で読み込みました")
	fmt.Printf("   byte length = %d\n", len(data))
	fmt.Printf("   string(data) = %q\n", string(data))

	if err := appendLine(path, "*** start ***"); err != nil {
		fmt.Println("append error:", err)
		return
	}
	if err := appendLine(path, "追加メッセージ 1"); err != nil {
		fmt.Println("append error:", err)
		return
	}
	if err := appendLine(path, "追加メッセージ 2"); err != nil {
		fmt.Println("append error:", err)
		return
	}
	if err := appendLine(path, "*** end ***"); err != nil {
		fmt.Println("append error:", err)
		return
	}
	fmt.Println("3. os.OpenFile + WriteString で追記しました")

	finalData, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("read error:", err)
		return
	}

	fmt.Println()
	fmt.Println("== 最終内容 ==")
	fmt.Println(string(finalData))

	fmt.Println("メモ: ReadFile / WriteFile は簡単な入出力に向く")
	fmt.Println("メモ: 追記や細かい制御が必要な時は OpenFile を使う")
	fmt.Println("メモ: OpenFile で開いたファイルは defer で Close する")
}
