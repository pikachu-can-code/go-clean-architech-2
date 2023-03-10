package main

import (
	"fmt"
	"log"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-2/common"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-2/components"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-2/components/logging"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-2/components/mailprovider/mail"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-2/components/uploadprovider"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-2/internal/controllers/grpc_services_v1"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-2/middleware"
	proto_v1 "github.com/nguyen-phi-khanh-monorevo/go-clean-architech-2/protoc-gen/proto/proto-v1"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("____CLEAN ARCHITECH Khanh chế____")
	env := common.Init(".env")

	// init sql connection, this connection will keep alive until the app is closed
	connStr := env.DBConnectionStr
	db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})
	// use db Debug to see sql query
	// db = db.Debug()
	if err != nil {
		log.Fatalln(err)
	}
	sql, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer sql.Close()

	// init S3 provider
	provider := uploadprovider.NewS3Provider(
		env.S3BucketName,
		env.S3Region,
		env.S3APIKey,
		env.S3Secret,
		env.S3Domain,
	)
	// init mail provider
	mailProvider := mail.NewMailProvider(env.BaseEmailPassword)

	// init logger
	logger := logging.NewAPILogger()

	// init App Context, this App Context will be passed to all components
	appCtx := components.NewAppContext(db, provider, env.SecretKeyJWT, mailProvider, &env, logger)

	if err := runGrpcServices(appCtx, env.ServerPort, env.HttpPort, env.RunGateway); err != nil {
		log.Fatalf("Failed to start gRPC Server at port %v with error: %v", env.ServerPort, err)
	}
}

func runGrpcServices(
	appCtx components.AppContext,
	port string,
	httpPort string,
	runGateway bool,
) error {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}
	s := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_recovery.UnaryServerInterceptor(middleware.RecoverOptions...),
			grpc_auth.UnaryServerInterceptor(middleware.Authenticate(appCtx)),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_recovery.StreamServerInterceptor(middleware.RecoverOptions...),
			grpc_auth.StreamServerInterceptor(middleware.Authenticate(appCtx)),
		),
	)
	//====== Init all services here ======//
	proto_v1.RegisterUserServiceServer(s, grpc_services_v1.NewGRPCUserServiceV1(appCtx))
	proto_v1.RegisterProductServiceServer(s, grpc_services_v1.NewGRPCProductServiceV1(appCtx))
	//====================================//
	go func() {
		if err := s.Serve(listen); err != nil {
			log.Fatalf("Failed to start gRPC Server: %v", err)
		}
	}()
	log.Printf("| Serving gRPC on http://0.0.0.0%s...\n", port)

	if runGateway {
		return RunGateway(port, httpPort)
	}

	return nil
}
