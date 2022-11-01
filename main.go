package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-1/common"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-1/components"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-1/components/logging"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-1/components/mailprovider/mail"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-1/components/uploadprovider"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-1/internal/controllers/http"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("____CLEAN ARCHITECH Khanh chế____")
	env := common.Init(".env.yml")

	// init sql connection, this connection will keep alive until the app is closed
	connStr := fmt.Sprintf(env.DBConnectionStr, env.DBPassword)
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

	// set release mode for gin
	if env.IsDeployed {
		gin.SetMode(gin.ReleaseMode)
	}
	route := gin.Default()

	http.NewRouter(route, appCtx)
	err = route.Run(env.HttpPort)
	if err != nil {
		log.Fatalf("Cannot start server at port %v with error: %v", env.HttpPort, err)
	}
}
