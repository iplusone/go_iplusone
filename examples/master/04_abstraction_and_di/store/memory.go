package store

// MemoryStore は UserStore の Fake 実装。テストや手元確認に使う。
type MemoryStore struct {
	data map[int]*User
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{data: make(map[int]*User)}
}

func (m *MemoryStore) FindByID(id int) (*User, error) {
	u, ok := m.data[id]
	if !ok {
		return nil, ErrNotFound
	}
	return u, nil
}

func (m *MemoryStore) Save(u *User) error {
	m.data[u.ID] = u
	return nil
}
