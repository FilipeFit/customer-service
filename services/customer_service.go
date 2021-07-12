package services

import (
	"github.com/filipeFit/customer-service/config"
	"github.com/filipeFit/customer-service/domain/api"
	"github.com/filipeFit/customer-service/handlers"
	"github.com/filipeFit/customer-service/repositories"
	"net/http"
)

type customerService struct{}

type customerServiceInterface interface {
	CreateCustomer(request *api.CreateCustomerRequest) (*api.CreateCustomerResponse, handlers.ApiError)
	UpdateCustomer(customerId uint64, request *api.CreateCustomerRequest) (*api.CreateCustomerResponse, handlers.ApiError)
	FindCustomerById(customerId uint64) (*api.CreateCustomerResponse, handlers.ApiError)
	FindAll() ([]api.CreateCustomerResponse, handlers.ApiError)
}

var (
	CustomerService    customerServiceInterface
	customerRepository = repositories.NewCustomerRepository(config.DB)
)

func init() {
	CustomerService = &customerService{}
}

func (s *customerService) CreateCustomer(request *api.CreateCustomerRequest) (*api.CreateCustomerResponse, handlers.ApiError) {
	customer, err := customerRepository.Create(request)
	if err != nil {
		return nil, handlers.NewApiError(http.StatusInternalServerError, "error in saving the customer")
	}
	response := api.ToCustomerResponse(customer)
	return response, nil
}

func (s *customerService) FindCustomerById(customerId uint64) (*api.CreateCustomerResponse, handlers.ApiError) {
	customer, err := customerRepository.FindById(customerId)
	if err != nil {
		return nil, handlers.NewApiError(http.StatusNotFound, "customer not found")
	}
	response := api.ToCustomerResponse(customer)
	return response, nil
}

func (s *customerService) FindAll() ([]api.CreateCustomerResponse, handlers.ApiError) {
	customers, err := customerRepository.FindAll()
	if err != nil {
		return nil, handlers.NewApiError(http.StatusInternalServerError, "error searching for customers")
	}

	var customersResponse []api.CreateCustomerResponse
	for _, customer := range customers {
		customerResponse := api.ToCustomerResponse(&customer)
		customersResponse = append(customersResponse, *customerResponse)
	}

	return customersResponse, nil
}

func (s *customerService) UpdateCustomer(customerId uint64, request *api.CreateCustomerRequest) (*api.CreateCustomerResponse, handlers.ApiError) {
	_, err := customerRepository.FindById(customerId)
	if err != nil {
		return nil, handlers.NewApiError(http.StatusNotFound, "customer not found")
	}
	updatedCustomer, err := customerRepository.Update(customerId, request)
	if err != nil {
		return nil, handlers.NewApiError(http.StatusBadRequest, "Error updating the customer")
	}
	response := api.ToCustomerResponse(updatedCustomer)
	return response, nil
}
