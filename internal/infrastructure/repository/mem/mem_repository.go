package mem

import (
	"context"
	"sync"

	"exchange/internal/domain/user"
)

type Repository struct {
	db map[string]struct{}
	mu sync.RWMutex
}

func NewMemoryRepository() *Repository {
	return &Repository{
		db: make(map[string]struct{}),
		mu: sync.RWMutex{},
	}
}

func (m *Repository) SaveUser(_ context.Context, eu *user.User) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.db[eu.Email] = struct{}{}

	return nil
}

func (m *Repository) GetByEmail(
	_ context.Context,
	email string,
) (*user.User, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	_, ok := m.db[email]
	if !ok {
		return nil, user.ErrNotFound
	}

	return user.NewUser(email), nil
}

func (m *Repository) GetAllEmails(
	_ context.Context,
) ([]string, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	emails := make([]string, 0, len(m.db))

	for key := range m.db {
		emails = append(emails, key)
	}

	return emails, nil
}

func (m *Repository) EmailExist(_ context.Context, email string) (bool, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	_, ok := m.db[email]

	return ok, nil
}
