package persistence

import (
	"context"
	"testing"

	"github.com/Andressep/QuoteMaker/internal/pkg/util"
	"github.com/Andressep/QuoteMaker/internal/pkg/utiltest"
	"github.com/stretchr/testify/require"
)

func TestSaveQuotation(t *testing.T) {
	db := utiltest.SetupTestDB(t)

	// Omitido: Código para limpiar la tabla de cotizaciones antes de la prueba.

	repo := NewWriteQuotationRepository(db)
	quotation := utiltest.CreateRandomQuotation(t, db)

	// Intenta guardar la cotización
	savedQuotation, err := repo.SaveQuotation(context.Background(), quotation)
	require.NoError(t, err)
	require.NotEmpty(t, savedQuotation)

	// Verificaciones básicas para asegurar que la cotización se guardó correctamente
	require.Equal(t, quotation.TotalPrice, savedQuotation.TotalPrice)
	require.False(t, savedQuotation.IsPurchased)
	require.False(t, savedQuotation.IsDelivered)

	// Omitido: Verificaciones adicionales y limpieza después de la prueba.
}

func TestUpdateQuotation(t *testing.T) {
	db := utiltest.SetupTestDB(t)

	// Primero, crea y guarda una cotización para tener algo que actualizar.
	repo := NewWriteQuotationRepository(db)
	quotation := utiltest.CreateRandomQuotation(t, db)

	// Modifica algunos campos de la cotización guardada.
	quotation.IsPurchased = true
	quotation.TotalPrice = util.RandomFloat(1000, 5000) // Asume que tienes una función util.RandomFloat

	// Actualiza la cotización
	updatedQuotation, err := repo.UpdateQuotation(context.Background(), quotation)
	require.NoError(t, err)

	// Verificaciones para asegurar que la cotización se actualizó correctamente
	require.True(t, updatedQuotation.IsPurchased)
	require.Equal(t, quotation.TotalPrice, updatedQuotation.TotalPrice)

}
