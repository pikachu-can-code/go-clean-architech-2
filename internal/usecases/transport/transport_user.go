package transport

import (
	"context"

	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-1/components"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-1/internal/entities"
)

type userTransport struct {
	appCtx components.AppContext
}

func NewUserTransport(appCtx components.AppContext) *userTransport {
	return &userTransport{appCtx: appCtx}
}

func (t *userTransport) CallAPIInMicroserviceBlaBla(ctx context.Context, user *entities.User) (*entities.User, error) {
	// Call some api
	return user, nil
}
