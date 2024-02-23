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
	MerchantHandlerInterface interface {
		GetMerchants(c *gin.Context)
		GetMerchantById(c *gin.Context)
		CreateMerchant(c *gin.Context)
		UpdateMerchant(c *gin.Context)
		DeleteMerchant(c *gin.Context)
	}
	MerchantHandler struct {
		MerchantDao dao.MerchantDaoInterface
	}
)

func (handler *MerchantHandler) GetMerchants(c *gin.Context) {
	merchantsParams := requests.MerchantParams{}
	errQueryParams := merchantsParams.QueryParamsToStruct(c.Request.URL.Query(), &merchantsParams)
	if errQueryParams != nil {
		c.JSON(http.StatusInternalServerError, responses.GetErrorResponse(errQueryParams.Error()))
		return
	}

	merchants, count, errGetMerchants := handler.MerchantDao.GetMerchants(merchantsParams)
	if errGetMerchants != nil {
		c.Status(http.StatusInternalServerError)
		_ = c.Error(errGetMerchants)
		return
	}

	c.JSON(http.StatusOK, responses.GetMerchantsResponse(merchants, count, merchantsParams.PaginationRequest))
}

func (handler *MerchantHandler) GetMerchantById(c *gin.Context) {
	merchant, errGetMerchant := handler.MerchantDao.GetMerchantById(c.Param("id"))
	if errGetMerchant == nil && merchant == nil {
		c.JSON(http.StatusNotFound, responses.GetErrorResponse("Not found"))
		return
	}

	if errGetMerchant != nil {
		c.JSON(http.StatusInternalServerError, responses.GetErrorResponse("Error connecting to database"+errGetMerchant.Error()))
		return
	}

	c.JSON(http.StatusOK, responses.GetMerchantResponse(merchant))
}

func (handler *MerchantHandler) CreateMerchant(c *gin.Context) {
	var merchant = &models.Merchant{}
	if errValidation := c.ShouldBind(merchant); errValidation != nil {
		c.Status(http.StatusNotAcceptable)
		_ = c.Error(errValidation)
		return
	}

	merchant, errorInsertMerchant := handler.MerchantDao.CreateMerchant(*merchant)
	if errorInsertMerchant != nil {
		c.Status(http.StatusInternalServerError)
		_ = c.Error(errorInsertMerchant)
		return
	}

	c.JSON(http.StatusOK, responses.GetMerchantResponse(merchant))
}

func (handler *MerchantHandler) UpdateMerchant(c *gin.Context) {
	var merchant = &models.Merchant{}
	if errValidation := c.ShouldBind(merchant); errValidation != nil {
		c.Status(http.StatusNotAcceptable)
		_ = c.Error(errValidation)
		return
	}

	merchantId, errUuid := uuid.Parse(c.Param("id"))
	if errUuid != nil {
		c.Status(http.StatusInternalServerError)
		_ = c.Error(errUuid)
		return
	}

	merchant.ID = merchantId
	merchant, errorUpdateMerchant := handler.MerchantDao.UpdateMerchant(*merchant)
	if errorUpdateMerchant != nil {
		c.Status(http.StatusInternalServerError)
		_ = c.Error(errorUpdateMerchant)
		return
	}

	c.JSON(http.StatusOK, responses.GetMerchantResponse(merchant))
}

func (handler *MerchantHandler) DeleteMerchant(c *gin.Context) {
	errorDeleteMerchant := handler.MerchantDao.DeleteMerchant(c.Param("id"))
	if errorDeleteMerchant != nil {
		if errorDeleteMerchant.Error() == "no rows affected" {
			c.Status(http.StatusNotFound)
		} else {
			c.Status(http.StatusInternalServerError)
		}
		_ = c.Error(errorDeleteMerchant)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func NewMerchantHandler(container *container.Container) MerchantHandlerInterface {
	return &MerchantHandler{
		MerchantDao: dao.NewMerchantDao(container),
	}
}
