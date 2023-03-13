package transport

import (
	"context"

	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-2/app/components"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-2/app/module/entities"
	proto_v1 "github.com/nguyen-phi-khanh-monorevo/go-clean-architech-2/protoc-gen/proto/proto-v1"
	"google.golang.org/grpc"
)

type userTransport struct {
	appCtx components.AppContext
}

func NewUserTransport(appCtx components.AppContext) *userTransport {
	return &userTransport{appCtx: appCtx}
}

func (t *userTransport) CallAPIInMicroserviceBlaBla(
	ctx context.Context,
	user *entities.User,
) (resp *proto_v1.UserResponse, err error) {
	// Call some grpc API in microservice bla bla
	conn, err := grpc.Dial(t.appCtx.GetEnv().UserServiceEndpoint, grpc.WithInsecure())
	client := proto_v1.NewUserServiceClient(conn)
	resp, err = client.Hello(ctx, &proto_v1.UserRequest{Id: user.FakeId.String()})
	if err != nil {
		return
	}
	return
}
