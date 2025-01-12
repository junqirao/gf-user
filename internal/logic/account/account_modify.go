package account

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"

	"gf-user/internal/dao"
	"gf-user/internal/model"
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

func (s sAccount) ModifyPassword(ctx context.Context, in *model.AccountModifyPasswordInput) (err error) {
	tokenInfo := service.Token().GetTokenInfoFromCtx(ctx)
	acc, err := s.IsValid(ctx, tokenInfo.AccountId)
	if err != nil {
		return
	}
	if err = s.VerifyMFACode(ctx, acc, in.MFACode); err != nil {
		return
	}
	if gmd5.MustEncrypt(fmt.Sprintf("%v%s", acc.Password, in.Nonce)) != in.Old {
		err = code.ErrAccountPassword
		return
	}
	_, err = dao.Account.Ctx(ctx).Where(dao.Account.Columns().Id, acc.Id).Update(g.Map{
		dao.Account.Columns().Password: in.New,
	})
	return
}
