package user

import (
	"context"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"

	"gf-user/internal/consts"
	"gf-user/internal/dao"
	"gf-user/internal/model/code"
	"gf-user/internal/model/do"
	"gf-user/internal/service"
)

func init() {
	service.RegisterUser(&sUser{})
}

type sUser struct {
}

func (u sUser) GetUserFromToken(ctx context.Context, spaceId int64) (usr *do.User, err error) {
	tokenInfo := service.Token().GetTokenInfoFromCtx(ctx)
	usr, err = u.GetUserByAccountId(ctx, tokenInfo.AccountId, spaceId)
	return
}

func (u sUser) GetUserByAccountId(ctx context.Context, accountId string, spaceId int64) (usr *do.User, err error) {
	v, err := dao.User.Ctx(ctx).
		Where(dao.User.Columns().Account, accountId).
		Where(dao.User.Columns().Space, spaceId).One()
	if err != nil {
		return
	}
	if v.IsEmpty() {
		err = code.ErrUserNotExist.WithDetail(spaceId)
		return
	}
	usr = new(do.User)
	err = v.Struct(usr)
	return
}

func (u sUser) Exist(ctx context.Context, accountId string, spaceId int64) (exist bool, err error) {
	count, err := dao.User.Ctx(ctx).
		Where(dao.User.Columns().Account, accountId).
		Where(dao.User.Columns().Space, spaceId).Count()
	if err != nil {
		return
	}
	exist = count > 0
	return
}

func (u sUser) CreateSpaceUser(ctx context.Context, account *do.Account, spaceId int64) (usr *do.User, err error) {
	usr = &do.User{
		Account:   account.Id,
		Space:     spaceId,
		Type:      consts.UserTypeNormal,
		Name:      account.Name,
		CreatedAt: gtime.Now(),
	}
	_, err = dao.User.Ctx(ctx).Insert(usr)
	return
}

func (u sUser) IsSpaceManager(ctx context.Context) (ok bool, err error) {
	tokenInfo := service.Token().GetTokenInfoFromCtx(ctx)
	if tokenInfo.SpaceId == 0 {
		return
	}
	usr, err := u.GetUserByAccountId(ctx, tokenInfo.AccountId, tokenInfo.SpaceId)
	if err != nil {
		return
	}
	if gconv.Int(usr.Type) == consts.UserTypeManager {
		ok = true
		return
	}
	space, err := service.Space().GetSpaceInfo(ctx, tokenInfo.SpaceId)
	if err != nil {
		return
	}
	ok = space.IsOwner
	return
}
