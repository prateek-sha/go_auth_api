package application

import (
	"github.com/gin-gonic/gin"
	"github.com/prateek-sha/go_auth_api/src/domain/access_token"
	"github.com/prateek-sha/go_auth_api/src/http"
	"github.com/prateek-sha/go_auth_api/src/repository/db"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atService := access_token.NewService(db.New())
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/acces_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/acces_token", atHandler.Create)
	router.PUT("/oauth/acces_token", atHandler.UpdateExpirationTime)
	router.Run(":8080")
}
