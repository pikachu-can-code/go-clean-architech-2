package proto_v1_test

import (
	"context"
	"testing"

	proto_v1 "github.com/nguyen-phi-khanh-monorevo/go-clean-architech-2/protoc-gen/proto/proto-v1"
)

func TestHello(t *testing.T) {
	ctx := context.Background()

	client, _, closer := UserServer(ctx, "http://0.0.0.0:8080/proto_v1.UserService/Hello")
	defer func() {
		closer()
	}()

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