package logic

import (
	"context"
	"errors"
	"fmt"
	"go-zero-bookmall/service/book/api/internal/svc"
	"go-zero-bookmall/service/book/api/internal/types"
	"go-zero-bookmall/service/book/model"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req *types.AddReq) (resp *types.AddReply, err error) {
	// todo: add your logic here and delete this line
	fmt.Print("testing add")

	if len(strings.TrimSpace(req.Plu)) == 0 || len(strings.TrimSpace(req.Sku)) == 0 || len(strings.TrimSpace(req.Name)) == 0 || len(strings.TrimSpace(req.InStocked)) == 0 {
		return nil, errors.New("Plu, Sku,Name and InStocked are required")
	}

	bookModel := new(model.Book)
	bookModel.Plu = req.Plu
	bookModel.Sku = req.Sku
	bookModel.Name = req.Name
	bookModel.Description = req.Description
	bookModel.InStocked = req.InStocked
	_, err = l.svcCtx.BookModel.Insert(l.ctx, bookModel)
	if err != nil {
		return nil, err
	}

	bookInfo, err := l.svcCtx.BookModel.FindOneBySku(l.ctx, req.Sku)
	if err != nil {
		return nil, err
	}

	return &types.AddReply{
		Id:          bookInfo.Id,
		Plu:         bookInfo.Plu,
		Sku:         bookInfo.Sku,
		Name:        bookInfo.Name,
		Image:       bookInfo.Image,
		Description: bookInfo.Description,
		InStocked:   bookInfo.InStocked,
	}, nil
}
