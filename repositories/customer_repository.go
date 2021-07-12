package repositories

import (
	"github.com/filipeFit/customer-service/domain"
	"github.com/filipeFit/customer-service/domain/api"
	"gorm.io/gorm"
)

type customerRepository struct {
	DB *gorm.DB
}

type customerRepositoryInterface interface {
	Create(request *api.CreateCustomerRequest) (*domain.Customer, error)
	FindById(id uint64) (*domain.Customer, error)
	FindAll() ([]domain.Customer, error)
	Update(customerId uint64, request *api.CreateCustomerRequest) (*domain.Customer, error)
	UpdateField(customerId uint64, field string, value interface{}) error
}

var (
	_ customerRepositoryInterface
)

func init() {
	_ = &customerRepository{}
}

func NewCustomerRepository(db *gorm.DB) *customerRepository {
	return &customerRepository{DB: db}
}

func (s *customerRepository) Create(request *api.CreateCustomerRequest) (*domain.Customer, error) {
	customer := api.ToCustomer(request)
	result := s.DB.Create(customer)
	if result.Error != nil {
		return nil, result.Error
	}

	return customer, nil
}

func (s *customerRepository) FindById(id uint64) (*domain.Customer, error) {
	var customer domain.Customer
	result := s.DB.First(&customer, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &customer, nil
}

func (s *customerRepository) FindAll() ([]domain.Customer, error) {
	var customers []domain.Customer
	result := s.DB.Find(&customers)
	if result.Error != nil {
		return nil, result.Error
	}
	return customers, nil
}

func (s *customerRepository) Update(customerId uint64, request *api.CreateCustomerRequest) (*domain.Customer, error) {
	customer := &domain.Customer{}
	result := s.DB.Model(customer).Where("id =?", customerId).Updates(api.ToCustomer(request))
	if result.Error != nil {
		return nil, result.Error
	}
	customer.ID = customerId
	return customer, nil
}

func (s *customerRepository) UpdateField(customerId uint64, field string, value interface{}) error {
	customer := &domain.Customer{}
	result := s.DB.Model(customer).Where("id =?", customerId).Update(field, value)
	if result.Error != nil {
		return result.Error
	}
	customer.ID = customerId
	return nil
}
