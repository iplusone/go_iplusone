package service

import (
	"errors"
	"fmt"

	"github.com/iplusone/go_iplusone/examples/master/04_abstraction_and_di/store"
)

// UserService は UserStore に依存するが、具体的な実装は知らない。
type UserService struct {
	store store.UserStore
}

// NewUserService はコンストラクタ注入の形。外から store を渡す。
func NewUserService(s store.UserStore) *UserService {
	return &UserService{store: s}
}

func (s *UserService) Register(id int, name string) error {
	_, err := s.store.FindByID(id)
	if err == nil {
		return fmt.Errorf("user %d already exists", id)
	}
	if !errors.Is(err, store.ErrNotFound) {
		return err
	}
	return s.store.Save(&store.User{ID: id, Name: name})
}

func (s *UserService) GetName(id int) (string, error) {
	u, err := s.store.FindByID(id)
	if err != nil {
		return "", fmt.Errorf("GetName: %w", err)
	}
	return u.Name, nil
}
