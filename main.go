package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-1/common"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-1/components"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-1/components/logging"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-1/components/mailprovider/mail"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-1/components/uploadprovider"
	v1 "github.com/nguyen-phi-khanh-monorevo/go-clean-architech-1/internal/controllers/http/v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("____CLEAN ARCHITECH Khanh chế____")
	env := common.Init(".env.yml")

	connStr := fmt.Sprintf(env.DBConnectionStr, env.DBPassword)
	db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})

	// db = db.Debug()
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
	appCtx := components.NewAppContext(db, provider, env.SecretKeyJWT, mailProvider, &env, logger)

	route := gin.Default()
	route.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	err = v1.NewRouter(route, appCtx)
	if err != nil {
		panic(err)
	}
	route.Run(env.HttpPort)
}
