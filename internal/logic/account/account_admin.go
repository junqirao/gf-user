package account

import (
	"context"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"gf-user/internal/consts"
	"gf-user/internal/model"
	"gf-user/internal/model/code"
	"gf-user/internal/service"
)

func (s sAccount) RegisterAdministrator(ctx context.Context, in *model.AccountRegisterInput) (out *model.UserAccount, err error) {
	initialized, err := service.Config().SystemInitialized(ctx)
	if err != nil {
		return
	}
	if initialized {
		err = code.ErrAdminAlreadyExist
		return
	}

	var (
		rawCod = g.Cfg().MustGet(ctx, "admin.code").String()
		nonce  = gconv.String(in.Extra["nonce"])
		cod    = gconv.String(in.Extra[consts.AccountExtraKeyAdminCode])
	)

	if rawCod == "" || gmd5.MustEncrypt(rawCod+nonce) != cod {
		err = code.ErrAdminCodeMismatch
		return
	}

	// overwrite after check
	in.Extra = map[string]any{
		consts.AccountExtraKeyAdminCode: rawCod,
	}
	return s.Register(ctx, in)
}
