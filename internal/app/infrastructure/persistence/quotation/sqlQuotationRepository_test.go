package quotation

import (
	"context"
	"math/rand"
	"testing"
	"time"

	domain "github.com/Andressep/QuoteMaker/internal/app/domain/quotation"
	"github.com/Andressep/QuoteMaker/internal/pkg/util"
	"github.com/Andressep/QuoteMaker/internal/pkg/utiltest"
	"github.com/stretchr/testify/require"
)

func CreateRandomQuotation(t *testing.T) domain.Quotation {
	rand.Seed(time.Now().UnixNano())

	db := utiltest.SetupTestDB(t)
	ctx := context.Background()
	repo := NewQuotationRepository(db)

	quotation := domain.Quotation{
		CreatedAt:  time.Now(),
		TotalPrice: util.RandomFloat(1000, 10000),
	}

	savedQuotation, err := repo.SaveQuotation(ctx, quotation)

	require.NoError(t, err)
	require.NotEqual(t, 0, savedQuotation.ID) // Asegúrate de que se generó un ID
	require.NotEmpty(t, savedQuotation)
	require.Equal(t, quotation.TotalPrice, savedQuotation.TotalPrice)
	require.NotZero(t, savedQuotation.ID)

	return savedQuotation
}

func TestSaveQuotation(t *testing.T) {
	db := utiltest.SetupTestDB(t)
	ctx := context.Background()
	repo := NewQuotationRepository(db)

	CreateRandomQuotation(t)
	t.Cleanup(
		func() {
			err := repo.DeleteQuotation(ctx, CreateRandomQuotation(t).ID)
			require.NoError(t, err)
		},
	)
}

func TestGetQuotationByID(t *testing.T) {
	db := utiltest.SetupTestDB(t)
	ctx := context.Background()
	repo := NewQuotationRepository(db)
	newQuotation := CreateRandomQuotation(t)
	t.Cleanup(
		func() {
			err := repo.DeleteQuotation(ctx, CreateRandomQuotation(t).ID)
			require.NoError(t, err)
		},
	)

	fetchedQuotation, err := repo.GetQuotationByID(ctx, newQuotation.ID)
	require.NoError(t, err)
	require.NotNil(t, fetchedQuotation)
	require.Equal(t, newQuotation.ID, fetchedQuotation.ID)
	require.Equal(t, newQuotation.CreatedAt, fetchedQuotation.CreatedAt)
}

func TestListQuotation(t *testing.T) {
	db := utiltest.SetupTestDB(t)
	ctx := context.Background()
	repo := NewQuotationRepository(db)

	for i := 0; i < 5; i++ {
		CreateRandomQuotation(t)
	}

	quotations, err := repo.ListQuotations(ctx, 5, 0)
	require.NoError(t, err)

	for _, quotation := range quotations {
		require.NotEmpty(t, quotation)
		require.Len(t, quotations, 5)
	}
}

func TestDeleteQuotation(t *testing.T) {
	db := utiltest.SetupTestDB(t)
	ctx := context.Background()
	repo := NewQuotationRepository(db)
	newQuotation := CreateRandomQuotation(t)

	// delete quotation
	err := repo.DeleteQuotation(ctx, newQuotation.ID)
	require.NoError(t, err)

	//verify
	_, err = repo.GetQuotationByID(ctx, newQuotation.ID)
	require.Error(t, err)
}

func TestUpdateQuotation(t *testing.T) {
	db := utiltest.SetupTestDB(t)
	ctx := context.Background()
	repo := NewQuotationRepository(db)
	t.Cleanup(
		func() {
			err := repo.DeleteQuotation(ctx, CreateRandomQuotation(t).ID)
			require.NoError(t, err)
		},
	)
	originalQuotation := CreateRandomQuotation(t)

	// update

	err := repo.UpdateQuotation(ctx, originalQuotation)
	require.NoError(t, err)

	// verify
	updateQuotation, err := repo.GetQuotationByID(ctx, originalQuotation.ID)
	require.NoError(t, err)
	require.Equal(t, originalQuotation.TotalPrice, updateQuotation.TotalPrice)
}
