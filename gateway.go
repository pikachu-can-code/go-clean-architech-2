package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	proto_v1 "github.com/nguyen-phi-khanh-monorevo/go-clean-architech-2/protoc-gen/proto/proto-v1"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

// RunGateway runs the gRPC-Gateway on a given port.
func RunGateway(port, httpPort string) error {
	conn, err := grpc.DialContext(
		context.Background(),
		fmt.Sprintf("0.0.0.0%v", port),
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		return err
	}

	gwmux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames: true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: false,
			},
		}),
	)

	//==== REGISTER SERVICE HERE ====//
	err = proto_v1.RegisterProductServiceHandlerClient(
		context.Background(),
		gwmux,
		proto_v1.NewProductServiceClient(conn),
	)
	if err != nil {
		return err
	}
	err = proto_v1.RegisterUserServiceHandlerClient(
		context.Background(),
		gwmux,
		proto_v1.NewUserServiceClient(conn),
	)
	if err != nil {
		return err
	}
	//=====================//

	gwServer := &http.Server{
		Addr:    httpPort,
		Handler: gwmux,
	}
	log.Printf("| Serving gRPC-Gateway on http://0.0.0.0%s...\n", httpPort)
	return gwServer.ListenAndServe()
}
