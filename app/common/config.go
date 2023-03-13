package common

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	DBConnectionStr     string
	DBPassword          string
	ServerPort          string
	HttpPort            string
	RunGateway          bool
	S3BucketName        string
	S3Region            string
	S3APIKey            string
	S3Secret            string
	S3Domain            string
	SecretKeyJWT        string
	FirebaseService     string
	BaseEmailPassword   string
	DefaultEndpoint     string
	IsDeployed          bool
	UserServiceEndpoint string
}

func Init(dirFile string) Env {
	err := godotenv.Load(dirFile)
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	var env Env
	env.DBConnectionStr = os.Getenv("DB_CONNECTION")
	env.DBPassword = os.Getenv("DB_PASSWORD")
	env.ServerPort = os.Getenv("SERVER_PORT")
	env.HttpPort = os.Getenv("HTTP_PORT")
	env.RunGateway = os.Getenv("RUN_GATEWAY") == "true"
	env.S3BucketName = os.Getenv("S3_BUCKET_NAME")
	env.S3Region = os.Getenv("S3_REGION")
	env.S3APIKey = os.Getenv("S3_API_KEY")
	env.S3Secret = os.Getenv("S3_SECRET")
	env.S3Domain = os.Getenv("S3_DOMAIN")
	env.SecretKeyJWT = os.Getenv("SECRET_KEY_JWT")
	env.FirebaseService = os.Getenv("FIREBASE_SERVICE")
	env.BaseEmailPassword = os.Getenv("EMAIL")
	env.DefaultEndpoint = os.Getenv("DEFAULT_ENDPOINT")
	env.UserServiceEndpoint = os.Getenv("USER_SERVICE_ENDPOINT")
	return env
}
