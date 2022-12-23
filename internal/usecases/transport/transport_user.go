package transport

import (
	"context"

	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-1/components"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-1/internal/entities"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-1/protoc-gen/proto/proto-v1"
	"google.golang.org/grpc"
)

type userTransport struct {
	appCtx components.AppContext
	client proto_v1.UserServiceClient
}

func NewUserTransport(appCtx components.AppContext, conn *grpc.ClientConn) *userTransport {
	return &userTransport{appCtx: appCtx, client: proto_v1.NewUserServiceClient(conn)}
}

func (t *userTransport) CallAPIInMicroserviceBlaBla(
	ctx context.Context,
	user *entities.User,
) (resp *proto_v1.UserResponse, err error) {
	// Call some grpc API in microservice bla bla
	resp, err = t.client.Hello(ctx, &proto_v1.UserRequest{Id: user.FakeId.String()})
	if err != nil {
		return
	}
	return
}
