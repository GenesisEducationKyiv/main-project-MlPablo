package filesysytem_test

import (
	"context"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"exchange/pkg/domain"
	"exchange/pkg/repository/filesysytem"
)

const testFilePath = "test.txt"

func TestFileSave(t *testing.T) {
	ctx := context.Background()

	repo, err := filesysytem.NewFileSystemRepository(testFilePath)
	require.NoError(t, err)

	defer os.Remove(testFilePath)

	testEmail := faker.Email()

	err = repo.SaveEmail(ctx, domain.NewEmailUser(testEmail))
	require.NoError(t, err)

	fileContent, err := os.ReadFile(testFilePath)
	require.NoError(t, err)

	assert.Equal(t, testEmail, strings.TrimSpace(string(fileContent)))
}

func TestSave(t *testing.T) {
	ctx := context.Background()

	repo, err := filesysytem.NewFileSystemRepository(testFilePath)
	require.NoError(t, err)

	defer os.Remove(testFilePath)

	testEmail := faker.Email()

	err = repo.SaveEmail(ctx, domain.NewEmailUser(testEmail))
	require.NoError(t, err)

	email, err := repo.GetByEmail(ctx, testEmail)
	require.NoError(t, err)

	assert.Equal(t, testEmail, email.Email)
}

func TestEmailExist(t *testing.T) {
	ctx := context.Background()
	batch := 10

	repo, err := filesysytem.NewFileSystemRepository(testFilePath)
	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(testFilePath)

	var exist bool

	for i := 0; i < batch; i++ {
		mail := faker.Email()
		err = repo.SaveEmail(ctx, domain.NewEmailUser(mail))
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

	repo, err := filesysytem.NewFileSystemRepository(testFilePath)
	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(testFilePath)

	emails := make([]string, batch)

	for i := 0; i < batch; i++ {
		mail := faker.Email()
		err = repo.SaveEmail(ctx, domain.NewEmailUser(mail))
		require.NoError(t, err)

		emails[i] = mail
	}

	getEmails, err := repo.GetAllEmails(ctx)
	require.NoError(t, err)
	require.True(t, reflect.DeepEqual(emails, getEmails), "slices elements are not equal")
}
