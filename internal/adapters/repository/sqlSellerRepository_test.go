package repository

import (
	"context"
	"math/rand"
	"testing"
	"time"

	domain "github.com/Andressep/QuoteMaker/internal/core/domain/entity"
	"github.com/Andressep/QuoteMaker/internal/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomSeller(t *testing.T) domain.Seller {
	rand.Seed(time.Now().UnixNano())

	seller := domain.Seller{
		Name: util.RandomString(10),
	}

	db := util.SetupTestDB(t)
	ctx := context.Background()
	repo := NewSellerRepository(db)

	savedSeller, err := repo.SaveSeller(ctx, seller)

	require.NoError(t, err)
	require.NotEqual(t, 0, savedSeller.ID)
	require.NotEmpty(t, savedSeller)
	require.Equal(t, seller.Name, savedSeller.Name)
	require.NotZero(t, savedSeller.ID)

	return savedSeller
}

func TestCreateSeller(t *testing.T) {
	CreateRandomSeller(t)
}

func TestGetSellerByID(t *testing.T) {
	db := util.SetupTestDB(t)
	ctx := context.Background()
	repo := NewSellerRepository(db)
	newSeller := CreateRandomSeller(t)

	savedSeller, err := repo.SaveSeller(ctx, newSeller)
	require.NoError(t, err)

	fetchedSeller, err := repo.GetSellerByID(ctx, savedSeller.ID)
	require.NoError(t, err)
	require.NotNil(t, fetchedSeller)
	require.Equal(t, savedSeller.ID, fetchedSeller.ID)
	require.Equal(t, savedSeller.Name, fetchedSeller.Name)
}

func TestListSellers(t *testing.T) {
	db := util.SetupTestDB(t)
	ctx := context.Background()
	repo := NewSellerRepository(db)

	for i := 0; i < 5; i++ {
		CreateRandomSeller(t)
	}

	sellers, err := repo.ListSellers(ctx, 5, 0)
	require.NoError(t, err)

	for _, seller := range sellers {
		require.NotEmpty(t, seller)
		require.Len(t, sellers, 5)
	}
}

func TestDeleteSeller(t *testing.T) {
	db := util.SetupTestDB(t)
	ctx := context.Background()
	repo := NewSellerRepository(db)
	newSeller := CreateRandomSeller(t)

	// delete product
	err := repo.DeleteSeller(ctx, int(newSeller.ID))
	require.NoError(t, err)

	//verify
	_, err = repo.GetSellerByID(ctx, int(newSeller.ID))
	require.Error(t, err)
}
