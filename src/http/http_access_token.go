package http

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/prateek-sha/go_auth_api/src/domain/access_token"
	errors "github.com/prateek-sha/go_auth_api/src/utils/error"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(*gin.Context)
	UpdateExpirationTime(*gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (h *accessTokenHandler) GetById(c *gin.Context) {
	access_token_id := strings.TrimSpace(c.Param("access_token_id"))

	accessToken, err := h.service.GetById(access_token_id)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

func (h *accessTokenHandler) Create(c *gin.Context) {
	var at access_token.AccessToken
	if err := c.ShouldBindJSON(&at); err != nil {
		restErr := errors.NewBadRequestError("invalid json")
		c.JSON(restErr.Status, restErr)
		return
	}
	if err := h.service.Create(at); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, at)
}

func (h *accessTokenHandler) UpdateExpirationTime(c *gin.Context) {
	var at access_token.AccessToken
	if err := c.ShouldBindJSON(&at); err != nil {
		restErr := errors.NewBadRequestError("invalid json")
		c.JSON(restErr.Status, restErr)
		return
	}
	if err := h.service.UpdateExpirationTime(at); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, at)
}
