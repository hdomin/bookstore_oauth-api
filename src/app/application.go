package app

import (
	"github.com/gin-gonic/gin"
	accesstoken "github.com/hdomin/bookstore_oauth-api/src/domain/access_token"
	"github.com/hdomin/bookstore_oauth-api/src/http"
	"github.com/hdomin/bookstore_oauth-api/src/repository/db"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atHandler := http.NewHandler(accesstoken.NewService(db.NewRepository()))

	router.GET("/oauth/access_token/:access_token", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run(":8080")
}
