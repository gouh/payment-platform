package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"payment-platform/internal/container"
	"payment-platform/internal/dao"
	"payment-platform/internal/models"
	"payment-platform/internal/requests"
	"payment-platform/internal/responses"
)

type (
	CustomerHandlerInterface interface {
		GetCustomers(c *gin.Context)
		GetCustomerById(c *gin.Context)
		CreateCustomer(c *gin.Context)
		UpdateCustomer(c *gin.Context)
		DeleteCustomer(c *gin.Context)
	}
	CustomerHandler struct {
		CustomerDao dao.CustomerDaoInterface
	}
)

func (handler *CustomerHandler) GetCustomers(c *gin.Context) {
	customersParams := requests.CustomerParams{}
	errQueryParams := customersParams.QueryParamsToStruct(c.Request.URL.Query(), &customersParams)
	if errQueryParams != nil {
		c.JSON(http.StatusInternalServerError, responses.GetErrorResponse(errQueryParams.Error()))
		return
	}

	customers, count, errGetCustomers := handler.CustomerDao.GetCustomers(customersParams)
	if errGetCustomers != nil {
		_ = c.Error(errGetCustomers)
		c.JSON(http.StatusInternalServerError, responses.GetErrorResponse(errGetCustomers.Error()))
		return
	}

	c.JSON(http.StatusOK, responses.GetCustomersResponse(customers, count, customersParams.PaginationRequest))
}

func (handler *CustomerHandler) GetCustomerById(c *gin.Context) {
	customer, errGetCustomer := handler.CustomerDao.GetCustomerById(c.Param("id"))
	if errGetCustomer == nil && customer == nil {
		c.JSON(http.StatusNotFound, responses.GetErrorResponse("Not found"))
		return
	}

	if errGetCustomer != nil {
		c.JSON(http.StatusInternalServerError, responses.GetErrorResponse("Error connecting to database"+errGetCustomer.Error()))
		return
	}

	c.JSON(http.StatusOK, responses.GetCustomerResponse(customer))
}

func (handler *CustomerHandler) CreateCustomer(c *gin.Context) {
	var customer = &models.Customer{}
	if errValidation := c.BindJSON(customer); errValidation != nil {
		_ = c.Error(errValidation)
		c.JSON(http.StatusNotAcceptable, errValidation)
		return
	}

	customer, errorInsertCustomer := handler.CustomerDao.CreateCustomer(*customer)
	if errorInsertCustomer != nil {
		c.Status(http.StatusInternalServerError)
		_ = c.Error(errorInsertCustomer)
		return
	}

	c.JSON(http.StatusOK, responses.GetCustomerResponse(customer))
}

func (handler *CustomerHandler) UpdateCustomer(c *gin.Context) {
	var customer = &models.Customer{}
	if errValidation := c.BindJSON(customer); errValidation != nil {
		_ = c.Error(errValidation)
		c.JSON(http.StatusNotAcceptable, errValidation)
		return
	}

	customerId, errUuid := uuid.Parse(c.Param("id"))
	if errUuid != nil {
		c.Status(http.StatusInternalServerError)
		_ = c.Error(errUuid)
		return
	}

	customer.ID = customerId
	customer, errorUpdateCustomer := handler.CustomerDao.UpdateCustomer(*customer)
	if errorUpdateCustomer != nil {
		c.Status(http.StatusInternalServerError)
		_ = c.Error(errorUpdateCustomer)
		return
	}

	c.JSON(http.StatusOK, responses.GetCustomerResponse(customer))
}

func (handler *CustomerHandler) DeleteCustomer(c *gin.Context) {
	errorDeleteCustomer := handler.CustomerDao.DeleteCustomer(c.Param("id"))
	if errorDeleteCustomer != nil {
		if errorDeleteCustomer.Error() == "no rows affected" {
			c.Status(http.StatusNotFound)
		} else {
			c.Status(http.StatusInternalServerError)
		}
		_ = c.Error(errorDeleteCustomer)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func NewCustomerHandler(container *container.Container) CustomerHandlerInterface {
	return &CustomerHandler{
		CustomerDao: dao.NewCustomerDao(container),
	}
}
