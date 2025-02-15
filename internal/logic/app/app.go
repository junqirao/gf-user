package app

import (
	"context"
	"errors"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/junqirao/gocomponents/kvdb"
	uuid "github.com/satori/go.uuid"

	"gf-user/internal/dao"
	"gf-user/internal/model"
	"gf-user/internal/model/code"
	"gf-user/internal/model/entity"
	"gf-user/internal/service"
)

const (
	stoName = "app_cache"
)

func init() {
	service.RegisterApp(&sApp{})
}

type sApp struct {
}

func (s sApp) Create(ctx context.Context, in *model.CreateAppInput) (out *model.CreateAppOutput, err error) {
	token := service.Token().GetTokenInfoFromCtx(ctx)
	id := uuid.NewV4().String()
	secret := grand.S(48)
	out = &model.CreateAppOutput{
		AppId:     id,
		AppSecret: secret,
	}
	app := &entity.App{
		Id:           id,
		Secret:       gmd5.MustEncrypt(secret),
		Name:         in.Name,
		Space:        int(token.SpaceId),
		Descriptions: in.Description,
		Profile:      gjson.MustEncodeString(in.Profile),
		ExpiredAt:    in.ExpiredAt,
		CreatedAt:    gtime.Now(),
	}
	_, err = dao.App.Ctx(ctx).Data(app).Insert()
	if err != nil {
		return
	}
	err = s.setAppCache(ctx, app)
	return
}

func (s sApp) List(ctx context.Context) (infos []*model.AppInfo, err error) {
	token := service.Token().GetTokenInfoFromCtx(ctx)
	result, err := dao.App.Ctx(ctx).Where(dao.App.Columns().Space, token.SpaceId).All()
	if err != nil {
		return
	}
	infos = make([]*model.AppInfo, 0)
	es := make([]*entity.App, 0)
	if err = result.Structs(&es); err != nil {
		return
	}
	for _, e := range es {
		infos = append(infos, model.NewAppInfo(e))
	}
	return
}

func (s sApp) Info(ctx context.Context, id string) (info *model.AppInfo, err error) {
	cache, err := s.getAppCache(ctx, id)
	if err != nil {
		return
	}
	info = model.NewAppInfo(cache)
	return
}

func (s sApp) Update(ctx context.Context, in *model.UpdateAppInput) (err error) {
	var app *entity.App
	defer func() {
		if err != nil || app == nil {
			return
		}
		err = s.setAppCache(ctx, app)
	}()

	v, err := dao.App.Ctx(ctx).Where(dao.App.Columns().Id, in.AppId).One()
	if err != nil {
		return
	}
	if v.IsEmpty() {
		err = code.ErrInvalidAppId
		return
	}
	app = new(entity.App)
	if err = v.Struct(&app); err != nil {
		return
	}

	if len(in.Profile) > 0 {
		pf := make(map[string]any)
		if gjson.Valid(app.Profile) {
			if err = gjson.Unmarshal([]byte(app.Profile), &pf); err != nil {
				return
			}
		}
		for key, val := range in.Profile {
			pf[key] = val
		}
		app.Profile = gjson.MustEncodeString(pf)
	}
	if in.ExpiredAt != nil {
		app.ExpiredAt = in.ExpiredAt
	}
	if in.Description != nil {
		app.Descriptions = *in.Description
	}
	if in.Name != nil {
		app.Name = *in.Name
	}
	_, err = dao.App.Ctx(ctx).Where(dao.App.Columns().Id, in.AppId).Data(app).Update()
	return
}

func (s sApp) Auth(ctx context.Context, in *model.ValidateAppInput) (err error) {
	info, err := s.getAppCache(ctx, in.AppId)
	if err != nil {
		return
	}
	if info.ExpiredAt != nil && info.ExpiredAt.Before(gtime.Now()) {
		err = code.ErrAppExpired
		_ = s.Remove(ctx, in.AppId)
		return
	}
	if in.AppSecret != gmd5.MustEncrypt(info.Secret+in.Nonce) {
		err = code.ErrInvalidAppSecret
		return
	}
	return
}

func (s sApp) Remove(ctx context.Context, id string) (err error) {
	_, err = dao.App.Ctx(ctx).Where(dao.App.Columns().Id, id).Delete()
	if err != nil {
		return
	}
	err = s.removeAppCache(ctx, id)
	return
}

func (s sApp) setAppCache(ctx context.Context, app *entity.App) error {
	return kvdb.Storages.GetStorage(stoName).Set(ctx, app.Id, app)
}

func (s sApp) getAppCache(ctx context.Context, id string) (cache *entity.App, err error) {
	v, err := kvdb.Storages.GetStorage(stoName).Get(ctx, id)
	if err != nil {
		if errors.Is(err, kvdb.ErrStorageNotFound) {
			err = code.ErrInvalidAppId
		}
		return
	}
	if len(v) == 0 {
		err = code.ErrInvalidAppId
		return
	}
	cache = new(entity.App)
	err = gconv.Struct(v[0].Value, &cache)
	return
}

func (s sApp) removeAppCache(ctx context.Context, id string) error {
	return kvdb.Storages.GetStorage(stoName).Delete(ctx, id)
}
