// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"fmt"
	"im-server/common/models"
	"im-server/common/models/list_query"
	"im-server/im_user/user_models"
	"strconv"

	"im-server/im_user/user_api/internal/svc"
	"im-server/im_user/user_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserSearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserSearchLogic {
	return &UserSearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserSearchLogic) UserSearch(req *types.UserSearchRequest) (resp *types.UserListResponse, err error) {
	//优先根据用户昵称搜索，昵称没有根据uid搜索
	keyword := req.Nickname
	if len(keyword) == 0 {
		keyword = strconv.Itoa(int(req.UserID))
	}

	list, count, err := list_query.ListQuery(l.svcCtx.DB, user_models.UserModel{}, list_query.Option{
		PageInfo: models.PageInfo{
			Key:   keyword,
			Page:  1,
			Limit: 10,
		},
		Where:   l.svcCtx.DB,
		Likes:   []string{"ID", "Nickname"},
		Preload: []string{"UserConfModel"},
	})
	if err != nil {
		fmt.Println("搜索用户失败")
		return nil, err
	}
	if len(list) == 0 {
		return nil, fmt.Errorf("搜索结果为空")
	}
	result := make([]types.UserInfoData, 0)
	for _, user := range list {
		result = append(result, types.UserInfoData{
			UserID:   user.ID,
			Nickname: user.Nickname,
			Avatar:   user.Avatar,
			Abstract: user.Abstract,
		})
	}
	return &types.UserListResponse{
		List:  result,
		Count: int(count),
	}, nil
}
