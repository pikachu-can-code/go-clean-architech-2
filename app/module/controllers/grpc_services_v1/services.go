package grpc_services_v1

import (
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-2/app/components"
	proto_v1 "github.com/nguyen-phi-khanh-monorevo/go-clean-architech-2/protoc-gen/proto/proto-v1"
)

type grpcUserServiceV1 struct {
	proto_v1.UserServiceServer
	appCtx components.AppContext
}

type grpcProductServiceV1 struct {
	proto_v1.ProductServiceServer
	appCtx components.AppContext
}

func NewGRPCUserServiceV1(appCtx components.AppContext) *grpcUserServiceV1 {
	return &grpcUserServiceV1{
		appCtx: appCtx,
	}
}

func NewGRPCProductServiceV1(appCtx components.AppContext) *grpcProductServiceV1 {
	return &grpcProductServiceV1{
		appCtx: appCtx,
	}
}
