package filesystem

import (
	"context"
	"fmt"
	"os"
	"strings"

	"exchange/internal/domain/user_domain"
)

func (f *Repository) SaveUser(_ context.Context, eu *user_domain.User) error {
	f.fm.Lock()
	defer f.fm.Unlock()

	file, err := os.OpenFile(f.filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return fmt.Errorf("failed to open data file: %w", err)
	}
	defer file.Close()

	if _, err = file.WriteString(addEndOfTheLine(eu.Email)); err != nil {
		return fmt.Errorf("failed write to file")
	}

	f.im.Lock()
	defer f.im.Unlock()

	f.index[eu.Email] = struct{}{}

	return nil
}

func (f *Repository) GetByEmail(
	_ context.Context,
	email string,
) (*user_domain.User, error) {
	f.im.RLock()
	defer f.im.RUnlock()

	_, ok := f.index[email]
	if !ok {
		return nil, user_domain.ErrNotFound
	}

	return user_domain.NewUser(email), nil
}

func (f *Repository) GetAllEmails(
	_ context.Context,
) ([]string, error) {
	f.fm.RLock()
	defer f.fm.RUnlock()

	if _, err := os.Stat(f.filePath); os.IsNotExist(err) {
		return []string{}, nil
	}

	data, err := os.ReadFile(f.filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file by path: %s", f.filePath)
	}

	rows := strings.Split(string(data), "\n")

	emails := make([]string, 0, len(rows))

	for _, row := range rows {
		if row == "" {
			continue
		}

		emails = append(emails, row)
	}

	return emails, nil
}

func (f *Repository) EmailExist(_ context.Context, email string) (bool, error) {
	f.fm.RLock()
	defer f.fm.RUnlock()

	if _, err := os.Stat(f.filePath); os.IsNotExist(err) {
		f.im.Lock()
		defer f.im.Unlock()

		f.index = make(map[string]struct{})

		return false, nil
	}

	f.im.RLock()
	defer f.im.RUnlock()

	_, ok := f.index[email]

	return ok, nil
}

func addEndOfTheLine(data string) string {
	return data + "\n"
}
