package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-1/components"
)

func NewRouterV1(r *gin.Engine, appCtx components.AppContext) {
	v1 := r.Group("/v1")

	user := v1.Group("/user")
	{
		user.POST("/register", RegisterUser(appCtx))
	}
}
