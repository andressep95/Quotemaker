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

func TestDeleteQuotation(t *testing.T) {
	db := utiltest.SetupTestDB(t)
	defer db.Close()

	// Asegura la creación del repositorio y cualquier otro setup necesario.
	writeRepo := NewWriteQuotationRepository(db)
	readRepo := NewReadQuotationRepository(db)

	// Primero, crea una cotización para luego eliminarla.
	quotation := utiltest.CreateRandomQuotation(t, db)
	require.NotEmpty(t, quotation.ID, "La cotización debe tener un ID válido")

	// Intenta eliminar la cotización creada.
	err := writeRepo.DeleteQuotation(context.Background(), quotation.ID)
	require.NoError(t, err, "No debería haber un error al eliminar una cotización existente")

	// Verifica que la cotización se haya eliminado correctamente.
	_, err = readRepo.GetQuotationByID(context.Background(), quotation.ID)
	require.Error(t, err, "Debería haber un error al intentar obtener una cotización eliminada")

	// Prueba eliminar una cotización que no existe y verifica el error.
	fakeUUID := "123e4567-e89b-12d3-a456-426614174000"
	err = writeRepo.DeleteQuotation(context.Background(), fakeUUID)
	require.Error(t, err, "Debería haber un error al intentar eliminar una cotización que no existe")
}
