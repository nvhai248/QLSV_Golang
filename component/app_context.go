package component

import (
	"studyGoApp/component/uploadprovider"

	"github.com/jmoiron/sqlx"
)

type AppContext interface {
	GetMainDBConnection() *sqlx.DB
	UploadProvider() uploadprovider.UploadProvider
}

type appCtx struct {
	db         *sqlx.DB
	upProvider uploadprovider.UploadProvider
}

func NewAppContext(db *sqlx.DB, upProvider uploadprovider.UploadProvider) *appCtx {
	return &appCtx{db: db, upProvider: upProvider}
}

func (ctx *appCtx) GetMainDBConnection() *sqlx.DB {
	return ctx.db
}

func (ctx *appCtx) UploadProvider() uploadprovider.UploadProvider {
	return ctx.upProvider
}
