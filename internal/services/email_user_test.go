package services_test

import (
	"context"
	"errors"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/require"

	"exchange/internal/domain"
	"exchange/internal/repository/mem"
	"exchange/internal/services"
)

func TestCreateUser(t *testing.T) {
	srv := services.NewEmailUserService(context.Background(), mem.NewMemoryRepository())

	err := srv.NewEmailUser(context.Background(), domain.NewEmailUser(faker.Word()))

	require.NoError(t, err)
}

func TestCreateManyUsers(t *testing.T) {
	srv := services.NewEmailUserService(context.Background(), mem.NewMemoryRepository())
	batch := 10

	for i := 0; i < batch; i++ {
		err := srv.NewEmailUser(context.Background(), domain.NewEmailUser(faker.Name()))
		require.NoError(t, err)
	}
}

func TestCreateExistedUser(t *testing.T) {
	srv := services.NewEmailUserService(context.Background(), mem.NewMemoryRepository())

	email := faker.Word()
	ctx := context.Background()

	err := srv.NewEmailUser(ctx, domain.NewEmailUser(email))
	require.NoError(t, err)

	err = srv.NewEmailUser(ctx, domain.NewEmailUser(email))
	require.True(t, errors.Is(err, domain.ErrAlreadyExist))
}
