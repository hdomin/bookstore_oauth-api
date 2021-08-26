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
	atService := accesstoken.NewService(db.NewRepository())
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token", atHandler.GetById)

	router.Run(":8080")
}
