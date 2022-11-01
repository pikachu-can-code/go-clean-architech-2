package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-1/components"
)

func NewRouterV1(r *gin.RouterGroup, appCtx components.AppContext) {
	//define API user version 1
	v1 := r.Group("/v1")
	{
		v1.POST("/register", RegisterUser(appCtx))
	}
}
