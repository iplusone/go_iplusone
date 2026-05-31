package main

import "fmt"

func main() {
	fmt.Println("== switch と default ==")

	level := "gold"
	switch level {
	case "bronze":
		fmt.Println("ブロンズ会員です")
	case "silver":
		fmt.Println("シルバー会員です")
	default:
		fmt.Println("default: どの case にも一致しませんでした")
	}

	fmt.Println()
	fmt.Println("== fallthrough なし ==")

	day := "sat"
	switch day {
	case "sat":
		fmt.Println("土曜日です")
	case "sun":
		fmt.Println("日曜日です")
	default:
		fmt.Println("平日です")
	}

	fmt.Println()
	fmt.Println("== fallthrough あり ==")

	score := 85
	switch {
	case score >= 80:
		fmt.Println("80点以上です")
		fallthrough
	case score >= 60:
		fmt.Println("60点以上です")
	default:
		fmt.Println("判定を終了します")
	}

	fmt.Println()
	fmt.Println("メモ: fallthrough は次の case 条件を再評価せず、そのまま次へ進む")
}
