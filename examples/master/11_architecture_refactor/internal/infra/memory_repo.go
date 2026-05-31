package infra

import (
	"github.com/iplusone/go_iplusone/examples/master/11_architecture_refactor/internal/domain"
)

type MemoryItemRepo struct {
	data map[int]*domain.Item
}

func NewMemoryItemRepo() *MemoryItemRepo {
	return &MemoryItemRepo{data: make(map[int]*domain.Item)}
}

func (r *MemoryItemRepo) FindByID(id int) (*domain.Item, error) {
	item, ok := r.data[id]
	if !ok {
		return nil, domain.ErrNotFound
	}
	return item, nil
}

func (r *MemoryItemRepo) Save(item *domain.Item) error {
	r.data[item.ID] = item
	return nil
}

func (r *MemoryItemRepo) List() ([]*domain.Item, error) {
	result := make([]*domain.Item, 0, len(r.data))
	for _, item := range r.data {
		result = append(result, item)
	}
	return result, nil
}
