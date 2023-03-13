package components

import (
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-2/app/common"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-2/app/components/logging"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-2/app/components/mailprovider/mail"
	"github.com/nguyen-phi-khanh-monorevo/go-clean-architech-2/app/components/uploadprovider"

	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	GetUploadProvider() uploadprovider.UploadProvider
	GetSecretKey() string
	GetMailInfo() *mail.MailProvider
	GetEnv() *common.Env
	GetLogging() logging.Logger
}

type appCtx struct {
	db             *gorm.DB
	uploadProvider uploadprovider.UploadProvider
	secretKey      string
	mailProvider   *mail.MailProvider
	env            *common.Env
	logging        logging.Logger
}

func NewAppContext(
	db *gorm.DB,
	uploadProvider uploadprovider.UploadProvider,
	secretKey string,
	mailProvider *mail.MailProvider,
	env *common.Env,
	logging logging.Logger,
) *appCtx {
	return &appCtx{
		db:             db,
		uploadProvider: uploadProvider,
		secretKey:      secretKey,
		mailProvider:   mailProvider,
		env:            env,
		logging:        logging,
	}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) GetUploadProvider() uploadprovider.UploadProvider {
	return ctx.uploadProvider
}

func (ctx *appCtx) GetSecretKey() string {
	return ctx.secretKey
}

func (ctx *appCtx) GetMailInfo() *mail.MailProvider {
	return ctx.mailProvider
}

func (ctx *appCtx) GetEnv() *common.Env {
	return ctx.env
}

func (ctx *appCtx) GetLogging() logging.Logger {
	return ctx.logging
}
