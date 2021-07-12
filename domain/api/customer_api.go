package api

import "github.com/filipeFit/customer-service/domain"

type CreateCustomerRequest struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Nick       string `json:"nick"`
	Document   string `json:"document"`
	Active     bool   `json:"active"`
	LeiCompany string `json:"leiCompany"`
}

type CreateCustomerResponse struct {
	ID         uint64 `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Nick       string `json:"nick"`
	Document   string `json:"document"`
	Active     bool   `json:"active"`
	LeiCompany string `json:"leiCompany"`
}

func ToCustomer(request *CreateCustomerRequest) *domain.Customer {
	customer := domain.Customer{
		Email:      request.Email,
		Name:       request.Name,
		Nick:       request.Nick,
		Active:     request.Active,
		Document:   request.Document,
		LeiCompany: request.LeiCompany,
	}

	return &customer
}

func ToCustomerResponse(customer *domain.Customer) *CreateCustomerResponse {
	customerResponse := CreateCustomerResponse{
		ID:         customer.ID,
		Email:      customer.Email,
		Name:       customer.Name,
		Nick:       customer.Nick,
		Active:     customer.Active,
		Document:   customer.Document,
		LeiCompany: customer.LeiCompany,
	}

	return &customerResponse
}
