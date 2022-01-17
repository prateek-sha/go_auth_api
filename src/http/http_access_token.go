package http

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/prateek-sha/go_auth_api/src/domain/access_token"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
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

	accessToken, err := h.service.GetUserById(access_token_id)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
	c.JSON(http.StatusNotImplemented, "implemt me")
}
