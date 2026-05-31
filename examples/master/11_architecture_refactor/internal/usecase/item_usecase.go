package usecase

import (
	"fmt"

	"github.com/iplusone/go_iplusone/examples/master/11_architecture_refactor/internal/domain"
)

type ItemUsecase struct {
	repo   domain.ItemRepository
	nextID int
}

func NewItemUsecase(repo domain.ItemRepository) *ItemUsecase {
	return &ItemUsecase{repo: repo, nextID: 1}
}

func (u *ItemUsecase) Create(name string) (*domain.Item, error) {
	if name == "" {
		return nil, fmt.Errorf("Create: name is required")
	}
	item := &domain.Item{ID: u.nextID, Name: name}
	u.nextID++
	if err := u.repo.Save(item); err != nil {
		return nil, fmt.Errorf("Create: %w", err)
	}
	return item, nil
}

func (u *ItemUsecase) Get(id int) (*domain.Item, error) {
	item, err := u.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("Get: %w", err)
	}
	return item, nil
}

func (u *ItemUsecase) ListAll() ([]*domain.Item, error) {
	items, err := u.repo.List()
	if err != nil {
		return nil, fmt.Errorf("ListAll: %w", err)
	}
	return items, nil
}
