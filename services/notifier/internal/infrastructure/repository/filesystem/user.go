package filesystem

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"notifier/internal/domain/user"
)

func (f *Repository) Save(_ context.Context, eu *user.User) error {
	f.fm.Lock()
	defer f.fm.Unlock()

	file, err := os.OpenFile(f.cfg.Path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return fmt.Errorf("failed to open data file: %w", err)
	}
	defer file.Close()

	if _, err = file.WriteString(addEndOfTheLine(eu.Email)); err != nil {
		return fmt.Errorf("failed write to file")
	}

	return nil
}

func (f *Repository) GetAllEmails(
	_ context.Context,
) ([]string, error) {
	f.fm.RLock()
	defer f.fm.RUnlock()

	data, err := os.ReadFile(f.cfg.Path)
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}, nil
		}

		return nil, fmt.Errorf("failed to read file by path: %s", f.cfg.Path)
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

	file, err := os.Open(f.cfg.Path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	r := bufio.NewReader(file)

	for {
		line, _, err := r.ReadLine() //nolint:govet // shadow here is ok
		if err != nil {
			if errors.Is(err, io.EOF) {
				return false, nil
			}

			return false, err
		}

		if string(line) == (email) {
			return true, nil
		}
	}
}

func (f *Repository) Delete(_ context.Context, email string) error {
	f.fm.Lock()
	defer f.fm.Unlock()

	file, err := os.Open(f.cfg.Path)
	if err != nil {
		return err
	}

	scann := bufio.NewScanner(file)
	buffer := bytes.NewBufferString("")

	for scann.Scan() {
		if scann.Text() == email {
			continue
		}

		if _, err = buffer.WriteString(addEndOfTheLine(scann.Text())); err != nil {
			return err
		}
	}

	if err = scann.Err(); err != nil {
		return err
	}

	return os.WriteFile(f.cfg.Path, buffer.Bytes(), 0600)
}

func addEndOfTheLine(data string) string {
	return data + "\n"
}
