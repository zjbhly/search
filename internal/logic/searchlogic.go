package logic

import (
	"context"
	"fmt"
	"strings"

	baseerror "search/common"
	"search/internal/svc"
	"search/internal/types"
	"search/model"

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

func (l *SearchLogic) Search(req *types.Request) (resp *types.Response, err error) {
	// 对address进行判断，为空则返回参数错误
	if len(strings.TrimSpace(req.Address)) == 0 {
		return nil, baseerror.NewDefaultError("参数错误")
	}
	address := &model.SearchPermission{
		Address: req.Address,
	}
	info, err := l.svcCtx.SearchPermissionModel.FindOne(l.ctx, address.Address)
	switch err {
	case nil:
		if info.Permission != 1 {
			return nil, baseerror.NewDefaultError("您还未获得白名单铸造资格，请关注最新活动消息。")
		} else {
			fmt.Println("恭喜您获得白名单铸造资格！")
		}
	case model.ErrNotFound:
		return nil, baseerror.NewDefaultError("用户签名不存在")
	default:
		return nil, err
	}
	return &types.Response{
		Permission: info.Permission,
	}, nil
}
