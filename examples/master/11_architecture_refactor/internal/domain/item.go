package domain

import "errors"

var ErrNotFound = errors.New("item not found")

type Item struct {
	ID   int
	Name string
}

type ItemRepository interface {
	FindByID(id int) (*Item, error)
	Save(item *Item) error
	List() ([]*Item, error)
}
