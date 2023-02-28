package middleware

import (
	"context"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-2/common"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-2/components"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-2/components/tokenprovider/jwt"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-2/internal/entities"
	"google.golang.org/grpc"
)

var ignoreMethod = []string{
	"/proto_v1.UserService/Hello",
}

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"wrong authen header",
		"ErrWrongAuthHeader",
	)
}

func Authenticate(appCtx components.AppContext) func(context.Context) (context.Context, error) {
	return func(ctx context.Context) (context.Context, error) {
		// ========== IGNORE METHOD NOT NEED AUTH =============
		method, _ := grpc.Method(ctx)
		appCtx.GetLogging().Infof("| %s", method)
		for _, imethod := range ignoreMethod {
			if method == imethod {
				return ctx, nil
			}
		}
		// ====================================================

		token, err := grpc_auth.AuthFromMD(ctx, "Bearer")
		if err != nil {
			panic(ErrWrongAuthHeader(err))
		}

		tokenProvider := jwt.NewTokenJWTProvider(appCtx.GetSecretKey())

		payload, err := tokenProvider.Validate(token)
		if err != nil {
			panic(err)
		}

		// store := authstorage.NewSQLStore(appCtx.GetMainDBConnection())
		// if _, err := store.FindToken(
		// 	ctx,
		// 	map[string]interface{}{
		// 		"token":   token,
		// 		"auth_id": payload.UserID,
		// 	}); err != nil {
		// 	panic(tokenprovider.ErrInvalidToken)
		// }

		// user, err := store.FindAuth(ctx, map[string]interface{}{"id": payload.UserID}, "Role")
		// if err != nil {
		// 	if err == common.RecordNotFound {
		// 		panic(authmodel.ErrUserBannedOrDeleted)
		// 	}
		// 	panic(authmodel.ErrCannotGetTokenLogin(err))
		// }
		user := entities.User{SQLModel: common.SQLModel{ID: payload.UserID}}

		// admin := false
		// for _, val := range user.GetRoles() {
		// 	if val == 5 || val == 6 {
		// 		admin = true
		// 	}
		// }
		// user.Mask(admin)

		newCtx := context.WithValue(ctx, common.CurrentUser, user)
		newCtx = context.WithValue(newCtx, common.TokenUser, token)
		return newCtx, nil
	}
}
