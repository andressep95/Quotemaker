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

func CreateRandomQuotation(t *testing.T) domain.Quotation {
	rand.Seed(time.Now().UnixNano())

	quotation := domain.Quotation{
		SellerID:   CreateRandomSeller(t).ID,
		CustomerID: CreateRandomCustomer(t).ID,
		CreatedAt:  time.Now(),
		TotalPrice: util.RandomFloat(1000, 10000),
	}

	db := util.SetupTestDB(t)
	ctx := context.Background()
	repo := NewQuotationRepository(db)

	savedQuotation, err := repo.SaveQuotation(ctx, quotation)

	require.NoError(t, err)
	require.NotEqual(t, 0, savedQuotation.ID) // Asegúrate de que se generó un ID
	require.NotEmpty(t, savedQuotation)
	require.Equal(t, quotation.SellerID, savedQuotation.SellerID)
	require.Equal(t, quotation.CustomerID, savedQuotation.CustomerID)
	require.Equal(t, quotation.TotalPrice, savedQuotation.TotalPrice)
	require.NotZero(t, savedQuotation.ID)

	return quotation
}

func TestSaveQuotation(t *testing.T) {
	CreateRandomQuotation(t)
}
