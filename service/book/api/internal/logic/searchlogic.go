package logic

import (
	"context"

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

	//result, err := l.svcCtx.BookModel.FindByKeyword(l.ctx, req.Keyword, req.Page, req.PageSize, req.OrderBy)
	// if err != nil {
	// 	return nil, err
	// }
	return

}
