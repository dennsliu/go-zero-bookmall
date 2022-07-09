package logic

import (
	"context"
	"fmt"

	"go-zero-bookmall/service/book/api/internal/svc"
	"go-zero-bookmall/service/book/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req *types.SearchReq) (resp *types.SearchReply, err error) {
	// todo: add your logic here and delete this line
	result, err := l.svcCtx.BookModel.FindByKeyword(req.Keyword, req.Page, req.PageSize, req.OrderBy)
	if err != nil {
		return nil, err
	}
	size := len(*result)
	fmt.Printf("------------result------size:%d", size)
	var rsp types.SearchReply
	rsp.Books = make([]types.Book, size, size)
	for i := 0; i < size; i++ {
		rsp.Books[i].Id = (*result)[i].Id
		rsp.Books[i].Plu = (*result)[i].Plu
		rsp.Books[i].Name = (*result)[i].Name
		rsp.Books[i].Image = (*result)[i].Image
		rsp.Books[i].Description = (*result)[i].Description
		rsp.Books[i].InStocked = (*result)[i].InStocked
	}
	return &rsp, nil
}
