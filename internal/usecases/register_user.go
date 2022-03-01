package usecases

import (
	"context"

	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-1/common"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-1/components"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-1/components/encoder/hasher"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-1/internal/entities"
)

type RegisterUserRepo interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*entities.User, error)
	Create(ctx context.Context, acc *entities.User) error
}

type RegisterUserTransport interface {
	CallAPIInMicroserviceBlaBla(ctx context.Context, user *entities.User) (*entities.User, error)
}

type registerUserUsecase struct {
	repo      RegisterUserRepo
	transport RegisterUserTransport
	appCtx    components.AppContext
}

func NewRegisterUserUsecase(repo RegisterUserRepo, transport RegisterUserTransport, appCtx components.AppContext) *registerUserUsecase {
	return &registerUserUsecase{repo: repo, transport: transport, appCtx: appCtx}
}

func (u *registerUserUsecase) Register(ctx context.Context, data *entities.User) (err error) {
	if user, err := u.repo.FindUser(ctx, map[string]interface{}{"email": data.Email}); user != nil || err != nil {
		if user != nil {
			return entities.ErrEmailExisted
		}
		if err != nil && err != common.RecordNotFound {
			u.appCtx.GetLogging().Debug("alo")
			return err
		}
	}

	data.Password, err = hasher.HashPassword(data.Password)
	if err != nil {
		return common.ErrInternal(err)
	}

	if err := u.repo.Create(ctx, data); err != nil {
		return err
	}

	// Call api orhter service to process
	data, err = u.transport.CallAPIInMicroserviceBlaBla(ctx, data)
	if err != nil {
		return err
	}

	// logging value
	u.appCtx.GetLogging().Infof("user data: %v", data)

	return
}
