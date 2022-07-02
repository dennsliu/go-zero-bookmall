package logic

import (
	"context"
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"go-zero-bookmall/service/user/api/internal/svc"
	"go-zero-bookmall/service/user/api/internal/types"
	"go-zero-bookmall/service/user/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/zeromicro/go-zero/core/logx"
)

type SignupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func init() {
	// 初始化随机数的资源库，如果不执行这行，不管运行多少次都返回同样的值
	rand.Seed(time.Now().UnixNano())
}
func NewSignupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignupLogic {
	return &SignupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func (l *SignupLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
func (l *SignupLogic) Signup(req *types.SignupReq) (resp *types.SignupReply, err error) {
	// todo: add your logic here and delete this line
	if len(strings.TrimSpace(req.Username)) == 0 || len(strings.TrimSpace(req.Password)) == 0 || len(strings.TrimSpace(req.Email)) == 0 {
		return nil, errors.New("Username, Email and Password are required")
	}
	gender := "male"
	if len(strings.TrimSpace(req.Gender)) == 0 {
		gender = req.Gender
	}
	userModel := new(model.User)
	userModel.Number = strconv.Itoa(rand.Int()) + strconv.Itoa(rand.Intn(99999))
	userModel.Username = req.Username
	userModel.Email = req.Email
	userModel.Password = req.Password
	userModel.Gender = gender
	_, err = l.svcCtx.UserModel.Insert(l.ctx, userModel)
	if err != nil {
		return nil, err
	}

	userInfo, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.Username)
	if err != nil {
		return nil, err
	}
	return &types.SignupReply{
		Id:       userInfo.Id,
		Username: userInfo.Username,
		Email:    userInfo.Email,
		Gender:   userInfo.Gender,
	}, nil
}
