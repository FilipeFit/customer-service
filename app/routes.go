package app

import (
	"github.com/filipeFit/customer-service/controllers/customer"
	"github.com/filipeFit/customer-service/controllers/heath"
)

func mapRoutes() {
	router.GET("/health", heath.HealthCheck)
	router.POST("/customers", customer.CreateCustomer)
	router.GET("/customers", customer.FindAll)
	router.GET("/customers/:customerId", customer.FindCustomerById)
	router.PUT("/customers/:customerId", customer.Update)
}
