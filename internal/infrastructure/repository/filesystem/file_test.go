package filesystem_test

import (
	"context"
	"os"
	"reflect"
	"strings"
	"sync"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"exchange/internal/domain/user"
	"exchange/internal/infrastructure/repository/filesystem"
)

const testFilePath = "test.txt"

func TestFileSaveUserEmail(t *testing.T) {
	ctx := context.Background()

	repo, err := filesystem.NewFileSystemRepository(testFilePath)
	require.NoError(t, err)

	defer os.Remove(testFilePath)

	testEmail := faker.Email()

	err = repo.Save(ctx, user.NewUser(testEmail))
	require.NoError(t, err)

	fileContent, err := os.ReadFile(testFilePath)
	require.NoError(t, err)

	assert.Equal(t, testEmail, strings.TrimSpace(string(fileContent)))
}

func TestEmailExist(t *testing.T) {
	ctx := context.Background()
	batch := 10

	repo, err := filesystem.NewFileSystemRepository(testFilePath)
	require.NoError(t, err)

	defer os.Remove(testFilePath)

	var exist bool

	for i := 0; i < batch; i++ {
		mail := faker.Email()
		err = repo.Save(ctx, user.NewUser(mail))
		require.NoError(t, err)

		exist, err = repo.EmailExist(ctx, mail)
		require.NoError(t, err)
		require.True(t, exist)
	}

	ok, err := repo.EmailExist(ctx, faker.Email())
	require.NoError(t, err)
	require.False(t, ok)
}

func TestGetAll(t *testing.T) {
	ctx := context.Background()
	batch := 20

	repo, err := filesystem.NewFileSystemRepository(testFilePath)
	require.NoError(t, err)

	defer os.Remove(testFilePath)

	emails := make([]string, batch)

	for i := 0; i < batch; i++ {
		mail := faker.Email()
		err = repo.Save(ctx, user.NewUser(mail))
		require.NoError(t, err)

		emails[i] = mail
	}

	getEmails, err := repo.GetAllEmails(ctx)
	require.NoError(t, err)
	require.True(t, reflect.DeepEqual(emails, getEmails), "slices elements are not equal")
}

func TestConcurrentWrite(t *testing.T) {
	ctx := context.Background()

	repo, err := filesystem.NewFileSystemRepository(testFilePath)
	require.NoError(t, err)

	defer os.Remove(testFilePath)

	batch := 1000
	wg := sync.WaitGroup{}
	wg.Add(batch)

	emailCh := make(chan string)
	defer close(emailCh)

	for i := 0; i < batch; i++ {
		go func() {
			defer wg.Done()

			err = repo.Save(ctx, user.NewUser("123"))
			require.NoError(t, err)
		}()
	}

	wg.Wait()

	getEmails, err := repo.GetAllEmails(ctx)
	require.NoError(t, err)
	require.Len(
		t,
		getEmails,
		batch,
	)
}
