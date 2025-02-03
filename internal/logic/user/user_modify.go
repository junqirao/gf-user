package user

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	"gf-user/internal/dao"
	"gf-user/internal/service"
)

func (u sUser) ModifyUserName(ctx context.Context, name string) (err error) {
	tokenInfo := service.Token().GetTokenInfoFromCtx(ctx)
	account, err := service.Account().IsValid(ctx, tokenInfo.AccountId)
	if err != nil {
		return
	}

	_, err = dao.User.Ctx(ctx).Where(g.Map{
		dao.User.Columns().Account: account.Id,
		dao.User.Columns().Id:      tokenInfo.UserId,
	}).Update(g.Map{
		dao.User.Columns().Name: name,
	})
	return
}
