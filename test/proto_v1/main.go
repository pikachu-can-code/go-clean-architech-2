package proto_v1_test

import (
	"context"
	"log"
	"net"

	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-2/common"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-2/components"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-2/components/logging"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-2/components/mailprovider/mail"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-2/components/uploadprovider"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-2/internal/controllers/grpc_services_v1"
	proto_v1 "github.com/nguyen-phi-khanh-monorevo/go-clean-architech-2/protoc-gen/proto/proto-v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func UserServer(ctx context.Context, target string) (client proto_v1.UserServiceClient, appCtx components.AppContext, closer func()) {
	appCtx = initAppCtx()
	var (
		buffer      = 101024 * 1024
		lis         = bufconn.Listen(buffer)
		baseServer  = grpc.NewServer()
		userService = grpc_services_v1.NewGRPCUserServiceV1(appCtx)
	)
	proto_v1.RegisterUserServiceServer(baseServer, userService)
	go func() {
		if err := baseServer.Serve(lis); err != nil {
			appCtx.GetLogging().DPanicf("Error start server: %v", err)
		}
	}()
	conn, err := grpc.DialContext(ctx, target, grpc.WithContextDialer(
		func(context.Context, string) (net.Conn, error) {
			return lis.Dial()
		},
	), grpc.WithInsecure())
	if err != nil {
		appCtx.GetLogging().DPanicf("Error connect server: %v", err)
	}
	closer = func() {
		err := lis.Close()
		if err != nil {
			appCtx.GetLogging().DPanicf("Error close listener: %v", err)
		}
	}
	client = proto_v1.NewUserServiceClient(conn)
	return
}

func initAppCtx() components.AppContext {
	env := common.Init("../../.env-ut")
	connStr := env.DBConnectionStr
	db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	sql, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer sql.Close()

	provider := uploadprovider.NewS3Provider(
		env.S3BucketName,
		env.S3Region,
		env.S3APIKey,
		env.S3Secret,
		env.S3Domain,
	)

	mailProvider := mail.NewMailProvider(env.BaseEmailPassword)

	logger := logging.NewAPILogger()

	return components.NewAppContext(db, provider, env.SecretKeyJWT, mailProvider, &env, logger)
}
