package customer

import (
	"context"
	"math/rand"
	"testing"
	"time"

	domain "github.com/Andressep/QuoteMaker/internal/app/domain/customer"
	"github.com/Andressep/QuoteMaker/internal/pkg/util"
	"github.com/Andressep/QuoteMaker/internal/pkg/utiltest"
	"github.com/stretchr/testify/require"
)

func CreateRandomCustomer(t *testing.T) domain.Customer {
	rand.Seed(time.Now().UnixNano())

	customer := domain.Customer{
		Name:    util.RandomString(5),
		Rut:     util.RandomString(10),
		Address: util.RandomString(15),
		Phone:   util.RandomString(9),
		Email:   util.RandomString(5) + "@gmail.com",
	}

	db := utiltest.SetupTestDB(t)
	ctx := context.Background()
	repo := NewCustomerRepository(db)

	savedCustomer, err := repo.SaveCustomer(ctx, customer)

	require.NoError(t, err)
	require.NotEqual(t, 0, savedCustomer.ID)
	require.NotEmpty(t, savedCustomer)
	require.Equal(t, customer.Email, savedCustomer.Email)
	require.Equal(t, customer.Rut, savedCustomer.Rut)
	require.NotZero(t, savedCustomer.ID)

	return savedCustomer
}

func TestSaveCustomer(t *testing.T) {
	CreateRandomCustomer(t)
}

func TestGetCustomerByID(t *testing.T) {
	db := utiltest.SetupTestDB(t)
	ctx := context.Background()
	repo := NewCustomerRepository(db)
	newCustomer := CreateRandomCustomer(t)

	fetchedCustomer, err := repo.GetCustomerByID(ctx, newCustomer.ID)
	require.NoError(t, err)
	require.NotNil(t, fetchedCustomer)
	require.Equal(t, newCustomer.ID, fetchedCustomer.ID)
	require.Equal(t, newCustomer.Name, fetchedCustomer.Name)
}

func TestListCustomers(t *testing.T) {
	db := utiltest.SetupTestDB(t)
	ctx := context.Background()
	repo := NewCustomerRepository(db)

	for i := 0; i < 5; i++ {
		CreateRandomCustomer(t)
	}

	customers, err := repo.ListCustomers(ctx, 5, 0)
	require.NoError(t, err)

	for _, customer := range customers {
		require.NotEmpty(t, customer)
		require.Len(t, customers, 5)
	}
}

func TestDeleteCustomer(t *testing.T) {
	db := utiltest.SetupTestDB(t)
	ctx := context.Background()
	repo := NewCustomerRepository(db)
	newCustomer := CreateRandomCustomer(t)

	// delete product
	err := repo.DeleteCustomer(ctx, int(newCustomer.ID))
	require.NoError(t, err)

	//verify
	_, err = repo.GetCustomerByID(ctx, int(newCustomer.ID))
	require.Error(t, err)
}

func TestUpdateCustomer(t *testing.T) {
	db := utiltest.SetupTestDB(t)
	ctx := context.Background()
	repo := NewCustomerRepository(db)
	originalCustomer := CreateRandomCustomer(t)

	// Update
	originalCustomer.Name = "New Customer Name"
	originalCustomer.Address = "Avda. Michigan"
	err := repo.UpdateCustomer(ctx, originalCustomer)
	require.NoError(t, err)

	// Verify
	updateCustomer, err := repo.GetCustomerByID(ctx, originalCustomer.ID)
	require.NoError(t, err)
	require.Equal(t, originalCustomer.Name, updateCustomer.Name)
	require.Equal(t, originalCustomer.Address, updateCustomer.Address)
}