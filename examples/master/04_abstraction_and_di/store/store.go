package store

import "errors"

var ErrNotFound = errors.New("not found")

type User struct {
	ID   int
	Name string
}

// UserStore は上位ロジックが依存する interface。実装を知らなくてよい。
type UserStore interface {
	FindByID(id int) (*User, error)
	Save(u *User) error
}
