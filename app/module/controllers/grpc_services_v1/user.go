package grpc_services_v1

import (
	"context"
	"strings"

	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-2/app/common"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-2/app/module/entities"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-2/app/module/usecases"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-2/app/module/usecases/repository"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-2/app/module/usecases/transport"
	proto_v1 "github.com/nguyen-phi-khanh-monorevo/go-clean-architech-2/protoc-gen/proto/proto-v1"
)

func (s *grpcUserServiceV1) Hello(ctx context.Context, req *proto_v1.UserRequest) (*proto_v1.UserResponse, error) {
	// Do something here
	return &proto_v1.UserResponse{Message: "Hello"}, nil
}

func (s *grpcUserServiceV1) RegistUser(
	ctx context.Context,
	req *proto_v1.RegistUserRequest,
) (*proto_v1.RegistUserRequest, error) {
	var (
		db      = s.appCtx.GetMainDBConnection().Begin()
		repo    = repository.NewUserRepo(db, s.appCtx)
		transp  = transport.NewUserTransport(s.appCtx)
		useCase = usecases.NewRegisterUserUsecase(repo, transp, s.appCtx)
	)
	defer func() {
		if r := recover(); r != nil {
			db.Rollback()
		} else {
			db.Commit()
		}
	}()

	req.Email = strings.TrimSpace(req.Email)
	req.Password = strings.TrimSpace(req.Password)
	if req.Email == "" || req.Password == "" {
		panic(entities.ErrUsernameOrPasswordInvalid)
	}
	if !common.ValidateEmail(req.Email) {
		panic(entities.ErrUsernameOrPasswordInvalid)
	}

	user := entities.UserCreate{
		Email:     req.Email,
		Password:  req.Password,
		LastName:  req.LastName,
		FirstName: req.FirstName,
	}
	userCreated, err := useCase.Register(ctx, &user)
	if err != nil {
		panic(err)
	}
	return &proto_v1.RegistUserRequest{
		Email:     userCreated.Email,
		LastName:  userCreated.LastName,
		FirstName: userCreated.FirstName,
	}, nil
}
