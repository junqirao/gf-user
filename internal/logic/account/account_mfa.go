package account

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/crypto/gaes"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/junqirao/gocomponents/mfa"

	"gf-user/internal/dao"
	"gf-user/internal/model/code"
	"gf-user/internal/model/do"
	"gf-user/internal/service"
)

func (s sAccount) GenerateMFAQRCode(ctx context.Context) (qrCode string, err error) {
	token := service.Token().GetTokenInfoFromCtx(ctx)
	account, err := s.IsValid(ctx, token.AccountId)
	if err != nil {
		return
	}
	if len(account.Mfa) > 0 {
		err = code.ErrAccountMfaAlreadyBind
		return
	}
	cfg := service.Config().GetMFAConfig(ctx)
	authenticator := mfa.NewGoogleAuthenticator(cfg.CodeLength, cfg.SecretLength)
	secret, err := authenticator.CreateSecret()
	if err != nil {
		return
	}
	err = s.setMfaBindCache(ctx, account.Id, secret)
	if err != nil {
		return
	}

	return authenticator.GenerateQRCode(fmt.Sprintf("%s (%s)", cfg.Title, account.Name), secret)
}

func (s sAccount) BindMFA(ctx context.Context, mfaCode string) (err error) {
	token := service.Token().GetTokenInfoFromCtx(ctx)
	account, err := s.IsValid(ctx, token.AccountId)
	if err != nil {
		return
	}
	if len(account.Mfa) > 0 {
		err = code.ErrAccountMfaAlreadyBind
		return
	}
	cfg := service.Config().GetMFAConfig(ctx)
	authenticator := mfa.NewGoogleAuthenticator(cfg.CodeLength, cfg.SecretLength)
	cacheKey := fmt.Sprintf("mfa:bind:%v", account.Id)
	secret, err := g.Redis().Get(ctx, cacheKey)
	if err != nil || secret.IsNil() || secret.IsEmpty() {
		return
	}

	sc := secret.String()
	if !authenticator.VerifyCode(sc, mfaCode, cfg.VerifyDiscrepancy) {
		err = code.ErrAccountMfaCode
		return
	}
	encrypt, err := gaes.Encrypt(secret.Bytes(), []byte(cfg.Secret))
	if err != nil {
		return
	}
	_, err = dao.Account.Ctx(ctx).Where(dao.Account.Columns().Id, account.Id).Data(g.Map{
		dao.Account.Columns().Mfa: encrypt,
	}).Update()
	if err == nil {
		_, err = g.Redis().Unlink(ctx, cacheKey)
	}
	return
}

func (s sAccount) VerifyMFACode(ctx context.Context, account *do.Account, mfaCode string) (err error) {
	cfg := service.Config().GetMFAConfig(ctx)
	if !cfg.Enable {
		return nil
	}
	if len(account.Mfa) == 0 {
		err = code.ErrAccountMfaNotBind
		return
	}
	authenticator := mfa.NewGoogleAuthenticator(cfg.CodeLength, cfg.SecretLength)
	secret, err := gaes.Decrypt(account.Mfa, []byte(cfg.Secret))
	if err != nil {
		return
	}
	if !authenticator.VerifyCode(string(secret), mfaCode, cfg.VerifyDiscrepancy) {
		err = code.ErrAccountMfaCode
	}
	return
}

func (s sAccount) setMfaBindCache(ctx context.Context, accountId any, secret string) (err error) {
	key := fmt.Sprintf("mfa:bind:%v", accountId)
	_, err = g.Redis().Set(ctx, key, secret)
	if err != nil {
		return
	}
	_, err = g.Redis().Expire(ctx, key, 300)
	return
}
