package main

import (
	"errors"
	"fmt"

	"github.com/iplusone/go_iplusone/examples/master/04_abstraction_and_di/service"
	"github.com/iplusone/go_iplusone/examples/master/04_abstraction_and_di/store"
)

func main() {
	// MemoryStore（Fake）を注入する。本番では DB 実装を渡せばよい。
	s := store.NewMemoryStore()
	svc := service.NewUserService(s)

	if err := svc.Register(1, "Alice"); err != nil {
		fmt.Println("登録エラー:", err)
		return
	}
	fmt.Println("Alice を登録しました")

	name, err := svc.GetName(1)
	if err != nil {
		fmt.Println("取得エラー:", err)
		return
	}
	fmt.Println("取得した名前:", name)

	// 重複登録
	if err := svc.Register(1, "Alice2"); err != nil {
		fmt.Println("重複登録エラー（期待通り）:", err)
	}

	// 存在しない ID
	_, err = svc.GetName(99)
	if err != nil {
		fmt.Println("存在しないIDのエラー（期待通り）:", err)
		fmt.Println("ErrNotFound か?", errors.Is(err, store.ErrNotFound))
	}
}
