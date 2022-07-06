package svc

import (
	"go-zero-bookmall/service/book/api/internal/config"
	"go-zero-bookmall/service/book/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	BookModel model.BookModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		BookModel: model.NewBookModel(conn, c.CacheRedis),
	}
}
