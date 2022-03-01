package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-1/common"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-1/components"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-1/components/tokenprovider/jwt"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-1/internal/usecases/repository"
)

var ignoreMethod = []string{
	"/helth_check",
}

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewFullErrorResponse(
		http.StatusUnauthorized,
		err,
		"wrong authen header",
		err.Error(),
		"ErrWrongAuthHeader",
	)
}

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	//"Authorization" : "Bearer {token}"

	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}

	return parts[1], nil
}

func RequireAuth(appCtx components.AppContext) gin.HandlerFunc {
	tokenProvider := jwt.NewTokenJWTProvider(appCtx.GetSecretKey())

	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))
		if err != nil {
			panic(err)
		}

		// ========== IGNORE METHOD NOT NEED AUTH =============
		// method := c.Request.URL.Path
		// for _, imethod := range ignoreMethod {
		// 	if method == imethod {
		// 		c.Next()
		// 	}
		// }
		// ====================================================

		db := appCtx.GetMainDBConnection()
		repo := repository.NewUserRepo(db, appCtx)

		payload, err := tokenProvider.Validate(token)
		if err != nil {
			panic(err)
		}

		user, err := repo.FindUser(c.Request.Context(), map[string]interface{}{"id": payload.UserID})
		if err != nil {
			panic(err)
		}

		if user.Status == 0 {
			panic(common.ErrNoPermission(errors.New("user has been deleted or banned")))
		}

		c.Set(common.CurrentUser, user)
		c.Set(common.TokenUser, token)
		c.Next()
	}
}
