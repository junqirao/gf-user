package account

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"

	"gf-user/internal/dao"
	"gf-user/internal/model/code"
	"gf-user/internal/service"
)

func (s sAccount) ModifyName(ctx context.Context, name string) (err error) {
	tokenInfo := service.Token().GetTokenInfoFromCtx(ctx)
	account, err := s.IsValid(ctx, tokenInfo.AccountId)
	if err != nil {
		return
	}

	_, err = dao.Account.Ctx(ctx).Where(dao.Account.Columns().Id, account.Id).Update(g.Map{
		dao.Account.Columns().Name: name,
	})
	return
}

func (s sAccount) ModifyAvatar(ctx context.Context, avatar string) (err error) {
	tokenInfo := service.Token().GetTokenInfoFromCtx(ctx)
	account, err := s.IsValid(ctx, tokenInfo.AccountId)
	if err != nil {
		return
	}

	_, err = dao.Account.Ctx(ctx).Where(dao.Account.Columns().Id, account.Id).Update(g.Map{
		dao.Account.Columns().Avatar: avatar,
	})
	return
}

func (s sAccount) ModifyPassword(ctx context.Context, oldPwd, newPwd string, nonce string) (err error) {
	tokenInfo := service.Token().GetTokenInfoFromCtx(ctx)
	account, err := s.IsValid(ctx, tokenInfo.AccountId)
	if err != nil {
		return
	}
	if gmd5.MustEncrypt(fmt.Sprintf("%v%s", account.Password, nonce)) != oldPwd {
		err = code.ErrAccountPassword
		return
	}
	_, err = dao.Account.Ctx(ctx).Where(dao.Account.Columns().Id, account.Id).Update(g.Map{
		dao.Account.Columns().Password: newPwd,
	})
	return
}
