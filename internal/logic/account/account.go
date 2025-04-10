package account

import (
	"context"
	"fmt"
	"net/url"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/junqirao/gocomponents/response"
	uuid "github.com/satori/go.uuid"

	"gf-user/internal/consts"
	"gf-user/internal/dao"
	"gf-user/internal/model"
	"gf-user/internal/model/code"
	"gf-user/internal/model/do"
	"gf-user/internal/model/entity"
	"gf-user/internal/service"
)

func init() {
	service.RegisterAccount(&sAccount{})
}

type sAccount struct {
}

func (s sAccount) Register(ctx context.Context, in *model.AccountRegisterInput) (out *model.UserAccount, err error) {
	count, err := dao.Account.Ctx(ctx).Where(dao.Account.Columns().Account, in.Account).Count()
	if err != nil {
		return
	}
	if count > 0 {
		err = response.CodeConflict.WithMessage(in.Account)
		return
	}
	now, _ := gtime.Now().ToZone("Asia/Shanghai")
	ea := entity.Account{
		Id:        uuid.NewV4().String(),
		Account:   in.Account,
		Password:  in.Password,
		Type:      in.Type,
		Status:    in.Status,
		Name:      in.Name,
		Email:     in.Email,
		Avatar:    in.Avatar,
		CreatedAt: now,
		UpdateAt:  nil,
		Extra:     "{}",
	}
	if len(in.Extra) > 0 {
		ea.Extra = gjson.MustEncodeString(in.Extra)
	}
	err = dao.Account.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		_, err = dao.Account.Ctx(ctx).Insert(ea)
		if err != nil {
			return
		}
		typ := consts.UserTypeNormal
		if in.Administrator {
			typ = consts.UserTypeManager
		}
		_, err = dao.User.Ctx(ctx).Insert(entity.User{
			Account:   ea.Id,
			Space:     consts.DefaultSpaceId,
			Type:      typ,
			Name:      in.Name,
			CreatedAt: now,
		})
		return
	})
	if err != nil {
		return
	}

	account, err := s.IsValid(ctx, ea.Id)
	if err != nil {
		return
	}
	usr, err := service.User().GetUserByAccountId(ctx, ea.Id, consts.DefaultSpaceId)
	if err != nil {
		return
	}
	if out, err = s.getUserAccount(ctx, usr, account); err != nil {
		return
	}
	s.setAvatar(ctx, out.Account)
	return
}

func (s sAccount) UserLogin(ctx context.Context, in *model.AccountLoginInput) (out *model.UserAccountLoginInfo, err error) {
	acc, err := dao.Account.Ctx(ctx).
		Where(dao.Account.Columns().Account, in.Account).
		Where(dao.Account.Columns().Type, consts.AccountTypeNormal).
		One()
	if err != nil {
		return
	}
	if acc.IsEmpty() {
		err = code.ErrAccountNotExist.WithDetail(in.Account)
		return
	}
	var account = new(do.Account)
	if err = acc.Struct(&account); err != nil {
		return
	}
	if account.Status == consts.AccountStatusFrozen {
		err = code.ErrAccountLocked
		return
	}

	if in.Password != gmd5.MustEncrypt(fmt.Sprintf("%v%s", account.Password, in.Nonce)) {
		err = code.ErrAccountPassword
		return
	}

	user, err := dao.User.Ctx(ctx).
		Where(dao.User.Columns().Account, account.Id).
		Where(dao.User.Columns().Space, consts.DefaultSpaceId).
		One()
	if err != nil {
		return
	}
	if user.IsEmpty() {
		err = code.ErrAccountNotExist.WithDetail(in.Account)
		return
	}
	var usr = new(do.User)
	if err = user.Struct(&usr); err != nil {
		return
	}
	ua, err := s.getUserAccount(ctx, usr, account)
	if err != nil {
		return
	}
	out = &model.UserAccountLoginInfo{
		UserAccount: ua,
	}

	ext := model.RefreshTokenExtraData{
		From: consts.TokenFromUnknown,
	}
	if req := ghttp.RequestFromCtx(ctx); req != nil {
		ext.ClientIP = req.GetClientIp()
		ext.UA = req.UserAgent()
	}
	if in.From != "" {
		ext.From = in.From
	}
	out.AccessToken, out.RefreshToken, err = service.Token().GenerateAccessToken(ctx, out.UserAccount, ext)
	if err != nil {
		return
	}
	s.setAvatar(ctx, out.UserAccount.Account, out.AccessToken)
	return
}

func (s sAccount) UserLogout(ctx context.Context, refreshToken string) (err error) {
	claims, err := service.Token().ParseRefreshToken(ctx, refreshToken)
	if err != nil {
		return
	}

	accountId := ""
	if len(claims.Audience) == 0 {
		return
	}
	accountId = claims.Audience[0]
	err = service.Token().RemoveRefreshToken(ctx, accountId, claims)
	return
}

func (s sAccount) IsValid(ctx context.Context, accountId string) (acc *do.Account, err error) {
	v, err := dao.Account.Ctx(ctx).Where(dao.Account.Columns().Id, accountId).One()
	if err != nil {
		return
	}
	if v.IsEmpty() {
		err = code.ErrAccountNotExist.WithDetail(accountId)
		return
	}
	acc = new(do.Account)
	if err = v.Struct(acc); err != nil {
		return
	}

	if gconv.Int(acc.Status) == consts.AccountStatusFrozen {
		err = code.ErrAccountLocked
	}
	return
}

func (s sAccount) RefreshToken(ctx context.Context, spaceId int64, refreshToken string) (res *model.UserAccountLoginInfo, err error) {
	claims, err := service.Token().ParseRefreshToken(ctx, refreshToken)
	if err != nil {
		return
	}

	if spaceId <= 0 {
		spaceId = consts.DefaultSpaceId
	}

	// for compatibility
	ctx = context.WithValue(ctx, consts.CtxKeyTokenInfo, &model.TokenInfo{
		AccountId: claims.Audience[0],
	})

	account, err := s.IsValid(ctx, claims.Audience[0])
	if err != nil {
		return
	}
	usr, err := service.User().GetUserByAccountId(ctx, gconv.String(account.Id), spaceId)
	if err != nil {
		return
	}
	ua, err := s.getUserAccount(ctx, usr, account)
	if err != nil {
		return
	}

	newAccessToken, newRefreshToken, err := service.Token().RefreshToken(ctx, ua, claims)
	if err != nil {
		return
	}

	s.setAvatar(ctx, ua.Account, newAccessToken)

	res = &model.UserAccountLoginInfo{
		UserAccount:  ua,
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	}
	return
}

func (s sAccount) GenerateAppToken(ctx context.Context, appId string, spaceId int64, refreshToken string) (res *model.UserAccountLoginInfo, err error) {
	claims, err := service.Token().ParseRefreshToken(ctx, refreshToken)
	if err != nil {
		return
	}

	// for compatibility
	ctx = context.WithValue(ctx, consts.CtxKeyTokenInfo, &model.TokenInfo{
		AccountId: claims.Audience[0],
	})

	// check app
	app, err := service.App().Info(ctx, appId)
	if err != nil {
		return
	}

	// if spaceId <= 0 {
	// 	spaceId = consts.DefaultSpaceId
	// }
	if spaceId != gconv.Int64(app.Space) {
		err = code.ErrAppSpaceNotAllowed.WithDetail(spaceId)
		return
	}

	// check account
	account, err := s.IsValid(ctx, claims.Audience[0])
	if err != nil {
		return
	}
	usr, err := service.User().GetUserByAccountId(ctx, gconv.String(account.Id), spaceId)
	if err != nil {
		return
	}
	ua, err := s.getUserAccount(ctx, usr, account)
	if err != nil {
		return
	}

	// sign app token
	accessToken, err := service.Token().GenerateAppToken(ctx, app.Id, ua, claims)
	if err != nil {
		return
	}

	s.setAvatar(ctx, ua.Account, accessToken)

	res = &model.UserAccountLoginInfo{
		UserAccount: ua,
		AccessToken: accessToken,
	}
	return
}

func (s sAccount) GetUserAccount(ctx context.Context, spaceId ...int64) (ua *model.UserAccount, err error) {
	tokenInfo := service.Token().GetTokenInfoFromCtx(ctx)
	account, err := s.IsValid(ctx, tokenInfo.AccountId)
	if err != nil {
		return
	}
	si := tokenInfo.SpaceId
	if len(spaceId) > 0 && spaceId[0] > 0 {
		si = spaceId[0]
	}
	usr, err := service.User().GetUserFromToken(ctx, si)
	if err != nil {
		return
	}
	if ua, err = s.getUserAccount(ctx, usr, account); err != nil {
		return
	}
	s.setAvatar(ctx, ua.Account)
	return
}

func (s sAccount) GetAccount(ctx context.Context, account string) (acc *do.Account, err error) {
	v, err := dao.Account.Ctx(ctx).Where(dao.Account.Columns().Account, account).One()
	if err != nil {
		return
	}
	if v.IsEmpty() {
		err = code.ErrAccountNotExist.WithDetail(account)
		return
	}
	acc = new(do.Account)
	if err = v.Struct(acc); err != nil {
		return
	}
	if gconv.Int(acc.Status) != consts.AccountStatusNormal {
		err = code.ErrAccountLocked
	}
	return
}

func (s sAccount) GetAccountById(ctx context.Context, id string) (acc *model.Account, err error) {
	v, err := s.IsValid(ctx, id)
	if err != nil {
		return
	}
	acc = model.NewAccount(v)
	s.setAvatar(ctx, acc)
	return
}

func (s sAccount) GetAccountByIds(ctx context.Context, id []string) (acs []*model.Account, err error) {
	accounts, err := dao.Account.Ctx(ctx).WhereIn(dao.Account.Columns().Id, id).All()
	if err != nil {
		return
	}
	for _, v := range accounts {
		acc := new(do.Account)
		if err = v.Struct(&acc); err != nil {
			return
		}
		account := model.NewAccount(acc)
		s.setAvatar(ctx, account)
		acs = append(acs, account)
	}
	return
}

func (s sAccount) getUserAccount(ctx context.Context, usr *do.User, account *do.Account) (ua *model.UserAccount, err error) {
	spaces, err := service.Space().GetSpaceList(ctx, gconv.String(account.Id))
	if err != nil {
		return
	}
	ua = model.NewUserAccount(account, usr, spaces...)
	return
}

func (s sAccount) setAvatar(ctx context.Context, acc *model.Account, token ...string) {
	accessToken := ""
	if len(token) > 0 && token[0] != "" {
		accessToken = token[0]
	} else {
		t := service.Token().GetTokenInfoFromCtx(ctx)
		accessToken = t.AccessToken
	}
	acc.AvatarKey = acc.Avatar
	key := gconv.String(acc.Avatar)
	if key == "" {
		return
	}
	values := url.Values{}
	values.Set("access_token", accessToken)
	values.Set("key", key)
	values.Set("account_id", gconv.String(acc.Id))
	// using internal redirect instead of sign storage url directly
	acc.Avatar = fmt.Sprintf("/v1/storage/account/avatar?%s", values.Encode())
}

func (s sAccount) Exists(ctx context.Context, account string) (exists bool, err error) {
	count, err := dao.Account.Ctx(ctx).Where(dao.Account.Columns().Account, account).Count()
	if err != nil {
		return
	}
	exists = count > 0
	return
}
