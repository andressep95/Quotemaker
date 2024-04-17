package persistence

import (
	"context"
	"fmt"
	"testing"

	"github.com/Andressep/QuoteMaker/internal/pkg/utiltest"
	"github.com/stretchr/testify/require"
)

/*func TestGetQuotationByID(t *testing.T) {
	db := utiltest.SetupTestDB(t)
	ctx := context.Background()
	readRepo := NewReadQuotationRepository(db)
	newQuotation := utiltest.CreateRandomQuotation(t, db)

	quotation, err := readRepo.GetQuotationByID(ctx, newQuotation.ID)
	require.NoError(t, err)
	require.NotEmpty(t, quotation)
}*/

func TestListQuotations(t *testing.T) {
	db := utiltest.SetupTestDB(t)
	ctx := context.Background()
	readRepo := NewReadQuotationRepository(db)

	for i := 0; i < 5; i++ {
		newQuotation := utiltest.CreateRandomQuotation(t, db)
		fmt.Println(i, newQuotation)
		require.NotEmpty(t, newQuotation.Products)
	}

	quotations, err := readRepo.ListQuotations(ctx, 5, 0)
	require.NoError(t, err)
	require.Len(t, quotations, 5)
}
