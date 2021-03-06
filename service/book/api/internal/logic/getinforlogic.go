package logic

import (
	"context"

	"go-zero-bookmall/service/book/api/internal/svc"
	"go-zero-bookmall/service/book/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInfoLogic {
	return &GetInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetInfoLogic) GetInfo(req *types.GetInfoReq) (resp *types.GetInfoReply, err error) {
	// todo: add your logic here and delete this line

	return
}
