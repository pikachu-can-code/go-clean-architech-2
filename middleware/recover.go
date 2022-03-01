package middleware

import (
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-1/common"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-1/components"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-1/components/logging"
)

func Recover(appCtx components.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger := logging.NewAPILogger()
				c.Header("Content-Type", "application/json")

				if appErr, ok := err.(*common.AppError); ok {
					logger.Errorf("[error] %v\n", appErr.Error())
					c.AbortWithStatusJSON(appErr.StatusCode, appErr)
					panic(err)
				}

				logger.Errorf("[error unknow] %v\n", err)
				debug.PrintStack()
				appErr := common.ErrInternal(err.(error))
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
				panic(err)
			}
		}()

		c.Next()
	}
}
