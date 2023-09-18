package component

import (
	"studyGoApp/component/uploadprovider"

	"github.com/jmoiron/sqlx"
)

type AppContext interface {
	GetMainDBConnection() *sqlx.DB
	SecretKey() string
	UploadProvider() uploadprovider.UploadProvider
}

type appCtx struct {
	db         *sqlx.DB
	secretKey  string
	upProvider uploadprovider.UploadProvider
}

func NewAppContext(db *sqlx.DB, secretKey string, upProvider uploadprovider.UploadProvider) *appCtx {
	return &appCtx{db: db, secretKey: secretKey, upProvider: upProvider}
}

func (ctx *appCtx) GetMainDBConnection() *sqlx.DB {
	return ctx.db
}

func (ctx *appCtx) UploadProvider() uploadprovider.UploadProvider {
	return ctx.upProvider
}

func (ctx *appCtx) SecretKey() string {
	return ctx.secretKey
}
