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
	Create(ctx context.Context, user *entities.UserCreate) (*entities.UserCreate, error)
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

func (u *registerUserUsecase) Register(ctx context.Context, data *entities.UserCreate) (resp *entities.UserCreate, err error) {
	if user, err := u.repo.FindUser(ctx, map[string]interface{}{"email": data.Email}); user != nil || err != nil {
		if user != nil {
			return nil, entities.ErrEmailExisted
		}
		if err != nil && err != common.RecordNotFound {
			return nil, err
		}
	}
	u.appCtx.GetLogging().Debug("Sample debug log")

	data.Password, err = hasher.HashPassword(data.Password)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	if resp, err = u.repo.Create(ctx, data); err != nil {
		return nil, err
	}
	resp.Password = ""

	// Call api orhter service to process
	userProcess := entities.User{}
	responseFromMS, err := u.transport.CallAPIInMicroserviceBlaBla(ctx, &userProcess)
	if err != nil {
		return nil, err
	}

	// logging value
	u.appCtx.GetLogging().Infof("user data: %v", responseFromMS)

	// Gen new uid for this account
	resp.SQLModel.GenUID(common.DbTypeAccount)

	return
}
