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
	"google.golang.org/grpc"
)

func RegisterUser(appCtx components.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data entities.UserCreate

		// Bind request data to struct
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		conn, err := grpc.Dial(appCtx.GetEnv().UserServiceEndpoint, grpc.WithInsecure())
		if err != nil {
			panic(err)
		}

		// declare dependencies and usecases
		var (
			repo    = repository.NewUserRepo(appCtx.GetMainDBConnection(), appCtx)
			transp  = transport.NewUserTransport(appCtx, conn)
			usecase = usecases.NewRegisterUserUsecase(repo, transp, appCtx)
			resp    *entities.UserCreate
		)

		// call usecase
		if resp, err = usecase.Register(c, &data); err != nil {
			panic(err)
		}

		// response
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(resp))
	}
}
