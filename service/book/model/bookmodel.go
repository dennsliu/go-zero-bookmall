package model

import (
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BookModel = (*customBookModel)(nil)

type (
	// BookModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBookModel.
	BookModel interface {
		bookModel
		FindByKeyword(keyword string, page int64, pageSize int64, orderBy string) (*[]Book, error)
		//Query(pos int, limit int) (*[]Book, error)
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

func (m *defaultBookModel) FindByKeyword(keyword string, page int64, pageSize int64, orderBy string) (*[]Book, error) {

	if orderBy == "" {
		orderBy = "id"
	}

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 20
	}
	var resp []Book
	query := fmt.Sprintf("select %s from %s order by %s limit %d", bookRows, m.table, orderBy, pageSize)
	err := m.QueryRowsNoCache(&resp, query)
	fmt.Println(err)
	fmt.Println(resp)
	switch err {
	case nil:
		return &resp, nil
	default:
		return nil, err
	}
}

// func (m *defaultBookModel) Query(pos int, limit int) (*[]Book, error) {
// 	var result []Book
// 	query := fmt.Sprintf("select %s from %s where (`id` > %d) limit %d",
// 		bookRows, m.table, pos, limit)
// 	err := m.QueryRowsNoCache(&result, query)
// 	switch err {
// 	case nil:
// 		return &result, nil
// 	case sqlc.ErrNotFound:
// 		return nil, ErrNotFound
// 	default:
// 		return nil, err
// 	}
// }
