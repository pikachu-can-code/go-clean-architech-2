package proto_v1_test

import (
	"context"
	"testing"

	proto_v1 "github.com/nguyen-phi-khanh-monorevo/go-clean-architech-2/protoc-gen/proto/proto-v1"
	"google.golang.org/grpc"
)

func HelloTest(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "http://0.0.0.0:8080/proto_v1.UserService/Hello",grpc.WithInsecure())
	if err != nil {
		t.Errorf("Error connect service: %s", err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			t.Errorf("Error close connection: %s", err)
		}
	}()

	client := proto_v1.NewUserServiceClient(conn)
	resp, err := client.Hello(ctx, &proto_v1.UserRequest{
		Id: "test",
	})
	if err != nil {
		t.Errorf("Error call service: %s", err)
	}
	if resp.Message != "Hello" {
		t.Errorf("Error call service: %s", err)
	}
}