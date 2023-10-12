package component

import (
	"studyGoApp/component/uploadprovider"
	"studyGoApp/pubsub"

	"github.com/jmoiron/sqlx"
)

type AppContext interface {
	GetMainDBConnection() *sqlx.DB
	SecretKey() string
	UploadProvider() uploadprovider.UploadProvider
	GetPubSub() pubsub.Pubsub
}

type appCtx struct {
	db         *sqlx.DB
	secretKey  string
	upProvider uploadprovider.UploadProvider
	pb         pubsub.Pubsub
}

func NewAppContext(
	db *sqlx.DB,
	secretKey string,
	upProvider uploadprovider.UploadProvider,
	pb pubsub.Pubsub,
) *appCtx {
	return &appCtx{db: db, secretKey: secretKey, upProvider: upProvider, pb: pb}
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

func (ctx *appCtx) GetPubSub() pubsub.Pubsub {
	return ctx.pb
}
