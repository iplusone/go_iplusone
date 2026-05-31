package main

import (
	"errors"
	"log/slog"
	"os"

	"github.com/iplusone/go_iplusone/examples/master/11_architecture_refactor/internal/domain"
	"github.com/iplusone/go_iplusone/examples/master/11_architecture_refactor/internal/infra"
	"github.com/iplusone/go_iplusone/examples/master/11_architecture_refactor/internal/usecase"
)

func main() {
	// JSON 形式の構造化ログを標準出力へ
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	repo := infra.NewMemoryItemRepo()
	uc := usecase.NewItemUsecase(repo)

	// 作成
	item, err := uc.Create("Go言語入門")
	if err != nil {
		slog.Error("create failed", "error", err)
		return
	}
	slog.Info("item created", "id", item.ID, "name", item.Name)

	item2, err := uc.Create("Docker実践")
	if err != nil {
		slog.Error("create failed", "error", err)
		return
	}
	slog.Info("item created", "id", item2.ID, "name", item2.Name)

	// 取得
	found, err := uc.Get(1)
	if err != nil {
		slog.Error("get failed", "error", err)
		return
	}
	slog.Info("item found", "id", found.ID, "name", found.Name)

	// 存在しない ID
	_, err = uc.Get(99)
	if err != nil {
		slog.Warn("item not found", "id", 99, "error", err)
		// errors.Is でラップされたエラーを判定できる
		if errors.Is(err, domain.ErrNotFound) {
			slog.Info("confirmed: ErrNotFound via errors.Is")
		}
	}

	// 一覧
	items, err := uc.ListAll()
	if err != nil {
		slog.Error("list failed", "error", err)
		return
	}
	slog.Info("items listed", "count", len(items))
}
