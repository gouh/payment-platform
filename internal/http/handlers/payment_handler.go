package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
	"payment-platform/internal/container"
	"payment-platform/internal/dao"
	"payment-platform/internal/models"
	"payment-platform/internal/requests"
	"payment-platform/internal/responses"
	"payment-platform/internal/utils"
	"time"
)

type (
	PaymentHandlerInterface interface {
		GetPayments(c *gin.Context)
		GetPaymentById(c *gin.Context)
		CreatePayment(c *gin.Context)
	}
	PaymentHandler struct {
		PaymentDao    dao.PaymentDaoInterface
		BankSimulator utils.BankSimulatorInterface
	}
)

func (handler *PaymentHandler) GetPayments(c *gin.Context) {
	paymentsParams := requests.PaymentParams{}
	errQueryParams := paymentsParams.QueryParamsToStruct(c.Request.URL.Query(), &paymentsParams)
	if errQueryParams != nil {
		c.JSON(http.StatusInternalServerError, responses.GetErrorResponse(errQueryParams.Error()))
		return
	}

	payments, count, errGetPayments := handler.PaymentDao.GetPayments(paymentsParams)
	if errGetPayments != nil {
		c.Status(http.StatusInternalServerError)
		_ = c.Error(errGetPayments)
		return
	}

	c.JSON(http.StatusOK, responses.GetPaymentsResponse(payments, count, paymentsParams.PaginationRequest))
}

func (handler *PaymentHandler) GetPaymentById(c *gin.Context) {
	payment, errGetPayment := handler.PaymentDao.GetPaymentById(c.Param("id"))
	if errGetPayment == nil && payment == nil {
		c.JSON(http.StatusNotFound, responses.GetErrorResponse("Not found"))
		return
	}

	if errGetPayment != nil {
		c.JSON(http.StatusInternalServerError, responses.GetErrorResponse("Error connecting to database"+errGetPayment.Error()))
		return
	}

	c.JSON(http.StatusOK, responses.GetPaymentResponse(payment))
}

func (handler *PaymentHandler) CreatePayment(c *gin.Context) {
	var payment = &models.Payment{}
	if errValidation := c.ShouldBind(payment); errValidation != nil {
		c.Status(http.StatusNotAcceptable)
		_ = c.Error(errValidation)
		return
	}

	payment.ID = uuid.New()
	payment.TransactionDate = time.Now()
	errAuth := handler.BankSimulator.AuthorizePayment(payment.ID.String())
	if errAuth != nil {
		payment.Status = "unauthorized"
		logrus.Error(errAuth)
	} else {
		payment.Status = "authorized"
	}

	payment, errorInsertPayment := handler.PaymentDao.CreatePayment(*payment)
	if errorInsertPayment != nil {
		c.Status(http.StatusInternalServerError)
		_ = c.Error(errorInsertPayment)
		return
	}

	c.JSON(http.StatusOK, responses.GetPaymentResponse(payment))
}

func NewPaymentHandler(container *container.Container) PaymentHandlerInterface {
	return &PaymentHandler{
		PaymentDao:    dao.NewPaymentDao(container),
		BankSimulator: utils.NewBankSimulator(),
	}
}
