package repositories

import (
	"github.com/filipeFit/customer-service/domain"
	"github.com/filipeFit/customer-service/domain/api"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
)

func prepareDatabase(t *testing.T) (error, api.CreateCustomerRequest, *customerRepository, *gorm.DB) {
	gormLogger := logger.Default.LogMode(logger.Info)
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{Logger: gormLogger})
	if err != nil {
		t.Fatal("not possible to open database ")
	}

	err = db.AutoMigrate(&domain.Customer{})
	if err != nil {
		t.Fatal("not possible to execute the migrations")
	}
	customer := api.CreateCustomerRequest{
		Name:       "Test",
		LeiCompany: "Itau",
		Document:   "123456789",
		Active:     true,
		Nick:       "test",
		Email:      "Test@Test.com"}

	customerRepository := NewCustomerRepository(db)
	return err, customer, customerRepository, db
}

func TestCustomerRepository_Create(t *testing.T) {
	err, customer, customerRepository, db := prepareDatabase(t)

	createdCustomer, err := customerRepository.Create(&customer)
	if err != nil {
		t.Fatal("fail to create a new customer")
	}

	assert.EqualValues(t, "test", createdCustomer.Nick)
	assert.EqualValues(t, "Test", createdCustomer.Name)
	assert.EqualValues(t, "Test@Test.com", createdCustomer.Email)
	assert.NotEqual(t, 0, createdCustomer.ID)
	assert.EqualValues(t, "Itau", createdCustomer.LeiCompany)
	assert.EqualValues(t, true, createdCustomer.Active)

	db.Exec("DELETE from customers")
}

func TestCustomerRepository_FindById(t *testing.T) {
	err, customer, customerRepository, db := prepareDatabase(t)

	createdCustomer, err := customerRepository.Create(&customer)
	if err != nil {
		t.Fatal("fail to create a new customer to query")
	}

	queryCustomer, err := customerRepository.FindById(createdCustomer.ID)
	if err != nil {
		t.Fatal("fail to Query customer")
	}

	assert.EqualValues(t, createdCustomer.ID, queryCustomer.ID)
	assert.EqualValues(t, createdCustomer.Nick, queryCustomer.Nick)
	assert.EqualValues(t, createdCustomer.Name, queryCustomer.Name)
	assert.EqualValues(t, createdCustomer.Email, queryCustomer.Email)
	assert.EqualValues(t, createdCustomer.Active, queryCustomer.Active)
	assert.EqualValues(t, createdCustomer.LeiCompany, queryCustomer.LeiCompany)

	db.Exec("DELETE from customers")
}

func TestCustomerRepository_FindAll(t *testing.T) {
	err, customer, customerRepository, db := prepareDatabase(t)

	_, err = customerRepository.Create(&customer)
	if err != nil {
		t.Fatal("fail to create a new customer to query")
	}

	customers, err := customerRepository.FindAll()
	assert.EqualValues(t, 1, len(customers))
	db.Exec("DELETE from customers")
}

func TestCustomerRepository_Update(t *testing.T) {
	err, customer, customerRepository, db := prepareDatabase(t)

	createdCustomer, err := customerRepository.Create(&customer)
	if err != nil {
		t.Fatal("fail to create a new customer to query")
	}
	customerRequest := api.CreateCustomerRequest{
		Name:       "TestUpdate",
		LeiCompany: "Bradesco",
		Document:   "77777777",
		Active:     false,
		Nick:       "testUpdate",
		Email:      "Test@update.com"}

	updatedCustomer, err := customerRepository.Update(createdCustomer.ID, &customerRequest)
	if err != nil {
		t.Fatal("error updating the customer")
	}

	assert.EqualValues(t, createdCustomer.ID, updatedCustomer.ID)
	assert.EqualValues(t, customerRequest.Nick, updatedCustomer.Nick)
	assert.EqualValues(t, customerRequest.Name, updatedCustomer.Name)
	assert.EqualValues(t, customerRequest.Email, updatedCustomer.Email)
	assert.EqualValues(t, customerRequest.Active, updatedCustomer.Active)
	assert.EqualValues(t, customerRequest.LeiCompany, updatedCustomer.LeiCompany)

	db.Exec("DELETE from customers")
}

func TestCustomerRepository_UpdateField(t *testing.T) {
	err, customer, customerRepository, db := prepareDatabase(t)

	createdCustomer, err := customerRepository.Create(&customer)
	if err != nil {
		t.Fatal("fail to create a new customer to query")
	}

	err = customerRepository.UpdateField(createdCustomer.ID, "active", false)
	if err != nil {
		t.Fatal("fail to update a customer field")
	}

	updatedCustomer, err := customerRepository.FindById(createdCustomer.ID)

	assert.EqualValues(
		t,
		createdCustomer.ID,
		updatedCustomer.ID,
	)
	assert.EqualValues(
		t,
		createdCustomer.Nick,
		updatedCustomer.Nick,
	)
	assert.EqualValues(
		t,
		createdCustomer.Name,
		updatedCustomer.Name,
	)
	assert.EqualValues(t, createdCustomer.Email, updatedCustomer.Email)
	assert.EqualValues(t, false, updatedCustomer.Active)
	assert.EqualValues(t, createdCustomer.LeiCompany, updatedCustomer.LeiCompany)

	db.Exec("DELETE from customers")
}
