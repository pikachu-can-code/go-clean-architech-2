package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-1/components"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-1/middleware"
)

func NewRouter(r *gin.Engine, appCtx components.AppContext) (err error) {
	r.GET("/health_check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "I am find! ok",
		})
	})
	r.Use(middleware.Recover(appCtx))

	v1 := r.Group("/v1")

	user := v1.Group("/user")
	{
		user.POST("/register", RegisterUser(appCtx))
	}

	return
}
