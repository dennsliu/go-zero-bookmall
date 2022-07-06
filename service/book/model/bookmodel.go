package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BookModel = (*customBookModel)(nil)

type (
	// BookModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBookModel.
	BookModel interface {
		bookModel
		//FindByKeyword(keyword string, page string, pageSize string, orderBy string) error
	}

	customBookModel struct {
		*defaultBookModel
	}
)

// NewBookModel returns a model for the database table.
func NewBookModel(conn sqlx.SqlConn, c cache.CacheConf) BookModel {
	return &customBookModel{
		defaultBookModel: newBookModel(conn, c),
	}
}

// func FindByKeyword(keyword string, page, pageSize int64, orderBy string) error {

// 	// if orderBy == "" {
// 	// 	orderBy = "id"
// 	// }

// 	// if page < 1 {
// 	// 	page = 1
// 	// }
// 	// if pageSize < 1 {
// 	// 	pageSize = 20
// 	// }
// 	// offset := (page - 1) * pageSize
// 	// var resp []Book
// 	// query := fmt.Sprintf("select %s from %s where `id` > %d limit %d", bookRows, m.table, offset, pageSize)
// 	// err := m.QueryRowNoCache(&resp, query)
// 	// switch err {
// 	// case nil:
// 	// 	return &resp, nil
// 	// default:
// 	// 	return nil, err
// 	// }
// 	return err
// }
