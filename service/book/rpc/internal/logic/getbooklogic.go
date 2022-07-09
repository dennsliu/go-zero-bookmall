package logic

import (
	"context"

	"go-zero-bookmall/service/book/rpc/internal/svc"
	"go-zero-bookmall/service/book/rpc/types/book"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBookLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBookLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBookLogic {
	return &GetBookLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetBookLogic) GetBook(in *book.IdRequest) (*book.BookResponse, error) {
	// todo: add your logic here and delete this line

	return &book.BookResponse{}, nil
}
