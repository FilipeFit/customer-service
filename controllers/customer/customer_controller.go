package customer

import (
	"github.com/filipeFit/customer-service/domain/api"
	"github.com/filipeFit/customer-service/handlers"
	"github.com/filipeFit/customer-service/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var (
	customerService = services.CustomerService
)

func CreateCustomer(c *gin.Context) {
	var request api.CreateCustomerRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := handlers.NewBadRequestError("invalid json body")
		c.JSON(apiErr.ResponseStatus(), apiErr)
		return
	}
	response, err := customerService.CreateCustomer(&request)
	if err != nil {
		c.JSON(err.ResponseStatus(), err)
		return
	}
	c.JSON(http.StatusCreated, response)
}

func FindCustomerById(c *gin.Context) {
	customerId, err := strconv.ParseUint(c.Param("customerId"), 10, 64)
	if err != nil {
		apiErr := handlers.NewBadRequestError("invalid customer id")
		c.JSON(apiErr.ResponseStatus(), apiErr)
		return
	}

	response, apiErr := customerService.FindCustomerById(customerId)
	if apiErr != nil {
		c.JSON(apiErr.ResponseStatus(), apiErr)
		return
	}
	c.JSON(http.StatusOK, response)
}

func FindAll(c *gin.Context) {
	response, apiErr := customerService.FindAll()
	if apiErr != nil {
		c.JSON(apiErr.ResponseStatus(), apiErr)
		return
	}
	c.JSON(http.StatusOK, response)
}

func Update(c *gin.Context) {
	var request api.CreateCustomerRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := handlers.NewBadRequestError("invalid json body")
		c.JSON(apiErr.ResponseStatus(), apiErr)
		return
	}
	customerId, err := strconv.ParseUint(c.Param("customerId"), 10, 64)
	if err != nil {
		apiErr := handlers.NewBadRequestError("invalid customer id")
		c.JSON(apiErr.ResponseStatus(), apiErr)
		return
	}
	response, apiErr := customerService.UpdateCustomer(customerId, &request)
	if apiErr != nil {
		c.JSON(apiErr.ResponseStatus(), apiErr)
		return
	}
	c.JSON(http.StatusOK, response)

}
