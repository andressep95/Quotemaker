package seller

import (
	"context"
	"math/rand"
	"testing"
	"time"

	domain "github.com/Andressep/QuoteMaker/internal/app/domain/seller"
	"github.com/Andressep/QuoteMaker/internal/pkg/util"
	"github.com/Andressep/QuoteMaker/internal/pkg/utiltest"
	"github.com/stretchr/testify/require"
)

func CreateRandomSeller(t *testing.T) domain.Seller {
	rand.Seed(time.Now().UnixNano())

	seller := domain.Seller{
		Name: util.RandomString(10),
	}

	db := utiltest.SetupTestDB(t)
	ctx := context.Background()
	repo := NewSellerRepository(db)

	// Inicia una transacción.
	tx, err := db.BeginTx(ctx, nil)
	require.NoError(t, err)
	defer tx.Rollback() // Asegúrate de revertir los cambios al final de la función.

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
	db := utiltest.SetupTestDB(t)
	ctx := context.Background()
	repo := NewSellerRepository(db)
	newSeller := CreateRandomSeller(t)

	fetchedSeller, err := repo.GetSellerByID(ctx, newSeller.ID)
	require.NoError(t, err)
	require.NotNil(t, fetchedSeller)
	require.Equal(t, newSeller.ID, fetchedSeller.ID)
	require.Equal(t, newSeller.Name, fetchedSeller.Name)
}

func TestListSellers(t *testing.T) {
	db := utiltest.SetupTestDB(t)
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
	db := utiltest.SetupTestDB(t)
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

func TestUpdateSeller(t *testing.T) {
	db := utiltest.SetupTestDB(t)
	ctx := context.Background()
	repo := NewSellerRepository(db)
	originalSeller := CreateRandomSeller(t)

	// update
	originalSeller.Name = "New original Name"
	err := repo.UpdateSeller(ctx, originalSeller)
	require.NoError(t, err)

	updateSeller, err := repo.GetSellerByID(ctx, originalSeller.ID)
	require.NoError(t, err)
	require.Equal(t, originalSeller.Name, updateSeller.Name)
}
