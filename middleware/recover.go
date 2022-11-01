package middleware

import (
	"errors"
	"reflect"
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
				}

				logger.Errorf("[error unknow] %v\n", err)
				var appErr  *common.AppError
				if reflect.TypeOf(err) == reflect.TypeOf("string") {
					appErr = common.ErrInternal(errors.New(err.(string)))
				} else {
					appErr = common.ErrInternal(err.(error))
				}
				debug.PrintStack()
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
			}
		}()

		c.Next()
	}
}
