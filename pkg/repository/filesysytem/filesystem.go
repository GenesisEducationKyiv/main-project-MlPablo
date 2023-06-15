package filesysytem

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync"

	"exchange/pkg/domain"
)

// I have decided to use a memory index for get operations for.
// This was done, because simple get operation is too heavy
// where we need to read the whole file and then iterate over every email (O(n)).
// On the big amount of data this can lead to performance issues.
// So better if we will index the file on the startup of the program
// and then we will add new items in the file and index.
// Get All operation made with file, just because I want to show read operation with filesystem)
// Don't forget about locks.
type fileSystemRepository struct {
	filePath string
	index    map[string]struct{}
	// file mutex
	fm sync.RWMutex
	// index mutex
	im sync.RWMutex
}

func NewFileSystemRepository(filePath string) (domain.EmailRepository, error) {
	f := &fileSystemRepository{
		filePath: filePath,
		fm:       sync.RWMutex{},
		im:       sync.RWMutex{},
	}

	if err := f.loadIndex(); err != nil {
		return nil, err
	}

	return f, nil
}

func (f *fileSystemRepository) SaveEmail(_ context.Context, eu *domain.EmailUser) error {
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

func (f *fileSystemRepository) GetByEmail(
	_ context.Context,
	email string,
) (*domain.EmailUser, error) {
	f.im.RLock()
	defer f.im.RUnlock()

	_, ok := f.index[email]
	if !ok {
		return nil, domain.ErrNotFound
	}

	return domain.NewEmailUser(email), nil
}

func (f *fileSystemRepository) GetAllEmails(
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

func (f *fileSystemRepository) EmailExist(_ context.Context, email string) (bool, error) {
	f.im.RLock()
	defer f.im.RUnlock()

	_, ok := f.index[email]

	return ok, nil
}

func addEndOfTheLine(data string) string {
	return data + "\n"
}
