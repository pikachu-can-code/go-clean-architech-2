package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-1/common"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-1/components"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-1/internal/entities"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-1/internal/usecases"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-1/internal/usecases/repository"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-1/internal/usecases/transport"
)

func RegisterUser(appCtx components.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data entities.User

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var (
			repo    = repository.NewUserRepo(appCtx.GetMainDBConnection(), appCtx)
			transp  = transport.NewUserTransport(appCtx)
			usecase = usecases.NewRegisterUserUsecase(repo, transp, appCtx)
		)
		if err := usecase.Register(c, &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
