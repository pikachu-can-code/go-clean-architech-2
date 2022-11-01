package http

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-1/components"
	v1 "github.com/nguyen-phi-khanh-monorevo/go-clean-architech-1/internal/controllers/http/v1"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-1/middleware"
)

func NewRouter(r *gin.Engine, appCtx components.AppContext) {
	// config CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 30 * time.Minute,
	}))

	// use recovery middleware, this middleware will catch panic and return 500
	r.Use(middleware.Recover(appCtx))

	// group API to user group: /user/v1/*
	user := r.Group("/user")

	user.GET("/health_check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "I am find! ok",
		})
	})

	v1.NewRouterV1(user, appCtx)
}
