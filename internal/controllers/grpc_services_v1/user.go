package grpc_services_v1

import (
	"context"

	proto_v1 "github.com/nguyen-phi-khanh-monorevo/go-clean-architech-2/protoc-gen/proto/proto-v1"
)

func (s *grpcUserServiceV1) Hello(ctx context.Context, req *proto_v1.UserRequest) (*proto_v1.UserResponse, error) {
	// Do something here
	return &proto_v1.UserResponse{Message: "Hello"}, nil
}