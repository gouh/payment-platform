package handlers

import (
	"braces.dev/errtrace"
	"github.com/gin-gonic/gin"
	"net/http"
	"payment-platform/internal/container"
	"payment-platform/internal/dao"
	"payment-platform/internal/requests"
	"payment-platform/internal/responses"
	"payment-platform/internal/utils"
)

type (
	TokenizedCardHandlerInterface interface {
		GetTokenizedCards(c *gin.Context)
		GetTokenizedCardById(c *gin.Context)
		CreateTokenizedCard(c *gin.Context)
		DeleteTokenizedCard(c *gin.Context)
	}
	TokenizedCardHandler struct {
		TokenizedCardDao dao.TokenizedCardDaoInterface
	}
)

func (handler *TokenizedCardHandler) GetTokenizedCards(c *gin.Context) {
	merchantsParams := requests.TokenizedCardParams{}
	errQueryParams := merchantsParams.QueryParamsToStruct(c.Request.URL.Query(), &merchantsParams)
	if errQueryParams != nil {
		c.JSON(http.StatusInternalServerError, responses.GetErrorResponse(errQueryParams.Error()))
		return
	}

	customerId := c.Param("id")
	merchantsParams.CustomerId = &customerId
	merchants, count, errGetTokenizedCards := handler.TokenizedCardDao.GetTokenizedCards(merchantsParams)
	if errGetTokenizedCards != nil {
		_ = c.Error(errGetTokenizedCards)
		c.JSON(http.StatusInternalServerError, responses.GetErrorResponse(errGetTokenizedCards.Error()))
		return
	}

	c.JSON(http.StatusOK, responses.GetTokenizedCardsResponse(merchants, count, merchantsParams.PaginationRequest))
}

func (handler *TokenizedCardHandler) GetTokenizedCardById(c *gin.Context) {
	merchant, errGetTokenizedCard := handler.TokenizedCardDao.GetTokenizedCardById(c.Param("id"), c.Param("token"))
	if errGetTokenizedCard == nil && merchant == nil {
		c.JSON(http.StatusNotFound, responses.GetErrorResponse("Not found"))
		return
	}

	if errGetTokenizedCard != nil {
		c.JSON(http.StatusInternalServerError, responses.GetErrorResponse("Error connecting to database"+errGetTokenizedCard.Error()))
		return
	}

	c.JSON(http.StatusOK, responses.GetTokenizedCardResponse(merchant))
}

func (handler *TokenizedCardHandler) CreateTokenizedCard(c *gin.Context) {
	customerId := c.Param("id")
	var tokenizedCardParams = &requests.TokenizedCardRequest{}
	if errValidation := c.ShouldBind(&tokenizedCardParams); errValidation != nil {
		c.Status(http.StatusBadRequest)
		_ = c.Error(errValidation)
		return
	}

	tokenizedCard, errUuid := utils.TokenizeCard(*tokenizedCardParams, customerId)
	if errUuid != nil {
		c.Status(http.StatusBadRequest)
		_ = c.Error(errUuid)
		return
	}

	tokenizedCard, errorInsertTokenizedCard := handler.TokenizedCardDao.CreateTokenizedCard(*tokenizedCard)
	if errorInsertTokenizedCard == nil && tokenizedCard == nil {
		c.Status(http.StatusBadRequest)
		_ = c.Error(errtrace.New("Card already exists"))
		return
	}

	if errorInsertTokenizedCard != nil {
		c.Status(http.StatusInternalServerError)
		_ = c.Error(errorInsertTokenizedCard)
		return
	}

	c.JSON(http.StatusOK, responses.GetTokenizedCardResponse(tokenizedCard))
}

func (handler *TokenizedCardHandler) DeleteTokenizedCard(c *gin.Context) {
	errorDeleteTokenizedCard := handler.TokenizedCardDao.DeleteTokenizedCard(c.Param("token"))
	if errorDeleteTokenizedCard != nil {
		if errorDeleteTokenizedCard.Error() == "no rows affected" {
			c.Status(http.StatusNotFound)
		} else {
			c.Status(http.StatusInternalServerError)
		}
		_ = c.Error(errorDeleteTokenizedCard)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func NewTokenizedCardHandler(container *container.Container) TokenizedCardHandlerInterface {
	return &TokenizedCardHandler{
		TokenizedCardDao: dao.NewTokenizedCardDao(container),
	}
}
